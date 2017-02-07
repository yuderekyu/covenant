package gateways

import (
	"fmt"
	"net/http"

	"github.com/pborman/uuid"

	g "github.com/ghmeier/bloodlines/gateways"
	"github.com/yuderekyu/covenant/config"
	"github.com/yuderekyu/covenant/models"
)

/*Covenant wraps all methods of the covenant API*/
type Covenant interface {
	NewSubscription(newSubscription *models.Subscription) (*models.Subscription, error) 
	GetAllSubscription(offset int, limit int) ([]*models.Subscription, error)
	GetSubscriptionById(id uuid.UUID) (*models.Subscription, error)
	UpdateSubscription(id uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(id uuid.UUID) (error)
}

type covenant struct {
	*g.BaseService
	url string
	client *http.Client
}

func NewCovenant(config config.Covenant) Covenant{
	return &covenant{
		BaseService: g.NewBaseService(),
		url: fmt.Sprintf("https://%s:%s/api/", config.Host, config.Port),
	}
}

func (c *covenant) NewSubscription(newSubscription *models.Subscription) (*models.Subscription, error){
	url := fmt.Sprintf("%ssubscription", c.url)

	var subscription *models.Subscription
	err := c.ServiceSend(http.MethodPost, url, newSubscription, subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func(c *covenant) GetAllSubscription(offset int, limit int)([]*models.Subscription, error){
	url := fmt.Sprintf("%ssubscription?offset=%d&limit=%d", c.url, offset, limit)
	
	var subscription []*models.Subscription
	err := c.ServiceSend(http.MethodGet, url, nil, subscription)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func(c *covenant) GetSubscriptionById(id uuid.UUID) (*models.Subscription, error){
	url := fmt.Sprintf("%ssubscription/%s", c.url, id)

	var subscription *models.Subscription
	err := c.ServiceSend(http.MethodGet, url, nil, subscription)
	if err!= nil{
		return nil, err
	}
	return subscription, nil
}

func(c *covenant) UpdateSubscription(id uuid.UUID) (*models.Subscription, error){
	url := fmt.Sprintf("%ssubscription%s", c.url, id)
	
	var subscription *models.Subscription
	err := c.ServiceSend(http.MethodPost, url, id, subscription)
	if(err != nil){
		return nil, err
	}
	return nil, nil
}

func(c *covenant) DeleteSubscription(id uuid.UUID) error{
	url := fmt.Sprintf("%ssubscription%s", c.url, id)

	err := c.ServiceSend(http.MethodDelete, url, nil, nil)
	if(err != nil){
		return err
	}
	return nil
}