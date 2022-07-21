package db

import (
	"fmt"
	"strconv"

	"getheadway/app/model"
)

type DbContext struct{}

func (db *DbContext) GetSubscriptionById(id string) model.Subscription {
	price, _ := strconv.ParseFloat(id, 64)

	return model.Subscription{
		Id:    id,
		Name:  fmt.Sprintf("Subscription %v", id),
		Price: price,
	}
}
