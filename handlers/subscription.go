package handlers

import(
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/pborman/uuid"

	"github.com/yuderekyu/covenant/models"
	"github.com/yuderekyu/covenant/gateways"
	"github.com/yuderekyu/covenant/helpers"
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
	Helper helpers.SubscriptionI
}

func NewSubscription(sql gateways.SQL) SubscriptionI {
	return &Subscription{Helper: helpers.NewSubscription(sql)}
}

func (s *Subscription) New(ctx *gin.Context) {
	var json models.Subscription
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Subscription Object"))
		fmt.Printf("%s", err.Error())
		return
	}

	subscription := models.NewSubscription(json.UserId, json.CreatedAt, json.StartAt, json.ShopId, json.OzInBag, json.BeanName, json.RoastName, json.Price)
	err = s.Helper.Insert(subscription)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": subscription})
}

func (s *Subscription) View(ctx *gin.Context) {
	id := ctx.Param("subscriptionId") //change 

	subscription, err := s.Helper.GetById(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return 
	}

	ctx.JSON(200, gin.H{"data": subscription})
}

func (s *Subscription) ViewAll(ctx *gin.Context) {
	offset, limit := getPaging(ctx)
	subscription, err := s.Helper.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": subscription})
}

func (s *Subscription) Update(ctx *gin.Context) {
	id := ctx.Param("subscriptionId")

	var json models.Subscription
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse(err.Error()))
		return
	}
	json.Id = uuid.Parse(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": json})
}

func (s *Subscription) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (s *Subscription) Cancel(ctx *gin.Context) {
	ctx.JSON(200, empty())
}