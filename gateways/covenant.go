package gateways

import (
	"fmt"
	"net/http"

	"github.com/pborman/uuid"

	"github.com/yuderekyu/covenant/config"
	"github.com/yuderekyu/covenant/models"
	"github.com/ghmeier/bloodlines/gateways"
)

/*Covenant wraps all methods of the covenant API*/
type Covenant interface {
	NewSubscription(newSubscription *models.Subscription) (*models.Subscription, error) 
	GetAllSubscription(offset int, limit int) ([]*models.Subscription, error)
	GetSubscriptionById(id uuid.UUID) (*models.Subscription, error)
	UpdateSubscription(id uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(id uuid.UUID) (*models.Subscription, error)
}

type covenant struct {
	*BaseService
	url string
	client *http.Client
}

func NewCovenant(config config.Covenant) Covenant{
	return &covenant{
		BaseService: NewBaseService(),
		url: fmt.Sprintf("https://%s:%s/api/", config.Host, config.Port),
	}
}

func (c *covenant) NewSubscription(newSubscription, *models.Subscription) (*models.Subscription, error){
	url := ftmt.Sprintf("%ssubscription", c.url)

	var subscription *models.Subscription
	err := c.ServiceSend(http.MethodPost, url, newSubscription, subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func(c *covenant) GetAllSubscription(offset int, limit int)([]*models.Subscription, error){

}

func(c *covenant) GetSubscriptionById(id uuid.UUID) (*models.Subscription, error){

}

func(c *covenant) UpdateSubscription(id uuid.UUID) (*models.Subscription, error){

}

func(c *covenant) DeleteSubscription(id uuid.UUID) (*models.Subscription, error){
	
}