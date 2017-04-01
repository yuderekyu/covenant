package helpers

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ghmeier/bloodlines/gateways"
	tmocks "github.com/jakelong95/TownCenter/_mocks"
	wmocks "github.com/lcollin/warehouse/_mocks/gateways"
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
	time := time.Now()
	quantity := 1;

	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), userID.String(), "ACTIVE", time, "MONTHLY", roasterID.String(), itemID.String(), quantity))

	subscription, err := s.GetByID(id.String())
	fmt.Println(subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(subscription.ID, id)
	assert.Equal(subscription.UserID, userID)
	assert.EqualValues(subscription.Status, models.ACTIVE)
	assert.Equal(subscription.CreatedAt, time)
	assert.Equal(subscription.Frequency, "MONTHLY")
	assert.Equal(subscription.RoasterID, roasterID)
	assert.Equal(subscription.ItemID, itemID)
	assert.Equal(subscription.Quantity, quantity)
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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestGetAllSuccess(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 10
	time := time.Now()
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), uuid.New(), "PENDING", time, "MONTHLY", uuid.New(), uuid.New(), 1).
			AddRow(uuid.New(), uuid.New(), "ACTIVE", time, "MONTHLY", uuid.New(), uuid.New(), 1).
			AddRow(uuid.New(), uuid.New(), "CANCELLED", time, "MONTHLY", uuid.New(), uuid.New(), 1))

	subscriptions, errTwo := s.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo)
	assert.Equal(3, len(subscriptions))
}

func TestGetAllFail(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 10
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestGetByRoasterSuccess(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	time := time.Now()
	offset, limit := 0, 10
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(subscription.RoasterID, offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), uuid.New(), "PENDING", time, "MONTHLY", subscription.RoasterID.String(), uuid.New(), 1).
			AddRow(uuid.New(), uuid.New(), "ACTIVE", time, "MONTHLY", subscription.RoasterID.String(), uuid.New(), 1))

	subscriptions, errTwo := s.GetByRoaster(subscription.RoasterID.String(), offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo)
	assert.Equal(2, len(subscriptions))
}

func TestGetByRoasterFail(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	offset, limit := 0, 10
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(subscription.RoasterID, offset, limit).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetByRoaster(subscription.RoasterID.String(), offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestGetByUserSuccess(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	offset, limit := 0, 10
	time := time.Now()
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(subscription.UserID, offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), subscription.UserID.String(), "PENDING", time, "MONTHLY", uuid.New(), uuid.New(), 1).
			AddRow(uuid.New(), subscription.UserID.String(), "ACTIVE", time, "MONTHLY", uuid.New(), uuid.New(), 1))

	subscriptions, errTwo := s.GetByUser(subscription.UserID.String(), offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo)
	assert.Equal(2, len(subscriptions))
}

func TestGetByUserFail(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	offset, limit := 0, 10
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId, quantity FROM subscription").
		WithArgs(subscription.UserID, offset, limit).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetByRoaster(subscription.UserID.String(), offset, limit)

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
		WithArgs(subscription.ID.String(), subscription.UserID.String(), string(models.ACTIVE), 
		subscription.CreatedAt, subscription.Frequency, subscription.RoasterID.String(), subscription.ItemID.String(), subscription.Quantity).
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
	WithArgs(subscription.ID.String(), subscription.UserID.String(), string(models.ACTIVE), 
	subscription.CreatedAt, subscription.Frequency, subscription.RoasterID.String(), subscription.ItemID.String(), subscription.Quantity).
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
		WithArgs(string(models.ACTIVE), subscription.Frequency, subscription.RoasterID.String(), 
		subscription.ItemID.String(), subscription.Quantity, subscription.ID.String()).
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
		WithArgs(string(models.ACTIVE), subscription.Frequency, subscription.RoasterID.String(), 
		subscription.ItemID.String(), subscription.Quantity, subscription.ID.String()).
		WillReturnError(fmt.Errorf("error"))

	errTwo := s.Update(subscription.ID.String(), subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestDeleteSuccess(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	db, mock, err :=sqlmock.New()
	s:= getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM subscription").
		ExpectExec().
		WithArgs(id.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errTwo := s.Delete(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo)
}

func TestDeleteFail(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	db, mock, err :=sqlmock.New()
	s:= getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM subscription").
		ExpectExec().
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("error"))

	errTwo := s.Delete(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestSetStatusSuccess(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	db, mock, err :=sqlmock.New()
	s:= getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE subscription").
		ExpectExec().
		WithArgs(string(subscription.Status), subscription.ID).
		WillReturnResult(sqlmock.NewResult(1,1))

	errTwo := s.SetStatus(subscription.ID.String(), subscription.Status)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(errTwo)
}

func TestSetStatusFail(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	db, mock, err :=sqlmock.New()
	s:= getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE subscription").
		ExpectExec().
		WithArgs(string(subscription.Status), subscription.ID).
		WillReturnError(fmt.Errorf("error"))

	errTwo := s.SetStatus(subscription.ID.String(), subscription.Status)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func getMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "userId", "status", "createdAt", "frequency", "roasterId", "itemId", "quantity"})
} 

type mockContext struct {
	tc *tmocks.TownCenterI
	w *wmocks.Warehouse
}

func getMockSubscription(s *sql.DB) *Subscription {
	mocks :=&mockContext{
		tc: &tmocks.TownCenterI{},
		w: &wmocks.Warehouse{},
	}
	return NewSubscription(&gateways.MySQL{DB: s}, mocks.tc, mocks.w)
}

func getDefaultSubscription() *models.Subscription {
	userID := uuid.NewUUID()
	roasterID := uuid.NewUUID()
	itemID := uuid.NewUUID() 
	quantity := 1

	return models.NewSubscription(userID, "Frequency", roasterID, itemID, quantity)
}