package helpers

import (
	"database/sql"
	"fmt"
	"testing"
	// "strconv"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/yuderekyu/covenant/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdSuccess(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	userID := uuid.NewUUID()
	roasterID := uuid.NewUUID()
	itemID := uuid.NewUUID()

	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, startAt, roasterId, itemId FROM subscription").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), userID.String(), "ACTIVE", "01/01/11", "02/02/11", roasterID.String(), itemID.String()))

	subscription, err := s.GetByID(id.String())
	fmt.Println(subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(subscription.ID, id)
	assert.Equal(subscription.UserID, userID)
	assert.EqualValues(subscription.Status, models.ACTIVE)
	assert.Equal(subscription.CreatedAt, "01/01/11")
	assert.Equal(subscription.StartAt, "02/02/11")
	assert.Equal(subscription.RoasterID, roasterID)
	assert.Equal(subscription.ItemID, itemID)
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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, startAt, roasterId, itemId FROM subscription").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestInsertSuccess(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	subscription := getDefaultSubscription()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO subscription").
		ExpectExec().
		WithArgs(subscription.ID.String(), subscription.UserID.String(), string(models.ACTIVE), subscription.CreatedAt, subscription.StartAt, subscription.RoasterID.String(), subscription.ItemID.String()).
		WillReturnResult(sqlmock.NewResult(1,1))

	errTwo := s.Insert(subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo) 
}

func TestInsertFail(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	subscription :=getDefaultSubscription()
	s := getMockSubscription(db)
		if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO subscription").
	ExpectExec().
	WithArgs(subscription.ID.String(), subscription.UserID.String(), string(models.ACTIVE), subscription.CreatedAt, subscription.StartAt, subscription.RoasterID.String(), subscription.ItemID.String()).
	WillReturnError(fmt.Errorf("error"))

	errTwo := s.Insert(subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)

}

func TestUpdateSuccess(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	subscription := getDefaultSubscription()
	s := getMockSubscription(db)
		if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE subscription").
		ExpectExec().
		WithArgs(string(models.ACTIVE), subscription.StartAt, subscription.RoasterID.String(), subscription.ItemID.String(), subscription.ID.String()).
		WillReturnResult(sqlmock.NewResult(1,1))

	errTwo := s.Update(subscription.ID.String(), subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo)
}

func TestUpdateFail(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	subscription := getDefaultSubscription()
	s := getMockSubscription(db)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE subscription").
		ExpectExec().
		WithArgs(string(models.ACTIVE), subscription.StartAt, subscription.RoasterID.String(), subscription.ItemID.String(), subscription.ID.String()).
		WillReturnError(fmt.Errorf("error"))

	errTwo := s.Update(subscription.ID.String(), subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func getMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "userId", "status", "createdAt", "startAt", "roasterId", "itemId"})
} 

func getMockSubscription(s *sql.DB) *Subscription {
	return NewSubscription(&gateways.MySQL{DB: s})
}

func getDefaultSubscription() *models.Subscription {
	userID := uuid.NewUUID()
	roasterID := uuid.NewUUID()
	itemID := uuid.NewUUID()

	return models.NewSubscription(userID, "test", "test", roasterID, itemID)
}