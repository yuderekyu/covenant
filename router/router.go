package router

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/yuderekyu/expresso-subscription/config"
	"github.com/yuderekyu/expresso-subscription/handlers"
	"github.com/yuderekyu/expresso-subscription/gateways"
)

type Subscription struct {
	router *gin.Engine
	subscription handlers.SubscriptionI
}

func New(config *config.Root) (*Subscription, error) {
	sql, err := gateways.NewSQL(config.SQL)

	if err != nil {
		fmt.Println("ERROR: could not connect to mysql.")
		fmt.Println(err.Error())
		return nil, err
	}

	s := &Subscription{
		subscription : handlers.NewSubscription(sql),
	}

	InitRouter(s)
	return s, nil
}

func InitRouter(s *Subscription) {
	s.router = gin.Default()

	subscription := s.router.Group("/api/subscription")
	{
		subscription.POST("", s.subscription.New)
		subscription.GET("/:subscriptionId", s.subscription.View)
		subscription.POST("/:subscriptionId", s.subscription.Update)
		subscription.POST("/:subscriptionId/deactivate", s.subscription.Deactivate)
		subscription.DELETE("/:subscriptionId", s.subscription.Cancel)
	}
}

func (s *Subscription) Start(port string) {
	s.router.Run(port)
}