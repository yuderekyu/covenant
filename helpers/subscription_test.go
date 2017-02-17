package helpers

import (
	"database/sql"
	"fmt"
	"testing"
	// "strconv"

	"github.com/ghmeier/bloodlines/gateways"
	t"github.com/jakelong95/TownCenter/gateways"
	g"github.com/lcollin/warehouse/gateways"
	tmocks "github.com/jakelong95/TownCenter/_mocks"
	//lmocks "github.com/lcollin/warehouse"
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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
		WithArgs(id.String()).
		WillReturnRows(getMockRows().AddRow(id.String(), userID.String(), "ACTIVE", "01/01/11", "MONTHLY", roasterID.String(), itemID.String()))

	subscription, err := s.GetByID(id.String())
	fmt.Println(subscription)
	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(subscription.ID, id)
	assert.Equal(subscription.UserID, userID)
	assert.EqualValues(subscription.Status, models.ACTIVE)
	assert.Equal(subscription.CreatedAt, "01/01/11")
	assert.Equal(subscription.Frequency, "MONTHLY")
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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestGetAllSuccess(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 10
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
		WithArgs(offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), uuid.New(), "PENDING", "01/01/17", "MONTHLY", uuid.New(), uuid.New()).
			AddRow(uuid.New(), uuid.New(), "ACTIVE", "01/10/17", "MONTHLY", uuid.New(), uuid.New()).
			AddRow(uuid.New(), uuid.New(), "CANCELLED", "01/10/17", "MONTHLY", uuid.New(), uuid.New()))

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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("error"))

	_, errTwo := s.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(errTwo)
}

func TestGetByRoasterSuccess(t *testing.T) {
	assert := assert.New(t)

	subscription := getDefaultSubscription()
	offset, limit := 0, 10
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
		WithArgs(subscription.RoasterID, offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), uuid.New(), "PENDING", "01/01/17", "MONTHLY", subscription.RoasterID.String(), uuid.New()).
			AddRow(uuid.New(), uuid.New(), "ACTIVE", "01/10/17", "MONTHLY", subscription.RoasterID.String(), uuid.New()))

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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
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
	db, mock, err := sqlmock.New()
	s := getMockSubscription(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
		WithArgs(subscription.UserID, offset, limit).
		WillReturnRows(getMockRows().
			AddRow(uuid.New(), subscription.UserID.String(), "PENDING", "01/01/17", "MONTHLY", uuid.New(), uuid.New()).
			AddRow(uuid.New(), subscription.UserID.String(), "ACTIVE", "01/10/17", "MONTHLY", uuid.New(), uuid.New()))

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

	mock.ExpectQuery("SELECT id, userId, status, createdAt, frequency, roasterId, itemId FROM subscription").
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
		WithArgs(subscription.ID.String(), subscription.UserID.String(), string(models.ACTIVE), subscription.CreatedAt, subscription.Frequency, subscription.RoasterID.String(), subscription.ItemID.String()).
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
	WithArgs(subscription.ID.String(), subscription.UserID.String(), string(models.ACTIVE), subscription.CreatedAt, subscription.Frequency, subscription.RoasterID.String(), subscription.ItemID.String()).
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
		WithArgs(string(models.ACTIVE), subscription.Frequency, subscription.RoasterID.String(), subscription.ItemID.String(), subscription.ID.String()).
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
		WithArgs(string(models.ACTIVE), subscription.Frequency, subscription.RoasterID.String(), subscription.ItemID.String(), subscription.ID.String()).
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
	return sqlmock.NewRows([]string{"id", "userId", "status", "createdAt", "frequency", "roasterId", "itemId"})
} 

func getMockSubscription(s *sql.DB) *Subscription {
	return NewSubscription(&gateways.MySQL{DB: s}, &t.TownCenter, &g.Warehouse)
}

func getDefaultSubscription() *models.Subscription {
	userID := uuid.NewUUID()
	roasterID := uuid.NewUUID()
	itemID := uuid.NewUUID()

	return models.NewSubscription(userID, "test", roasterID, itemID)
}