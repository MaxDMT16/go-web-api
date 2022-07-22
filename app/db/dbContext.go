package db

import (
	"fmt"

	"getheadway/app/models"
)

type DbContext struct{}

func (db *DbContext) GetSubscriptionById(id int) (models.Subscription, error) {
	return models.Subscription{
		Id:    id,
		Name:  fmt.Sprintf("Subscription %v", id),
		Price: float64(id),
	}, nil
}
