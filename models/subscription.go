package models

import (
	"database/sql"
	"errors"

	"github.com/pborman/uuid"
)

type Subscription struct {
	Id uuid.UUID `json: "id"`
	UserId int `json: "userId"`
	Status SubscriptionStatus `json:"status"` 
	CreatedAt string `json:"createdAt"` 
	StartAt string `json:"startAt"` 
	ShopId int `json: "shopId"`
	OzInBag int `json: "ozInBag"`
	BeanName string `json:"beanName"`
	RoastName string `json: "roastName"`
	Price int `json: "price"`
}

func NewSubscription(userId int, createdAt string, startAt string, shopId int, ozInBag int, beanName string, roastName string, price int) *Subscription {
	return &Subscription{ 
		Id: uuid.NewUUID(), 
		UserId: userId, 
		Status: ACTIVE, 
		CreatedAt: createdAt, 
		StartAt: startAt, 
		ShopId: shopId, 
		OzInBag: ozInBag, 
		BeanName: beanName,
		RoastName: roastName,
		Price: price,
	}
}


func SubscriptionFromSql(rows *sql.Rows) ([]*Subscription, error) {
	subscription := make([]*Subscription,0)

	for rows.Next() {
		s := &Subscription{}
		var sStatus string
		rows.Scan(&s.Id, &s.UserId, &sStatus, &s.CreatedAt, &s.StartAt, &s.ShopId, &s.OzInBag, &s.BeanName, &s.RoastName, &s.Price)
		subscription = append(subscription, s)

		var ok bool
		s.Status, ok = toSubscriptionType(sStatus)
		if !ok {
			return nil, errors.New("invalid subscriptionStatus string")
		}
		subscription = append(subscription, s)
	}


	return subscription, nil
}

func toSubscriptionType(s string) (SubscriptionStatus, bool) {
	switch s {
	case ACTIVE:
		return ACTIVE, true
	case PENDING:
		return PENDING, true
	case CANCELLED:
		return CANCELLED, true
	case INACTIVE:
		return INACTIVE, true
	default:
		return "", false
	}
}

/*SubscriptionStatus is an enum wrapper for valid subscription type*/
type SubscriptionStatus string 

const (
	ACTIVE = "ACTIVE"
	PENDING = "PENDING"
	CANCELLED = "CANCELLED"
	INACTIVE = "INACTIVE"
)
