package helpers

import (
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/yuderekyu/covenant/models"
)

type baseHelper struct {
	sql gateways.SQL
}

type SubscriptionI interface {
	GetByID(string) (*models.Subscription, error)
	GetAll(int, int) ([]*models.Subscription, error)
	Insert(*models.Subscription) error
	Update(string, *models.Subscription) error
	SetStatus(string, models.SubscriptionStatus) error
	Delete(string) error
}

type Subscription struct {
	*baseHelper
}

func NewSubscription(sql gateways.SQL) *Subscription {
	return &Subscription{baseHelper: &baseHelper{sql: sql}}
}

func (s *Subscription) GetByID(id string) (*models.Subscription, error) {
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, startAt, roasterId, itemID FROM subscription WHERE id =?", id)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}
	return subscription[0], err
}

func (s *Subscription) GetAll(offset int, limit int) ([]*models.Subscription, error) {
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, startAt, shopId, ozInBag, beanName, roastName, price FROM subscription ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

func (s *Subscription) Insert(subscription *models.Subscription) error {
	err := s.sql.Modify(
		"INSERT INTO subscription (id, userId, status, createdAt, startAt, roasterId, itemId) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		subscription.ID,
		subscription.UserID,
		string(subscription.Status),
		subscription.CreatedAt,
		subscription.StartAt,
		subscription.RoasterID,
		subscription.ItemID,
	)
	return err
}

func (s *Subscription) Update(id string, subscription *models.Subscription) error {
	err := s.sql.Modify("UPDATE subscription SET status=?, startAt=?, roasterId=?, itemId=? WHERE id=?",
		string(subscription.Status),
		subscription.StartAt,
		subscription.RoasterID,
		subscription.ItemID,
		id,
	)
	return err
}

func (s *Subscription) Delete(id string) error {
	err := s.sql.Modify("DELETE FROM subscription where id=?", id)
	return err
}

func (s *Subscription) SetStatus(id string, status models.SubscriptionStatus) error {
	err := s.sql.Modify("UPDATE subscription SET status=? WHERE id=?", string(status), id)
	return err
}
