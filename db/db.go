package db

import (
	"gopkg.in/mgo.v2"
)

type DBConnection struct {
	session *mgo.Session
}

func NewConnection(host string) (conn *DBConnection, err error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}

	return conn, nil
}

func (conn *DBConnection) Use(dbName, tableName string) (db *mgo.Collection) {
	return conn.session.DB(dbName).C(tableName)
}

func (conn *DBConnection) Close() {
	conn.session.Close()
	return
}
