package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	_, err = s.db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)

}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGet() {

	clientExpected, _ := entity.NewClient("John Doe", "f1@gmail.com")
	s.clientDB.Save(clientExpected)

	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES ('1', 'John Doe', 'f1@gmail.com', '2020-01-01')")
	clientResult, err := s.clientDB.Get(clientExpected.ID)

	s.Nil(err)
	s.Equal(clientExpected.Name, clientResult.Name)
	s.Equal(clientExpected.Email, clientResult.Email)
	s.Equal(clientExpected.ID, clientResult.ID)

}

func (s *ClientDBTestSuite) TestSave() {

	clientExpected, _ := entity.NewClient(
		"John Doe",
		"f1@gmail.com",
	)
	err := s.clientDB.Save(clientExpected)

	s.Nil(err)
}
