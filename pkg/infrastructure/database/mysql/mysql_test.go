package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/JhonX2011/GFAWBP/test/doubles/database"
	mocks "github.com/JhonX2011/GFAWBP/test/mocks/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	testConnName = "test-conn-name"
)

type TestMySQLConnectionsScenery struct {
	mockConnections *mocks.MockConnections
	aError          error
	aResult         *sql.DB
	aListResult     []*sql.DB
	aIResult        IDBClient
	iMySQLConn      IDBClient
}

func givenTestMySQLConnectionsScenery(t *testing.T) *TestMySQLConnectionsScenery {
	t.Parallel()
	return &TestMySQLConnectionsScenery{}
}

func (s *TestMySQLConnectionsScenery) givenMockConnections() {
	s.mockConnections = &mocks.MockConnections{}
}

func (s *TestMySQLConnectionsScenery) givenIMysqlConn() {
	s.iMySQLConn = &mySQLConn{connections: s.mockConnections}
}

func (s *TestMySQLConnectionsScenery) whenGetIs(name string) {
	s.mockConnections.On("Get", name).Return(&sql.DB{}, nil)
}

func (s *TestMySQLConnectionsScenery) whenGetIsError(err error) {
	s.mockConnections.On("Get", mock.Anything).Return(nil, err)
}

func (s *TestMySQLConnectionsScenery) whenGetIsExecuted(name string) {
	s.aResult, s.aError = s.iMySQLConn.Get(name)
}

func (s *TestMySQLConnectionsScenery) whenListIs() {
	listConn := append([]*sql.DB{}, &sql.DB{}, &sql.DB{})
	s.mockConnections.On("List").Return(listConn)
}

func (s *TestMySQLConnectionsScenery) whenListIsEmpty() {
	listEmpty := make([]*sql.DB, 0)
	s.mockConnections.On("List").Return(listEmpty)
}

func (s *TestMySQLConnectionsScenery) whenListIsExecuted() {
	s.aListResult = s.iMySQLConn.List()
}

func (s *TestMySQLConnectionsScenery) whenCloseIs() {
	s.mockConnections.On("Close").Return(nil)
}

func (s *TestMySQLConnectionsScenery) whenCloseIsError(err error) {
	s.mockConnections.On("Close").Return(err)
}

func (s *TestMySQLConnectionsScenery) whenNewMySQLIsExecuted(conn []byte) {
	s.aIResult, s.aError = NewMysqlConn(conn)
}

func (s *TestMySQLConnectionsScenery) whenCloseIsExecuted() {
	s.aError = s.iMySQLConn.Close()
}

func (s *TestMySQLConnectionsScenery) thenNoError(t *testing.T) {
	assert.NoError(t, s.aError)
}

func (s *TestMySQLConnectionsScenery) thenHaveError(t *testing.T) {
	assert.NotNil(t, s.aError)
}

func TestNewMysqlOK(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.whenNewMySQLIsExecuted(database.GetMysqlConnStringOK())
	s.thenNoError(t)
}

func TestNewMysqlError(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.whenNewMySQLIsExecuted(database.GetMysqlConnString())
	s.thenHaveError(t)
}

func TestNewMysqlJsonError(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.whenNewMySQLIsExecuted([]byte("invalid json"))
	s.thenHaveError(t)
}

func TestMysqlGetOK(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.givenMockConnections()
	s.givenIMysqlConn()
	s.whenGetIs(testConnName)
	s.whenGetIsExecuted(testConnName)
	s.thenNoError(t)
}

func TestMysqlGetError(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.givenMockConnections()
	s.givenIMysqlConn()
	s.whenGetIsError(fmt.Errorf("unknown connection name %s", testConnName))
	s.whenGetIsExecuted(testConnName)
	s.thenHaveError(t)
}

func TestMysqlListOK(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.givenMockConnections()
	s.givenIMysqlConn()
	s.whenListIs()
	s.whenListIsExecuted()
	s.thenNoError(t)
	assert.Len(t, s.aListResult, 2)
}

func TestMysqlListEmpty(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.givenMockConnections()
	s.givenIMysqlConn()
	s.whenListIsEmpty()
	s.whenListIsExecuted()
	s.thenNoError(t)
	assert.Len(t, s.aListResult, 0)
}

func TestMysqlCloseOK(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.givenMockConnections()
	s.givenIMysqlConn()
	s.whenCloseIs()
	s.whenCloseIsExecuted()
	s.thenNoError(t)
}

func TestMysqlCloseError(t *testing.T) {
	s := givenTestMySQLConnectionsScenery(t)
	s.givenMockConnections()
	s.givenIMysqlConn()
	s.whenCloseIsError(fmt.Errorf("error closing connections"))
	s.whenCloseIsExecuted()
	s.thenHaveError(t)
}
