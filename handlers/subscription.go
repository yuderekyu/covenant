package handlers

import (
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/pborman/uuid"
	"github.com/yuderekyu/covenant/helpers"
	"github.com/yuderekyu/covenant/models"
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"
)

/*SubscriptionI contains the methods for a subscription handler*/
type SubscriptionI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	ViewByRoaster(ctx *gin.Context)
	ViewByUser(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Time() gin.HandlerFunc
	GetJWT() gin.HandlerFunc
	CreateOrder(ctx *gin.Context)
}

/*Subscription is the handler for all subscription api calls*/
type Subscription struct {
	*handlers.BaseHandler
	Subscription helpers.SubscriptionI
}

/*NewSubscription returns a subscription handler*/
func NewSubscription(ctx *handlers.GatewayContext) SubscriptionI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.subscription"))
	return &Subscription{
		Subscription: helpers.NewSubscription(ctx.Sql, ctx.TownCenter, ctx.Warehouse, ctx.Coinage),
		BaseHandler:  &handlers.BaseHandler{Stats: stats},
	}
}

/*New adds the given subscription entry to the database*/
func (s *Subscription) New(ctx *gin.Context) {
	var json models.RequestSubscription
	err := ctx.BindJSON(&json)
	if err != nil {
		s.UserError(ctx, "Error: unable to parse json", err)
		return
	}

	//Check if user already has subscription with the specific ItemID
	subs, err := s.Subscription.GetByUserAndItem(json.UserID, json.ItemID)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}
	if subs != nil {
		s.UserError(ctx, "Error: you've alread subscribed to this", nil)
		return
	}

	//Check if customer account exists with Coinage
	_, err = s.Subscription.CheckCustomer(json.UserID)
	if err != nil {
		s.UserError(ctx, "Please update your payment information before subscribing", nil)
		return
	}

	//Create subscription within Covenant
	subscription := models.NewSubscription(json.UserID, json.Frequency, json.RoasterID, json.ItemID, json.Quantity)
	err = s.Subscription.Insert(subscription)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}

	//Create subscription within Coinage
	err = s.Subscription.Subscribe(subscription.UserID, subscription.RoasterID, subscription.ItemID, subscription.Frequency, subscription.Quantity)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}

	s.Success(ctx, subscription)
}

/*View returns the subscription entry with the given subscriptionId*/
func (s *Subscription) View(ctx *gin.Context) {
	id := ctx.Param("subscriptionId")

	subscription, err := s.Subscription.GetByID(id)
	if err != nil {
		s.ServerError(ctx, err, nil)
		return
	}

	if subscription == nil {
		s.UserError(ctx, "Error: Subscription does not exist", id)
		return
	}

	s.Success(ctx, subscription)
}

/*ViewAll returns a list of subscriptions with offset and limit determining the entries and amount*/
func (s *Subscription) ViewAll(ctx *gin.Context) {
	offset, limit := s.GetPaging(ctx)
	subscriptions, err := s.Subscription.GetAll(offset, limit)
	if err != nil {
		s.ServerError(ctx, err, nil)
		return
	}

	s.Success(ctx, subscriptions)
}

/*ViewByRoaster returns a list of subscriptions with the given roasterId,
with the offset and limit determining the entries and amount*/
func (s *Subscription) ViewByRoaster(ctx *gin.Context) {
	roasterID := ctx.Param("roasterId")
	offset, limit := s.GetPaging(ctx)

	subscriptions, err := s.Subscription.GetByRoaster(roasterID, offset, limit)
	if err != nil {
		s.ServerError(ctx, err, nil)
		return
	}

	s.Success(ctx, subscriptions)
}

/*ViewByUser returns a list of subscriptions with the given roasterId,
with the offset and limit determining the entries and amount*/
func (s *Subscription) ViewByUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	offset, limit := s.GetPaging(ctx)

	subscriptions, err := s.Subscription.GetByUser(userID, offset, limit)
	if err != nil {
		s.ServerError(ctx, err, nil)
		return
	}

	s.Success(ctx, subscriptions)
}

/*Update overwrites a subscription with the given subscriptionId*/
func (s *Subscription) Update(ctx *gin.Context) {
	id := ctx.Param("subscriptionId")

	var json models.Subscription
	err := ctx.BindJSON(&json)
	if err != nil {
		s.UserError(ctx, "Error: unable to parse json", err)
		return
	}
	json.ID = uuid.Parse(id)

	err = s.Subscription.Update(id, &json)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}

	s.Success(ctx, json)
}

/*Delete removes a subscription with the given subscriptionId*/
func (s *Subscription) Delete(ctx *gin.Context) {
	id := ctx.Param("subscriptionId")

	err := s.Subscription.Delete(id)
	if err != nil {
		s.ServerError(ctx, err, id)
		return
	}
	s.Success(ctx, nil)
}

/*CreateOrder creates an order through Warehouse*/
func (s *Subscription) CreateOrder(ctx *gin.Context) {
	var json models.RequestOrder
	err := ctx.BindJSON(&json)
	if err != nil {
		s.UserError(ctx, "Error: unable to parse json", err)
		return
	}
	//check if subscription already exists
	sub, err := s.Subscription.GetByUserAndItem(json.UserID, json.ItemID)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}
	//TODO: pass itemID and add itemID to warehouse order struct
	order, err := s.Subscription.NewOrder(sub.UserID, sub.ID, sub.Quantity)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}
	s.Success(ctx, order)
}
