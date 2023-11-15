package manager

import (
	"context"
	"log"
	"time"

	"acc.accommodation.com/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AvailableIntervalManager struct {
	available_interval_collection *mongo.Collection

	client *mongo.Client
}

func NewAvailableIntervalManager(available_interval_collection *mongo.Collection, client *mongo.Client) *AvailableIntervalManager {
	return &AvailableIntervalManager{
		available_interval_collection: available_interval_collection,
		client:                        client,
	}
}

func (a *AvailableIntervalManager) GetAvailableIntervalsByAccommodationId(ctx context.Context, accommodationId primitive.ObjectID) ([]*model.AvailableInterval, error) {
	// Find the accommodation
	availableIntervals := []*model.AvailableInterval{}

	//cursor with sorted start_date and start or end is greather than today
	cursor, err := a.available_interval_collection.Find(ctx, bson.M{
		"accommodation_id": accommodationId,
		"$or": bson.A{
			bson.M{"start_date": bson.M{"$gte": time.Now()}},
			bson.M{"end_date": bson.M{"$gte": time.Now()}},
		},
	}, options.Find().SetSort(bson.D{{"start_date", 1}}))

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		availableInterval := &model.AvailableInterval{}
		if err = cursor.Decode(availableInterval); err != nil {
			return nil, err
		}
		availableIntervals = append(availableIntervals, availableInterval)
	}
	return availableIntervals, nil

}

// get affected intervals by new interval
func (a *AvailableIntervalManager) getAffectedIntervals(ctx context.Context, availableIntervals *model.AvailableInterval) ([]*model.AvailableInterval, error) {
	// Find the accommodation
	affectedIntervals := []*model.AvailableInterval{}
	cursor, err := a.available_interval_collection.Find(ctx, bson.M{
		"accommodation_id": availableIntervals.AccommodationId,
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": availableIntervals.StartDate}},
				{"start_date": bson.M{"$lte": availableIntervals.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": availableIntervals.StartDate}},
				{"end_date": bson.M{"$lte": availableIntervals.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": availableIntervals.StartDate}},
				{"end_date": bson.M{"$gte": availableIntervals.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": availableIntervals.StartDate}},
				{"end_date": bson.M{"$lte": availableIntervals.EndDate}},
			}},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		availableInterval := &model.AvailableInterval{}
		if err = cursor.Decode(availableInterval); err != nil {
			return nil, err
		}
		affectedIntervals = append(affectedIntervals, availableInterval)
	}
	return affectedIntervals, nil
}

// update affected intervals
func (a *AvailableIntervalManager) updateAffectedIntervals(ctx context.Context, affectedIntervals []*model.AvailableInterval, newInterval *model.AvailableInterval) (bool, error) {
	shoudInsert := true

	// update affected intervals
	// ---|----|---- <- new interval
	// -|---|------- -> update
	// -|---------|- -> update and insert
	// -----|----|-- -> update
	// -----|--|---- -> delete
	for _, affectedInterval := range affectedIntervals {
		// start before
		if affectedInterval.StartDate.Before(newInterval.StartDate) {
			if affectedInterval.EndDate.Before(newInterval.EndDate) || affectedInterval.EndDate.Equal(newInterval.EndDate) {
				affectedInterval.EndDate = newInterval.StartDate.AddDate(0, 0, -1)
			} else {
				affectedInterval2 := &model.AvailableInterval{
					Id:              primitive.NewObjectID(),
					AccommodationId: newInterval.AccommodationId,
					StartDate:       newInterval.EndDate.AddDate(0, 0, 1),
					EndDate:         affectedInterval.EndDate,
					IsAvailable:     affectedInterval.IsAvailable,
				}

				_, err := a.available_interval_collection.InsertOne(ctx, affectedInterval2)
				if err != nil {
					return shoudInsert, err
				}
				affectedInterval.EndDate = newInterval.StartDate.AddDate(0, 0, -1)
			}
			// start after
		} else if affectedInterval.StartDate.After(newInterval.StartDate) || affectedInterval.StartDate.Equal(newInterval.StartDate) {
			if affectedInterval.EndDate.After(newInterval.EndDate) {
				affectedInterval.StartDate = newInterval.EndDate.AddDate(0, 0, 1)
			} else {
				// delete affected interval
				_, err := a.available_interval_collection.DeleteOne(ctx, bson.M{"_id": affectedInterval.Id})
				if err != nil {
					return shoudInsert, err
				}
				continue
			}
		}
		_, err := a.available_interval_collection.UpdateOne(ctx, bson.M{"_id": affectedInterval.Id}, bson.M{"$set": affectedInterval})
		if err != nil {
			return shoudInsert, err
		}
	}
	return shoudInsert, nil
}

// add new interval and handle overlapping intervals (merge them)
func (a *AvailableIntervalManager) AddAvailableInterval(ctx context.Context, newInterval *model.AvailableInterval) (interface{}, error) {

	session, err := a.client.StartSession()
	if err != nil {
		log.Println("Cannot start session")
		return nil, err
	}
	err = session.StartTransaction()
	if err != nil {
		log.Println("Cannot start transaction")
		return nil, err
	}
	// get Affected Intervals
	affectedIntervals, err := a.getAffectedIntervals(ctx, newInterval)
	if err != nil {
		log.Println("Cannot get affected intervals")
		return nil, err
	}
	// update affected intervals
	_, err = a.updateAffectedIntervals(ctx, affectedIntervals, newInterval)
	if err != nil {
		log.Println("Cannot update affected intervals")
		session.AbortTransaction(ctx)
		return nil, err
	}
	// insert new interval
	if !newInterval.IsAvailable {
		_, err = a.available_interval_collection.InsertOne(ctx, newInterval)
		if err != nil {
			log.Println("Cannot insert new interval")
			session.AbortTransaction(ctx)
			return nil, err
		}
	}
	err = session.CommitTransaction(ctx)
	if err != nil {
		log.Println("Cannot commit transaction")
		return nil, err
	}
	return "Success", nil

}

// filter accommodation ids by available intervals so that we can get only accommodations that are available in given period
func (a *AvailableIntervalManager) FilterAccommodationIdsByAvailableIntervals(ctx context.Context, accommodationIds []primitive.ObjectID, startDate time.Time, endDate time.Time) ([]primitive.ObjectID, error) {
	// Find the intervals
	availableIntervals := []*model.AvailableInterval{}
	//cursor with sorted start_date and get intervals that are in given period beetwen start and end date
	cursor, err := a.available_interval_collection.Find(ctx, bson.M{
		"accommodation_id": bson.M{"$in": accommodationIds},
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": startDate}},
				{"start_date": bson.M{"$lte": endDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": startDate}},
				{"end_date": bson.M{"$lte": endDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": startDate}},
				{"end_date": bson.M{"$gte": endDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": startDate}},
				{"end_date": bson.M{"$lte": endDate}},
			}},
		},
	}, options.Find().SetSort(bson.D{{Key: "start_date", Value: 1}}))

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		availableInterval := &model.AvailableInterval{}
		if err = cursor.Decode(availableInterval); err != nil {
			return nil, err
		}
		availableIntervals = append(availableIntervals, availableInterval)
	}
	// remove all accommodation ids that are not available in given period
	for _, availableInterval := range availableIntervals {
		for i, accommodationId := range accommodationIds {
			if accommodationId == availableInterval.AccommodationId {
				if !availableInterval.IsAvailable {
					accommodationIds = append(accommodationIds[:i], accommodationIds[i+1:]...)
				}
				break
			}
		}
	}
	return accommodationIds, nil

}
