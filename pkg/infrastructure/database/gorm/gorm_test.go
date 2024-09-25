package gorm

import (
	"context"
	"errors"
	"testing"

	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	gt "github.com/JhonX2011/GFAWBP/test/generic"
	mocks "github.com/JhonX2011/GFAWBP/test/mocks/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type gormConnectionsScenery struct {
	gormClient   IClientGorm
	db           *gorm.DB
	config       *mic.DBConnection
	sqlDialector gorm.Dialector
	mockSQL      *mocks.MockSQLForGormConnection
	callback     func()
	gt.GenericTest
}

func givenGormConnectionsScenery(t *testing.T) *gormConnectionsScenery {
	t.Parallel()
	return &gormConnectionsScenery{
		config: &mic.DBConnection{
			MaxRetries:        3,
			RetryIntervalTime: 10,
			LogQueries:        true,
		},
		mockSQL: &mocks.MockSQLForGormConnection{},
	}
}

func (s *gormConnectionsScenery) givenSQLDialector(sqlDialector gorm.Dialector) {
	s.sqlDialector = sqlDialector
}

func (s *gormConnectionsScenery) givenDB(db *gorm.DB) {
	s.db = db
}

func (s *gormConnectionsScenery) givenGormDB() {
	s.gormClient = &gormClient{
		db:     s.db,
		config: s.config,
	}
}

func (s *gormConnectionsScenery) whenNewGormClientIsCall() {
	s.AResult, s.AError = NewGormClient(s.sqlDialector, s.config)
}

func (s *gormConnectionsScenery) whenGetDBIsCall(ctx context.Context) {
	s.AResult = s.gormClient.GetDB(ctx)
}

func (s *gormConnectionsScenery) whenRetryQueryIsCall(ctx context.Context, err error) {
	s.AResult = s.gormClient.RetryQuery(ctx, func() *gorm.DB {
		return &gorm.DB{Error: err}
	})
}

func (s *gormConnectionsScenery) whenBeginIsCall(ctx context.Context) {
	s.AResult, s.callback, s.AError = s.gormClient.Begin(ctx)
	s.callback()
}

func (s *gormConnectionsScenery) whenCommitIsCall(ctx context.Context) {
	s.AError = s.gormClient.Commit(ctx)
}

func (s *gormConnectionsScenery) whenRollbackIsCall(ctx context.Context) {
	s.AError = s.gormClient.Rollback(ctx)
}

func (s *gormConnectionsScenery) thenHaveExpectedError(t *testing.T, expectedError error) {
	assert.Equal(t, s.AResult.(*gorm.DB).Error, expectedError)
}

func TestNewGormClientOK(t *testing.T) {
	db, _, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	test := givenGormConnectionsScenery(t)
	test.givenSQLDialector(mocks.GetDialector(db))
	test.whenNewGormClientIsCall()
	test.ThenNoHaveError(t)
}

func TestNewGormClientFail(t *testing.T) {
	db, _, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	closeDB()

	test := givenGormConnectionsScenery(t)
	test.givenSQLDialector(mocks.GetDialector(db))
	test.whenNewGormClientIsCall()
	test.ThenHaveError(t)
}

func TestRetryQueryWithContextCancelled(t *testing.T) {
	test := givenGormConnectionsScenery(t)
	test.givenGormDB()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	test.whenRetryQueryIsCall(ctx, errors.New("context canceled"))
	test.thenHaveExpectedError(t, errors.New("context canceled"))
}

func TestRetryQueryWithError(t *testing.T) {
	test := givenGormConnectionsScenery(t)
	test.givenGormDB()
	ctx := context.Background()
	test.whenRetryQueryIsCall(ctx, errors.New("simulated error"))
	test.thenHaveExpectedError(t, errors.New("simulated error"))
}

func TestRetryQueryWithOK(t *testing.T) {
	test := givenGormConnectionsScenery(t)
	test.givenGormDB()
	ctx := context.Background()
	test.whenRetryQueryIsCall(ctx, nil)
	test.ThenNoHaveError(t)
}

func TestBeginOk(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.Background()
	test.whenBeginIsCall(ctx)
	test.ThenNoHaveError(t)
}

func TestBeginFail(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin().WillReturnError(gorm.ErrInvalidTransaction)

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.Background()
	test.whenBeginIsCall(ctx)
	test.ThenHaveError(t)
}

func TestBeginOkWithTransaction(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()
	database := gormDB.Begin()
	if database.Error != nil {
		return
	}

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.WithValue(context.Background(), transactionCtx{}, &transactionDB{db: database})
	test.whenBeginIsCall(ctx)
	test.ThenHaveError(t)
}

func TestGetDBOk(t *testing.T) {
	db, _, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.Background()
	test.whenGetDBIsCall(ctx)
	test.ThenNoHaveError(t)
}

func TestGetDBOkWithTransaction(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()
	database := gormDB.Begin()
	if database.Error != nil {
		return
	}

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.WithValue(context.Background(), transactionCtx{}, &transactionDB{db: database})
	test.whenGetDBIsCall(ctx)
	test.ThenNoHaveError(t)
}

func TestCommitOk(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()
	database := gormDB.Begin()
	if database.Error != nil {
		return
	}

	dbMock.ExpectCommit()

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.WithValue(context.Background(), transactionCtx{}, &transactionDB{db: database})
	test.whenCommitIsCall(ctx)
	test.ThenNoHaveError(t)
}

func TestCommitFailNoTransaction(t *testing.T) {
	db, _, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.Background()
	test.whenCommitIsCall(ctx)
	test.ThenHaveError(t)
}

func TestCommitFail(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()
	database := gormDB.Begin()
	if database.Error != nil {
		return
	}

	dbMock.ExpectCommit().WillReturnError(gorm.ErrInvalidTransaction)

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.WithValue(context.Background(), transactionCtx{}, &transactionDB{db: database})
	test.whenCommitIsCall(ctx)
	test.ThenHaveError(t)
}

func TestRollbackOk(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()
	database := gormDB.Begin()
	if database.Error != nil {
		return
	}

	dbMock.ExpectRollback()

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.WithValue(context.Background(), transactionCtx{}, &transactionDB{db: database})
	test.whenRollbackIsCall(ctx)
	test.ThenNoHaveError(t)
}

func TestRollbackFailNoTransaction(t *testing.T) {
	db, _, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.Background()
	test.whenRollbackIsCall(ctx)
	test.ThenHaveError(t)
}

func TestRollbackFail(t *testing.T) {
	db, dbMock, closeDB, err := mocks.GetDB()
	if err != nil {
		return
	}
	defer closeDB()

	gormDB, err := mocks.GetGormDB(db)
	if err != nil {
		return
	}

	dbMock.ExpectBegin()
	database := gormDB.Begin()
	if database.Error != nil {
		return
	}

	dbMock.ExpectRollback().WillReturnError(gorm.ErrInvalidTransaction)

	test := givenGormConnectionsScenery(t)
	test.givenDB(gormDB)
	test.givenGormDB()
	ctx := context.WithValue(context.Background(), transactionCtx{}, &transactionDB{db: database})
	test.whenRollbackIsCall(ctx)
	test.ThenHaveError(t)
}
