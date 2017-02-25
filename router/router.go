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

/*Subscription is the main server object which routes requests*/
type Subscription struct {
	router       *gin.Engine
	subscription handlers.SubscriptionI
}

/*New returns a Subscription struct configured by the given config file*/
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

/*InitRouter connects the given endpoints to the router with gin*/
func InitRouter(s *Subscription) {
	s.router = gin.Default()
	s.router.Use(h.GetCors())

	subscription := s.router.Group("/api")
	{
		subscription.Use(s.subscription.Time())
		subscription.POST("/subscription", s.subscription.New)
		subscription.GET("/subscription", s.subscription.ViewAll)
		subscription.GET("/subscription/:subscriptionId", s.subscription.View)
		subscription.PUT("/subscription/:subscriptionId", s.subscription.Update)
		subscription.DELETE("/subscription/:subscriptionId", s.subscription.Delete)

		subscription.GET("/roaster/subscription/:roasterId", s.subscription.ViewByRoaster)
		subscription.GET("/user/subscription/:userId", s.subscription.ViewByUser)
	}
}

/*Start begins the Covenant server*/
func (s *Subscription) Start(port string) {
	s.router.Run(port)
}
