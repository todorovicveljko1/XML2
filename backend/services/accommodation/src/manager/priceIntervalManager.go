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

type PriceIntervalManager struct {
	price_interval_collection *mongo.Collection

	client *mongo.Client
}

func NewPriceIntervalManager(price_interval_collection *mongo.Collection, client *mongo.Client) *PriceIntervalManager {
	return &PriceIntervalManager{
		price_interval_collection: price_interval_collection,
		client:                    client,
	}
}

func (a *PriceIntervalManager) GetPriceIntervalsByAccommodationId(ctx context.Context, accommodationId primitive.ObjectID) ([]*model.PriceInterval, error) {

	// Find the accommodation
	priceIntervals := []*model.PriceInterval{}
	//cursor with sorted start_date and start or end is greather than today
	cursor, err := a.price_interval_collection.Find(ctx, bson.M{
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
		priceInterval := &model.PriceInterval{}
		if err = cursor.Decode(priceInterval); err != nil {
			return nil, err
		}
		priceIntervals = append(priceIntervals, priceInterval)
	}
	return priceIntervals, nil

}

// get affected intervals by new interval
func (a *PriceIntervalManager) getAffectedIntervals(ctx context.Context, priceIntervals *model.PriceInterval) ([]*model.PriceInterval, error) {
	// Find the accommodation
	affectedIntervals := []*model.PriceInterval{}
	cursor, err := a.price_interval_collection.Find(ctx, bson.M{
		"accommodation_id": priceIntervals.AccommodationId,
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": priceIntervals.StartDate}},
				{"start_date": bson.M{"$lte": priceIntervals.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": priceIntervals.StartDate}},
				{"end_date": bson.M{"$lte": priceIntervals.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": priceIntervals.StartDate}},
				{"end_date": bson.M{"$gte": priceIntervals.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": priceIntervals.StartDate}},
				{"end_date": bson.M{"$lte": priceIntervals.EndDate}},
			}},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		priceInterval := &model.PriceInterval{}
		if err = cursor.Decode(priceInterval); err != nil {
			return nil, err
		}
		affectedIntervals = append(affectedIntervals, priceInterval)
	}
	return affectedIntervals, nil
}

func (a *PriceIntervalManager) updatePriceIntervals(ctx context.Context, affectedIntervals []*model.PriceInterval, newInterval *model.PriceInterval) (bool, error) {
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
				affectedInterval2 := &model.PriceInterval{
					Id:              primitive.NewObjectID(),
					AccommodationId: newInterval.AccommodationId,
					StartDate:       newInterval.EndDate.AddDate(0, 0, 1),
					EndDate:         affectedInterval.EndDate,
					Price:           affectedInterval.Price,
				}

				_, err := a.price_interval_collection.InsertOne(ctx, affectedInterval2)
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
				_, err := a.price_interval_collection.DeleteOne(ctx, bson.M{"_id": affectedInterval.Id})
				if err != nil {
					return shoudInsert, err
				}
				continue
			}
		}
		_, err := a.price_interval_collection.UpdateOne(ctx, bson.M{"_id": affectedInterval.Id}, bson.M{"$set": affectedInterval})
		if err != nil {
			return shoudInsert, err
		}
	}
	return shoudInsert, nil
}

func (a *PriceIntervalManager) AddPriceInterval(ctx context.Context, newInterval *model.PriceInterval) (interface{}, error) {

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
	_, err = a.updatePriceIntervals(ctx, affectedIntervals, newInterval)
	if err != nil {
		log.Println("Cannot update affected intervals")
		session.AbortTransaction(ctx)
		return nil, err
	}
	// insert new interval
	res, err := a.price_interval_collection.InsertOne(ctx, newInterval)
	if err != nil {
		log.Println("Cannot insert new interval")
		session.AbortTransaction(ctx)
		return nil, err
	}
	err = session.CommitTransaction(ctx)
	if err != nil {
		log.Println("Cannot commit transaction")
		return nil, err
	}
	return res.InsertedID, nil

}
