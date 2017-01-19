package helpers

import (
	"gopkg.in/alexcesaro/statsd.v2"

	"github.com/yuderekyu/expresso-subscription/models"
	"github.com/yuderekyu/expresso-subscription/gateways"
)

type baseHelper struct {
	sql gateways.SQL
	stats *statsd.Client
}

/**/
type SubscriptionI interface {
	GetById(string) (*models.Subscription, error)
	GetAll(int, int) ([]*models.Subscription, error)
	Insert(*models.Subscription) error
	Update(*models.Subscription) error
	SetStatus(string, models.SubscriptionStatus) error
}

type Subscription struct {
	*baseHelper
}

func NewSubscription(sql gateways.SQL) *Subscription {
	return &Subscription{baseHelper: &baseHelper{sql: sql}}
}

/*Change to String*/
func (s *Subscription) GetById(id string) (*models.Subscription, error) {
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, startAt, shopId, ozInBag, beanName, roastName, price FROM subscription WHERE id =?", id)
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
	if err!= nil {
		return nil, err
	}
	return subscription, err
}

func (s *Subscription) Insert(subscription *models.Subscription) error {
	err := s.sql.Modify(
		"INSERT INTO subscription (id, userId, status, createdAt, startAt, shopId, ozInBag, beanName, roastName, price) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		subscription.Id,
		subscription.UserId,
		string(subscription.Status),
		subscription.CreatedAt,
		subscription.StartAt,
		subscription.ShopId, 
		subscription.OzInBag,
		subscription.BeanName,
		subscription.RoastName,
		subscription.Price,
		)
	return err
}

func (s *Subscription) Update(subscription *models.Subscription) error {
	err := s.sql.Modify("UPDATE subscription SET status=?, startAt=?, shopId=?, ozInBag=?, beanName=?, roastName=?, price=? WHERE id=?", 
		string(subscription.Status),
		subscription.StartAt,
		subscription.ShopId,
		subscription.OzInBag,
		subscription.BeanName,
		subscription.RoastName,
		subscription.Price,
		subscription.Id,
	)
	return err
}

func (s *Subscription) SetStatus(id string, status models.SubscriptionStatus) error {
	err := s.sql.Modify("UPDATE subscription SET status=? WHERE id=?", string(status), id)
	return err
}