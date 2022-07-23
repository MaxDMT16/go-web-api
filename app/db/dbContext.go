package db

import (
	"fmt"

	"getheadway/app/models"
)

type SubscriptionsRepository struct{}

func (db *SubscriptionsRepository) GetSubscriptionById(id int) (models.Subscription, error) {
	return models.Subscription{
		Id:    id,
		Name:  fmt.Sprintf("Subscription %v", id),
		Price: float64(id),
	}, nil
}
