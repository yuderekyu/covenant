package router

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/yuderekyu/expresso-subscription/handlers"
	"github.com/yuderekyu/expresso-subscription/gateways"
)

type Subscription struct {
	router *gin.Engine
	subscription handlers.SubscriptionIfc
	// orders handlers.OrdersI
}

func New() (*Subscription, error) {
	sql, err := gateways.NewSql()

	if err != nil {
		fmt.Println("ERROR: could not connect to mysql.")
		fmt.Println(err.Error())
		return nil, err
	}

	s := &Subscription{
		subscription : handlers.NewSubscription(sql),
	}

	s.router = gin.Default()

	subscription := s.router.Group("/api/subscription")
	{
		subscription.POST("", s.subscription.New)
		subscription.GET("", s.subscription.View)
		subscription.POST("/:subscriptionId", s.subscription.Update)
		subscription.POST("/:subscriptionId/deactivate", s.subscription.Deactivate)
		subscription.DELETE("/:subscriptionId", s.subscription.Cancel)
	}

	// order := s.router.Group("/api/subscription/order")
	// {
	// 	order.POST("", s.)
	// }

	return s, nil
}

func (s *Subscription) Start(port string) {
	s.router.Run(port)
}