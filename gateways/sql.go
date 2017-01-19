package gateways

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //load the driver anonymously 

	"github.com/yuderekyu/expresso-subscription/config"
)

type SQL interface {
	Modify(string, ...interface{}) error
	Select(string, ...interface{}) (*sql.Rows, error)
	Destroy()
}

/*MySQL implements SQL with the mysql driver */
type MySQL struct {
	DB *sql.DB
}

/*NewSQL returns an instance of MySQL with the given connection configuration */
func NewSQL(config config.MySQL) (*MySQL, error) {
	db, err := sql.Open(
		"mysql",
		config.User+":"+config.Password+"@tcp("+config.Host+":"+string(config.Port)+")/"+config.Database+"?parseTime=true",
	)

	return &MySQL{DB: db}, err
}

/*Modify executes any query which changes the db and doesn't return result rows */
func (s *MySQL) Modify(query string, values ...interface{}) error {
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		fmt.Printf("ERROR: unable to prepare query %s\n", query)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		fmt.Printf("ERROR: unable to execute query %s\n", query)
		return err
	}

	//success
	return nil
}

/*Select gets rows from a select query*/
func (s *MySQL) Select(query string, values ...interface{}) (*sql.Rows, error) {
	if values == nil {
		values = make([]interface{}, 0)
	}
	rows, err := s.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("ERROR: unable to run select query %s\n", query)
		return nil, err
	}

	return rows, nil
}

/*Destroy cleans up the MySQL instance*/
func (s *MySQL) Destroy() {
	s.DB.Close()
}