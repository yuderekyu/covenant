package gateways

import (
	"fmt"
	"net/http"

	"github.com/ghmeier/bloodlines/config"
	g "github.com/ghmeier/bloodlines/gateways"
	"github.com/jakelong95/TownCenter/models"

	"github.com/pborman/uuid"
)

/*TownCenterI describes the functions for interacting with town center*/
type TownCenterI interface {
	GetUser(uuid.UUID) (*models.User, error)
	GetAllUsers(int, int) ([]*models.User, error)
	UpdateUser(uuid.UUID, *models.User) error
}

/*TownCenter contains instrumentation for accessing TownCenter service*/
type TownCenter struct {
	*g.BaseService
	host   string
	port   string
	url    string
	client *http.Client
}

/*NewTownCenter creates and returns a TownCenter gateway*/
func NewTownCenter(config config.TownCenter) TownCenterI {
	return &TownCenter{
		BaseService: g.NewBaseService(),
		host:	config.Host,
		port:	config.Port,
		url:    fmt.Sprintf("https://%s:%s/api/", config.Host, config.Port),
		client: &http.Client{},
	}
}

/*GetUser  gets information about a user based on the user ID*/
func (t *TownCenter) GetUser(id uuid.UUID) (*models.User, error) {
	url := fmt.Sprintf("%suser/%s", t.url, id.String())

	var user *models.User
	err := t.ServiceSend(http.MethodGet, url, nil, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

/*GetAllUsers gets information about all the users, paginated with an offset and limit per page*/
func (t *TownCenter) GetAllUsers(offset, limit int) ([]*models.User, error) {
	url := fmt.Sprintf("%suser?offset=%d&limit=%d", t.url, offset, limit)

	var users []*models.User
	err := t.ServiceSend(http.MethodGet, url, nil, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

/*UpdateUser updates the information about a user based on user id*/
func (t *TownCenter) UpdateUser(id uuid.UUID, user *models.User) error {
	url := fmt.Sprintf("%suser/%s", t.url, id.String())
	return t.ServiceSend(http.MethodPut, url, user, nil)
}

/*GetRoaster gets information about a roaster based on the roaster ID*/
func (t *TownCenter) GetRoaster(id uuid.UUID) (*models.Roaster, error) {
	url := fmt.Sprintf("%sroaster/%s", t.url, id.String())

	var roaster *models.Roaster
	err := t.ServiceSend(http.MethodGet, url, nil, &roaster)
	if err != nil {
		return nil, err
	}

	return roaster, nil
}

/*GetAllRoasters gets information about all the roasters, paginated with an offset and limit per page*/
func (t *TownCenter) GetAllRoasters(offset, limit int) ([]*models.Roaster, error) {
	url := fmt.Sprintf("%sroaster?offset=%d&limit=%d", t.url, offset, limit)

	var roasters []*models.Roaster
	err := t.ServiceSend(http.MethodGet, url, nil, &roasters)
	if err != nil {
		return nil, err
	}

	return roasters, nil
}

/*UpdateRoaster updates the information about a roaster based on roaster id*/
func (t *TownCenter) UpdateRoaster(id uuid.UUID, roaster *models.Roaster) error {
	url := fmt.Sprintf("%sroaster/%s", t.url, id.String())
	return t.ServiceSend(http.MethodPut, url, roaster, nil)
}
