package gopsql

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cadyrov/goerr"

	_ "github.com/lib/pq" //for psql will work
)

func (c *Config) Connect() (db *DB, e goerr.IError) {
	cu, e := c.ConnectionURL()

	if e != nil {
		return
	}

	dbPs, err := sql.Open("postgres", cu)
	if err != nil {
		e = goerr.New(err.Error())
	}

	db = &DB{false, dbPs}

	return
}

func (c *Config) ConnectionURL() (url string, e goerr.IError) {
	url = "host=%s port=%d user=%s password=%s dbname=%s"

	if c.Host == "" || c.Port == 0 || c.UserName == "" || c.DBName == "" || c.Password == "" {
		e = goerr.New(fmt.Sprintf("config isn't full "+url, c.Host,
			c.Port, c.UserName, c.Password, c.DBName)).HTTP(http.StatusBadRequest)

		return
	}

	if c.SslMode != "" {
		url += " sslmode=" + c.SslMode
	}

	if c.Binary {
		url += " binary_parameters=yes"
	}

	url = fmt.Sprintf(url, c.Host, c.Port, c.UserName, c.Password, c.DBName)

	return
}
