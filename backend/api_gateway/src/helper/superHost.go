package helper

import (
	"context"

	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
)

func IsSuperHost(clients *client.Clients, hostId string) (bool, error) {
	isSuperHost, err := clients.ReservationClient.CheckForSuperHost(context.Background(), &pb.IdRequest{Id: hostId})
	if err != nil {
		return false, err
	}
	// get rating
	rating, err := clients.RatingClient.HostRating(context.Background(), &pb.RatingIdRequest{Id: hostId})
	if err != nil {
		return false, err

	}
	// check if user has rating greather than 4.7
	if isSuperHost.Value && rating.Rating >= 4.7 {
		return true, nil
	}
	return false, nil
}

func FilterOutSuperHostIds(clients *client.Clients, hostIds []string) ([]string, error) {
	var filteredHostIds []string
	for _, hostId := range hostIds {
		isSuperHost, err := IsSuperHost(clients, hostId)
		if err != nil {
			return nil, err
		}
		if isSuperHost {
			filteredHostIds = append(filteredHostIds, hostId)
		}
	}
	return filteredHostIds, nil
}
