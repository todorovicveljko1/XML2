package helper

import (
	"context"
	"log"

	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
)

/*
message SendNotificationRequest {
    string type = 1;
    string resource_id = 2;
    string body = 3;
    string user_id = 4;
}
*/

type Notification struct {
	Type       string
	ResourceId string
	Body       string
	UserId     string
}

// Notification types as constants
const (
	ReservationCreated      = "reservation_created"
	ReservationCanceled     = "reservation_canceled"
	ReservationAccepted     = "reservation_accepted"
	ReservationRejected     = "reservation_rejected"
	RatingModified          = "rating_modified"
	SuperHostStatusModified = "super_host_status_modified" // tricky one
)

func SendNotification(ctx context.Context, clients *client.Clients, notification *Notification) (bool, error) {
	res, err := clients.NotificationClient.SendNotification(ctx, &pb.SendNotificationRequest{
		Type:       notification.Type,
		ResourceId: notification.ResourceId,
		Body:       notification.Body,
		UserId:     notification.UserId,
	})
	if err != nil {
		log.Println("Can't send notification to user: ", err)
		return false, err
	}

	return res.Success, nil

}
