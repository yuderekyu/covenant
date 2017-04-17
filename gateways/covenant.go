package gateways

import (
	"fmt"
	"net/http"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/config"
	g "github.com/ghmeier/bloodlines/gateways"
	wareM "github.com/lcollin/warehouse/models"
	"github.com/yuderekyu/covenant/models"
)

/*Covenant wraps all methods of the covenant API*/
type Covenant interface {
	NewSubscription(request *models.Subscription) (*models.Subscription, error)
	GetAllSubscription(offset int, limit int) ([]*models.Subscription, error)
	GetSubscriptionById(id uuid.UUID) (*models.Subscription, error)
	GetSubscriptionByRoaster(roasterID uuid.UUID, offset int, limit int) (*models.Subscription, error)
	GetSubscriptionByUser(userID uuid.UUID, offset int, limit int) (*models.Subscription, error)
	UpdateSubscription(id uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(id uuid.UUID) error
	NewOrder(request *models.RequestOrder) (*wareM.Order, error)
}

/*covenant is the structure for the Covenant service*/
type covenant struct {
	*g.BaseService
	url    string
	client *http.Client
}

/*NewCovenant creates and returns a Covenant struct which points to the service denoted in the config*/
func NewCovenant(config config.Covenant) Covenant {
	var url string
	if config.Port != "" {
		url = fmt.Sprintf("http://%s:%s/api/", config.Host, config.Port)
	} else {
		url = fmt.Sprintf("https://%s/api/", config.Host)
	}

	return &covenant{
		BaseService: g.NewBaseService(),
		url:         url,
	}
}

/*NewSubscription creates a new subscription*/
func (c *covenant) NewSubscription(request *models.Subscription) (*models.Subscription, error) {
	url := fmt.Sprintf("%ssubscription", c.url)

	var subscription models.Subscription
	err := c.ServiceSend(http.MethodPost, url, request, &subscription)
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

/*GetAllSubscription returns a list of subscription with the offset and limit determining the entries and amount*/
func (c *covenant) GetAllSubscription(offset int, limit int) ([]*models.Subscription, error) {
	url := fmt.Sprintf("%ssubscription?offset=%d&limit=%d", c.url, offset, limit)

	subscription := make([]*models.Subscription, 0)
	err := c.ServiceSend(http.MethodGet, url, nil, &subscription)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

/*GetSubscriptionById returns a subscription with the given id*/
func (c *covenant) GetSubscriptionById(id uuid.UUID) (*models.Subscription, error) {
	url := fmt.Sprintf("%ssubscription/%s", c.url, id.String())

	var subscription models.Subscription
	err := c.ServiceSend(http.MethodGet, url, nil, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

/*GetSubscriptionByRoaster returns a list of subscription of the given roasterID, with the offset and limit determining the entries and amount*/
func (c *covenant) GetSubscriptionByRoaster(roasterID uuid.UUID, offset int, limit int) (*models.Subscription, error) {
	url := fmt.Sprintf("%sroaster/subscription/%s?offset=%d&limit=%d", c.url, roasterID.String(), offset, limit)

	var subscription models.Subscription
	err := c.ServiceSend(http.MethodGet, url, nil, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

/*GetSubscriptionByRoaster returns a list of subscription of the given userID, with the offset and limit determining the entries and amount*/
func (c *covenant) GetSubscriptionByUser(userID uuid.UUID, offset int, limit int) (*models.Subscription, error) {
	url := fmt.Sprintf("%suser/subscription/%s?offset=%d&limit=%d", c.url, userID.String(), offset, limit)

	var subscription models.Subscription
	err := c.ServiceSend(http.MethodGet, url, nil, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

/*UpdateSubscription overwrites a subscription by the given id */
func (c *covenant) UpdateSubscription(id uuid.UUID) (*models.Subscription, error) {
	url := fmt.Sprintf("%ssubscription%s", c.url, id)

	var subscription models.Subscription
	err := c.ServiceSend(http.MethodPost, url, id, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

/*DeleteSubscription removes a subscription by the given id*/
func (c *covenant) DeleteSubscription(id uuid.UUID) error {
	url := fmt.Sprintf("%ssubscription%s", c.url, id)

	err := c.ServiceSend(http.MethodDelete, url, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *covenant) NewOrder(request *models.RequestOrder) (*wareM.Order, error) {
	url := fmt.Sprintf("%sorder", c.url)
	var order wareM.Order

	err := c.ServiceSend(http.MethodPost, url, request, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
