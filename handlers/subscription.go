package handlers

import (
	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
	"github.com/yuderekyu/covenant/helpers"
	"github.com/yuderekyu/covenant/models"
)

type SubscriptionI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
	Cancel(ctx *gin.Context)
}

type Subscription struct {
	*handlers.BaseHandler
	Helper helpers.SubscriptionI
}

func NewSubscription(ctx *handlers.GatewayContext) SubscriptionI {
	return &Subscription{
		Helper:      helpers.NewSubscription(ctx.Sql),
		BaseHandler: &handlers.BaseHandler{Stats: ctx.Stats},
	}
}

func (s *Subscription) New(ctx *gin.Context) {
	var json models.Subscription
	err := ctx.BindJSON(&json)

	if err != nil {
		s.UserError(ctx, "Error: unable to parse json", err)
		return
	}

	subscription := models.NewSubscription(json.UserId, json.CreatedAt, json.StartAt, json.ShopId, json.OzInBag, json.BeanName, json.RoastName, json.Price)
	err = s.Helper.Insert(subscription)
	if err != nil {
		s.ServerError(ctx, err, json)
		return
	}

	s.Success(ctx, subscription)
}

func (s *Subscription) View(ctx *gin.Context) {
	id := ctx.Param("subscriptionId") //change

	subscription, err := s.Helper.GetById(id)
	if err != nil {
		s.ServerError(ctx, err, nil)
		return
	}

	s.Success(ctx, subscription)
}

func (s *Subscription) ViewAll(ctx *gin.Context) {
	offset, limit := s.GetPaging(ctx)
	subscriptions, err := s.Helper.GetAll(offset, limit)
	if err != nil {
		s.ServerError(ctx, err, nil)
		return
	}

	s.Success(ctx, subscriptions)
}

func (s *Subscription) Update(ctx *gin.Context) {
	id := ctx.Param("subscriptionId")

	var json models.Subscription
	err := ctx.BindJSON(&json)

	if err != nil {
		s.UserError(ctx, "Error: unable to parse json", err)
		return
	}
	json.Id = uuid.Parse(id)

	// TODO: call helper update method.

	// if err != nil {
	// 	s.ServerError(ctx, err, json)
	// 	return
	// }

	s.Success(ctx, json)
}

func (s *Subscription) Deactivate(ctx *gin.Context) {
	s.Success(ctx, nil)
}

func (s *Subscription) Cancel(ctx *gin.Context) {
	s.Success(ctx, nil)
}
