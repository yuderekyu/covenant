package router

import (
	"fmt"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/gateways"
	h "github.com/ghmeier/bloodlines/handlers"
	"github.com/yuderekyu/covenant/handlers"
)

type Subscription struct {
	router       *gin.Engine
	subscription handlers.SubscriptionI
}

func New(config *config.Root) (*Subscription, error) {
	sql, err := gateways.NewSQL(config.SQL)

	if err != nil {
		fmt.Println("ERROR: could not connect to mysql.")
		fmt.Println(err.Error())
		return nil, err
	}

	stats, err := statsd.New(
		statsd.Address(config.Statsd.Host+":"+config.Statsd.Port),
		statsd.Prefix(config.Statsd.Prefix),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx := &h.GatewayContext{
		Sql:   sql,
		Stats: stats,
	}

	s := &Subscription{
		subscription: handlers.NewSubscription(ctx),
	}

	InitRouter(s)
	return s, nil
}

func InitRouter(s *Subscription) {
	s.router = gin.Default()
	s.router.Use(h.GetCors())

	subscription := s.router.Group("/api/subscription")
	{
		subscription.POST("", s.subscription.New)
		subscription.GET("", s.subscription.ViewAll)
		subscription.GET("/:subscriptionId", s.subscription.View)
		subscription.GET("/:roasterId", s.subscription.ViewByRoaster)
		subscription.GET("/:userId", s.subscription.ViewByUser)
		subscription.POST("/:subscriptionId", s.subscription.Update)
		subscription.DELETE("/:subscriptionId", s.subscription.Delete)
	}
}

func (s *Subscription) Start(port string) {
	s.router.Run(port)
}
