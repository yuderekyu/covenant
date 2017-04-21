package helpers

import (
	"errors"
	"github.com/ghmeier/bloodlines/gateways"
	c "github.com/ghmeier/coinage/gateways"
	coinM "github.com/ghmeier/coinage/models"
	t "github.com/jakelong95/TownCenter/gateways"
	w "github.com/lcollin/warehouse/gateways"
	wareM "github.com/lcollin/warehouse/models"
	"github.com/pborman/uuid"
	"github.com/yuderekyu/covenant/models"
)

const (
	SELECT_ALL = "SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity, nextOrder FROM subscription"
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
	GetByUserAndItem(userID uuid.UUID, itemID uuid.UUID) (*models.Subscription, error)
	Insert(*models.Subscription) error
	Update(string, *models.Subscription) error
	SetStatus(string, models.SubscriptionStatus) error
	Delete(string) error
	NewOrder(*models.Subscription, *models.RequestOrder) (*wareM.Order, error)
	Subscribe(uuid.UUID, uuid.UUID, uuid.UUID, string, uint64) error
	CheckCustomer(uuid.UUID) (*coinM.Customer, error)
}

/*Subscription is the helper for subscription entries*/
type Subscription struct {
	*baseHelper
	TownCenter t.TownCenterI
	Warehouse  w.Warehouse
	Coinage    c.Coinage
}

/*NewSubscription returns a new Subscription helper*/
func NewSubscription(sql gateways.SQL, tc t.TownCenterI, wh w.Warehouse, coin c.Coinage) *Subscription {
	return &Subscription{
		baseHelper: &baseHelper{sql: sql},
		TownCenter: tc,
		Warehouse:  wh,
		Coinage:    coin,
	}
}

/*GetById returns the subscription referenced by provided id, otherwise nil*/
func (s *Subscription) GetByID(id string) (*models.Subscription, error) {
	rows, err := s.sql.Select(SELECT_ALL+" WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}

	if len(subscription) == 0 {
		return nil, nil
	}

	return subscription[0], err
}

/*GetAll returns <limit> subscription entries from <offset> number*/
func (s *Subscription) GetAll(offset int, limit int) ([]*models.Subscription, error) {
	rows, err := s.sql.Select(SELECT_ALL+" ORDER BY id ASC LIMIT ?,?", offset, limit)
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
	rows, err := s.sql.Select(SELECT_ALL+" WHERE roasterId=? ORDER BY id ASC LIMIT ?,?", roasterID, offset, limit)
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
func (s *Subscription) GetByUser(userID string, offset int, limit int) ([]*models.Subscription, error) {
	rows, err := s.sql.Select(SELECT_ALL+" WHERE userId=? ORDER BY id ASC LIMIT ?,?", userID, offset, limit)
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

/*GetByUserAndItem checks if the subscription with the given userID and itemID exists; returns this subscription entry */
func (s *Subscription) GetByUserAndItem(userID uuid.UUID, itemID uuid.UUID) (*models.Subscription, error) {
	rows, err := s.sql.Select(SELECT_ALL+" WHERE userId=? AND itemId=?", userID.String(), itemID.String())
	if err != nil {
		return nil, err
	}

	subscription, err := models.SubscriptionFromSql(rows)
	if err != nil {
		return nil, err
	}

	if len(subscription) == 0 {
		return nil, nil
	}

	if len(subscription) > 1 {
		return nil, errors.New("Subscription already exist for user")
	}

	return subscription[0], err
}

/*Insert adds the given subscription entry*/
func (s *Subscription) Insert(subscription *models.Subscription) error {
	err := s.sql.Modify(
		"INSERT INTO subscription (id, userId, status, createdAt, frequency, roasterId, itemId, quantity, nextOrder) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		subscription.ID,
		subscription.UserID,
		string(subscription.Status),
		subscription.CreatedAt,
		string(subscription.Frequency),
		subscription.RoasterID,
		subscription.ItemID,
		subscription.Quantity,
		subscription.NextOrder,
	)
	return err
}

/*Update upserts the subscription with the given id*/
func (s *Subscription) Update(id string, subscription *models.Subscription) error {
	err := s.sql.Modify("UPDATE subscription SET status=?, frequency=?, roasterId=?, itemId=?, quantity=?, nextOrder=? WHERE id=?",
		string(subscription.Status),
		string(subscription.Frequency),
		subscription.RoasterID,
		subscription.ItemID,
		subscription.Quantity,
		subscription.NextOrder,
		id,
	)
	return err
}

/*Delete removes the subscription with the given id*/
func (s *Subscription) Delete(id string) error {
	err := s.sql.Modify("DELETE FROM subscription where id=?", id)
	// s.Coinage.DeleteSubscription()
	return err
}

/*SetStatus updates the status of the subscription with the given id*/
func (s *Subscription) SetStatus(id string, status models.SubscriptionStatus) error {
	err := s.sql.Modify("UPDATE subscription SET status=? WHERE id=?", string(status), id)
	return err
}

/*CreateOrder calls Warehouse's newOrder function to create a new subscription*/
func (s *Subscription) NewOrder(sub *models.Subscription, req *models.RequestOrder) (*wareM.Order, error) {
	order := wareM.NewOrder(sub.UserID, sub.ID, int(sub.Quantity))
	newOrder, err := s.Warehouse.NewOrder(order)
	if err != nil {
		return nil, err
	}

	sub.Status = models.ACTIVE
	sub.NextOrder = req.NextOrder
	s.Update(sub.ID.String(), sub)
	return newOrder, err
}

/*Subscribe calls Coinage Suscribe function to create a new subscription*/
func (s *Subscription) Subscribe(id uuid.UUID, roasterID uuid.UUID, itemID uuid.UUID, frequency string, quantity uint64) error {
	newFreq := coinM.Frequency(frequency)
	subscriptionRequest := coinM.NewSubscribeRequest(roasterID, itemID, newFreq, quantity)
	err := s.Coinage.NewSubscription(id, subscriptionRequest)
	return err
}

/*CheckCustomer checks coinage if the specified customer account exists*/
func (s *Subscription) CheckCustomer(id uuid.UUID) (*coinM.Customer, error) {
	customer, err := s.Coinage.Customer(id)
	return customer, err
}
