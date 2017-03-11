package router

import (
	"testing"

	"github.com/ghmeier/bloodlines/config"
	"github.com/yuderekyu/covenant/handlers"

	bmocks "github.com/ghmeier/bloodlines/_mocks/gateways"
	ghandlers "github.com/ghmeier/bloodlines/handlers"
	tmocks "github.com/jakelong95/TownCenter/_mocks"
	lmocks "github.com/lcollin/warehouse/_mocks/gateways"
	cmocks "github.com/yuderekyu/covenant/_mocks/helpers"

	"github.com/stretchr/testify/assert"
	"gopkg.in/alexcesaro/statsd.v2"
)

func TestNewSucess(t *testing.T) {
	assert := assert.New(t)

	r, err := New(&config.Root{SQL: config.MySQL{}})

	assert.NoError(err)
	assert.NotNil(r)
}

func getMockCovenant() *Subscription {
	sql := new(bmocks.SQL)
	towncenter := new(tmocks.TownCenterI)
	warehouse := new(lmocks.Warehouse)
	stats, _ := statsd.New()
	ctx := &ghandlers.GatewayContext{
		Sql:        sql,
		TownCenter: towncenter,
		Warehouse:  warehouse,
		Stats:      stats,
	}
	return &Subscription{
		subscription: handlers.NewSubscription(ctx),
	}
}

func mockSubscription() (*Subscription, *cmocks.SubscriptionI) {
	s := getMockCovenant()
	mock := new(cmocks.SubscriptionI)
	s.subscription = &handlers.Subscription{Subscription: mock, BaseHandler: &ghandlers.BaseHandler{Stats: nil}}
	InitRouter(s)

	return s, mock
}
