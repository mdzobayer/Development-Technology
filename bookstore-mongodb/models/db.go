package models

import (
	"gopkg.in/mgo.v2"
)

type Datastore interface {
	AllBooks() ([]Book, error)
}

type DB struct {
	collection *mgo.Collection
}

func (db *DB) AllBooks() ([]Book, error) {

	var bks []Book

	//fmt.Println("Debug 4")

	err := db.collection.Find(nil).All(&bks)

	if err != nil {
		return nil, err
	}
	//fmt.Println("Debug 5")

	//defer db.collection.Close()

	return bks, nil
}

func NewDB(dbName, collectionName string) (*DB, error) {

	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		return nil, err
	}

	//defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbName).C(collectionName)

	if err != nil {
		return nil, err
	}
	return &DB{c}, nil
}
