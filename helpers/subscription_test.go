package helpers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/yuderekyu/covenant/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdSuccess(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	userId := uuid.NewUUID()
	shopId := uuid.NewUUID()

	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, startAt, shopId, ozInBag, beanName, roastName, price FROM subscription").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), userId.String(), "ACTIVE", "01/01/11", "02/02/11", shopId.String(), 7.5, "arabica", "dark", 10.50))

	subscription, err := s.GetById(id.String())
	fmt.Println(subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(subscription.Id, id)
	assert.Equal(subscription.UserId, userId)
	assert.EqualValues(subscription.Status, models.ACTIVE)
	assert.Equal(subscription.CreatedAt, "01/01/11")
	assert.Equal(subscription.StartAt, "02/02/11")
	assert.Equal(subscription.ShopId, shopId)
	assert.Equal(subscription.OzInBag, float64(7.5))
	assert.Equal(subscription.BeanName, "arabica")
	assert.Equal(subscription.RoastName, "dark")
	assert.Equal(subscription.Price, float54(.50)
}

func TestGetByIdFail(t *testing.T) {
	assert := assert.New(t)
	id := uuid.NewUUID()

	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, startAt, shopId, ozInBag, beanName, roastName, price FROM subscription").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetById(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

// func TestGetByIdMapFail(t *testing.T) {
// 	assert :assert.New(t)

// 	id := uuid.NewUUID()
// 	db, mock, _ := sqlmock.New()
// 	s := getMockSubscription(db)

// 	mock.ExpectQuery("SELECT id, userId, status, createdAt, startAt, shopId, ozInBag, beanName, roastName, price FROM subscription").
// 	WithArgs(id.String()).WillReturnRows(getMockRows().AddRow(id.String(), ""))

// }

func getMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "userId", "status", "createdAt", "startAt", "shopId", "ozInBag", "beanName", "roastName", "price"})
} 

func getMockSubscription(s *sql.DB) *Subscription {
	return NewSubscription(&gateways.MySQL{DB: s})
}