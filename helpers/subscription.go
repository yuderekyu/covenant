package helpers

import (
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/yuderekyu/covenant/models"
	t"github.com/jakelong95/TownCenter/gateways"
	w"github.com/lcollin/warehouse/gateways"
)

type baseHelper struct {
	sql gateways.SQL
}

/*SubscriptionI describes the functions for manipulating subscription models*/
type SubscriptionI interface {
	GetByID(string) (*models.Subscription, error)
	GetAll(int, int) ([]*models.Subscription, error)
	GetByRoaster(string, int, int) ([]*models.Subscription, error)
	GetByUser(string, int, int) ([]*models.Subscription, error)
	Insert(*models.Subscription) error
	Update(string, *models.Subscription) error
	SetStatus(string, models.SubscriptionStatus) error
	Delete(string) error
}

/*Subscription is the helper for subscription entries*/
type Subscription struct {
	*baseHelper
	TownCenter t.TownCenterI
	Warehouse w.Warehouse
}

/*NewSubscription returns a new Subscription helper*/
func NewSubscription(sql gateways.SQL, tc t.TownCenterI, wh w.Warehouse) *Subscription {
	return &Subscription{
		baseHelper: &baseHelper{sql: sql},
		TownCenter: tc,
		Warehouse: wh,
	}
}

/*GetById returns the subscription referenced by provided id, otherwise nil*/
func (s *Subscription) GetByID(id string) (*models.Subscription, error) {
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription WHERE id =?", id)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}

	if(len(subscription) == 0) {
		return nil, nil
	}
	
	return subscription[0], err
}

/*GetAll returns <limit> subscription entries from <offset> number*/
func (s *Subscription) GetAll(offset int, limit int) ([]*models.Subscription, error) {
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

/*GetAll returns <limit> subscription entries from <offset> number referenced by provided roasterID*/
func (s *Subscription) GetByRoaster(roasterID string, offset int, limit int) ([]*models.Subscription, error) {
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription WHERE roasterId = ? ORDER BY id ASC LIMIT ?,?", roasterID, offset, limit)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

/*GetAll returns <limit> subscription entries from <offset> number referenced by provided userID*/
func (s *Subscription) GetByUser(userID string, offset int, limit int) ([]*models.Subscription, error){
	rows, err := s.sql.Select("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription WHERE userId = ? ORDER BY id ASC LIMIT ?,?", userID, offset, limit)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

/*Insert adds the given subscription entry*/
func (s *Subscription) Insert(subscription *models.Subscription) error {
	err := s.sql.Modify(
		"INSERT INTO subscription (id, userId, status, createdAt, frequency, roasterId, itemId) VALUE(?, ?, ?, ?, ?, ?, ?)",
		subscription.ID,
		subscription.UserID,
		string(subscription.Status),
		subscription.CreatedAt,
		subscription.Frequency,
		subscription.RoasterID,
		subscription.ItemID,
	)
	return err
}

/*Update upserts the subscription with the given id*/
func (s *Subscription) Update(id string, subscription *models.Subscription) error {
	err := s.sql.Modify("UPDATE subscription SET status=?, frequency=?, roasterId=?, itemId=? WHERE id=?",
		string(subscription.Status),
		subscription.Frequency,
		subscription.RoasterID,
		subscription.ItemID,
		id,
	)
	return err
}

/*Delete removes the subscription with the given id*/
func (s *Subscription) Delete(id string) error {
	err := s.sql.Modify("DELETE FROM subscription where id=?", id)
	return err
}

/*SetStatus updates the status of the subscription with the given id*/
func (s *Subscription) SetStatus(id string, status models.SubscriptionStatus) error {
	err := s.sql.Modify("UPDATE subscription SET status=? WHERE id=?", string(status), id)
	return err
}
