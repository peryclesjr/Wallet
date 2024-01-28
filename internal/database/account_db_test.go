package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	DB        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.DB = db

	_, err = s.DB.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	s.Nil(err)

	_, err = s.DB.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.Nil(err)

	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("Peter Parker", "spider@gmail.com")

}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.DB.Close()
	s.DB.Exec("DROP TABLE accounts")
	s.DB.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	accountExpected := entity.NewAccount(s.client)
	err := s.accountDB.Save(accountExpected)

	s.Nil(err)
	s.Equal(accountExpected.ID, accountExpected.ID)

}

func (s *AccountDBTestSuite) TestGet() {
	s.DB.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?,?,?,?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	accountExpected := entity.NewAccount(s.client)
	err := s.accountDB.Save(accountExpected)
	s.Nil(err)

	acccount, err := s.accountDB.Get(accountExpected.ID)
	s.Nil(err)
	s.Equal(accountExpected.ID, acccount.ID)
	s.Equal(accountExpected.Client.ID, acccount.Client.ID)
	s.Equal(accountExpected.Balance, acccount.Balance)

}
