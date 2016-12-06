package handlers

import(
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/pborman/uuid"

	"github.com/yuderekyu/expresso-subscription/containers"
	"github.com/yuderekyu/expresso-subscription/gateways"
)

type SubscriptionIfc interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
	Cancel(ctx *gin.Context)
}

type Subscription struct {
	sql *gateways.Sql
}

func NewSubscription(sql *gateways.Sql) SubscriptionIfc {
	return &Subscription{sql: sql}
}

func (s *Subscription) New(ctx *gin.Context) {
	var json containers.Subscription
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(400, errResponse("Invalid Subscription Object"))
		fmt.Printf("%s", err.Error())
		return
	}

	//todo add in reference to orders params?
	err = s.sql.Modify(
		"INSERT INTO subscription VALUE(?, ?, ?, ?, ?, ?, ?, ?)",
		uuid.New(),
		json.OrderId,
		json.Type,
		json.UserId,
		json.Status,
		json.CreatedAt,
		json.StartAt,
		json.TotalPrice)

	if err != nil {
		ctx.JSON(500, &gin.H{"error": err, "message": err.Error()})
		return
	}
	ctx.JSON(200, empty())
}

func (s *Subscription) ViewAll(ctx *gin.Context) {
	rows, err := s.sql.Select("SELECT * FROM subscription")
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	subscription, err := containers.FromSql(rows)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": subscription})
}

func (s *Subscription) View(ctx *gin.Context) {
	id := ctx.Param("subscriptionId")
	if id == "" {
		ctx.JSON(500, errResponse("subscriptionId is a required parameter"))
		return
	}

	rows, err := s.sql.Select("SELECT * FROM subscription WHERE id=?")
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	subscription, err := containers.FromSql(rows)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": subscription})
}

func (s *Subscription) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (s *Subscription) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (s *Subscription) Cancel(ctx *gin.Context) {
	ctx.JSON(200, empty())
}