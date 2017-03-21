package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yuderekyu/covenant/models"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestSubscriptionNewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	covenant, sMock := mockSubscription()
	sMock.On("Insert", mock.AnythingOfType("*models.Subscription")).Return(nil)

	userMock := uuid.NewUUID()
	roasterMock := uuid.NewUUID()
	itemMock := uuid.NewUUID()

	s := getSubscriptionString(models.NewSubscription(userMock, "FREQUENCY", roasterMock, itemMock))

	/*records mutations for inspection*/
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/subscription", s)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestSubscriptionNewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	userMock := uuid.NewUUID()
	roasterMock := uuid.NewUUID()
	itemMock := uuid.NewUUID()

	covenant, sMock := mockSubscription()
	sMock.On("Insert", mock.AnythingOfType("*models.Subscription")).Return(fmt.Errorf("error"))

	s := getSubscriptionString(models.NewSubscription(userMock, "test", roasterMock, itemMock))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/subscription", s)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestSubscriptionViewSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	covenant, sMock := mockSubscription()
	sMock.On("GetByID", id.String()).Return(&models.Subscription{}, nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/subscription/"+id.String(), nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestSubscriptionViewFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	id := uuid.NewUUID()
	covenant, sMock := mockSubscription()
	sMock.On("GetByID", id.String()).Return(nil, fmt.Errorf("error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/subscription/"+id.String(), nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestContentViewAllSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	covenant, sMock := mockSubscription()
	sMock.On("GetAll", 0, 20).Return(make([]*models.Subscription, 0), nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/subscription", nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestContentViewAllFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	covenant, sMock := mockSubscription()
	sMock.On("GetAll", 0, 20).Return(make([]*models.Subscription, 0), fmt.Errorf("error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/subscription", nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestSubscriptionGetByRoasterSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	roasterID := uuid.NewUUID()

	covenant, sMock := mockSubscription()
	sMock.On("GetByRoaster", roasterID.String(), 0, 4).Return(make([]*models.Subscription, 0), nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster/subscription/"+roasterID.String()+"?offset=0&limit=4", nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestSubscriptionGetByRoasterFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	roasterID := uuid.NewUUID()

	covenant, sMock := mockSubscription()
	sMock.On("GetByRoaster", roasterID.String(), 0, 4).Return(make([]*models.Subscription, 0), fmt.Errorf("error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/roaster/subscription/"+roasterID.String()+"?offset=0&limit=4", nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func TestSubscriptionGetByUserSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	userID := uuid.NewUUID()

	covenant, sMock := mockSubscription()
	sMock.On("GetByUser", userID.String(), 0, 4).Return(make([]*models.Subscription, 0), nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/user/subscription/"+userID.String()+"?offset=0&limit=4", nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestSubscriptionGetByUserFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	userID := uuid.NewUUID()

	covenant, sMock := mockSubscription()
	sMock.On("GetByUser", userID.String(), 0, 4).Return(make([]*models.Subscription, 0), fmt.Errorf("error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/user/subscription/"+userID.String()+"?offset=0&limit=4", nil)
	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

// func TestSubscriptionUpdateSuccess(t *testing.T) {
// 	assert := assert.New(t)

// 	gin.SetMode(gin.TestMode)
// 	userID := uuid.NewUUID()
// 	roasterID := uuid.NewUUID()
// 	itemID := uuid.NewUUID()

// 	subscription := models.NewSubscription(userID, "FREQUENCY", roasterID, itemID)
// 	id := subscription.ID
// 	fmt.Printf("\nCreatedAt: %s\n", subscription.CreatedAt)
// 	covenant, sMock := mockSubscription()
// 	sMock.On("Update", id.String(), subscription).Return(nil)

// 	recorder := httptest.NewRecorder()
// 	request, _ := http.NewRequest("PUT", "/api/subscription/"+subscription.ID.String(),
// 		getSubscriptionString(subscription))
// 	covenant.router.ServeHTTP(recorder, request)

// 	assert.Equal(200, recorder.Code)
// }

// func TestSubscriptionUpdateFail(t *testing.T) {
// 	assert := assert.New(t)

// 	gin.SetMode(gin.TestMode)
// 	userID := uuid.NewUUID()
// 	roasterID := uuid.NewUUID()
// 	itemID := uuid.NewUUID()

// 	subscription := models.NewSubscription(userID, "FREQUENCY", roasterID, itemID)
// 	id := subscription.ID
// 	fmt.Printf("\nCreatedAt: %s\n", subscription.CreatedAt)
// 	covenant, sMock := mockSubscription()
// 	sMock.On("Update", id.String(), subscription).Return(fmt.Errorf("error"))

// 	recorder := httptest.NewRecorder()
// 	request, _ := http.NewRequest("PUT", "/api/subscription/"+subscription.ID.String(),
// 		getSubscriptionString(subscription))

// 	covenant.router.ServeHTTP(recorder, request)

// 	assert.Equal(500, recorder.Code)
// }

func TestSubscriptionDeleteSuccess(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	userID := uuid.NewUUID()
	roasterID := uuid.NewUUID()
	itemID := uuid.NewUUID()

	subscription := models.NewSubscription(userID, "FREQUENCY", roasterID, itemID)

	covenant, sMock := mockSubscription()
	sMock.On("Delete", subscription.ID.String()).Return(nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/subscription/"+subscription.ID.String(), nil)

	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(200, recorder.Code)
}

func TestUserDeleteFail(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	userID := uuid.NewUUID()
	roasterID := uuid.NewUUID()
	itemID := uuid.NewUUID()

	subscription := models.NewSubscription(userID, "01/01/01", roasterID, itemID)

	covenant, sMock := mockSubscription()
	sMock.On("Delete", subscription.ID.String()).Return(fmt.Errorf("error"))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/subscription/"+subscription.ID.String(), nil)

	covenant.router.ServeHTTP(recorder, request)

	assert.Equal(500, recorder.Code)
}

func getSubscriptionString(m *models.Subscription) io.Reader {
	s, _ := json.Marshal(m) //convert to json
	return bytes.NewReader(s)
}
