package badgerdb

import (
	"fmt"
	"github.com/dgraph-io/badger"
)

type database struct {
	Client *badger.DB
}

//Open opens a Badger client to the database at dir()
func Open() (*database, error) {
	//Setup Badger options
	op := badger.DefaultOptions("data")

	//Open Badger Client
	client, err := badger.Open(op)
	if err != nil {
		return nil, fmt.Errorf("db open fail: %w", err)
	}

	//Put client inside the Database and return
	return &database{client}, nil
}

//Close closes the Badger client to the database at dir()
func (Db *database) Close() {
	if err := Db.Client.Close(); err != nil {
		panic(fmt.Errorf("db close fail:%w", err))
	}
}

//GetEntry function to fetch the data from the database
func (Db *database) GetEntry(key []byte) (value []byte, err error) {
	//Define a view method on the database
	err = Db.Client.View(func(view *badger.Txn) error {
		//Attempt to get the Item for the given key
		item, err := view.Get(key)
		if err != nil {
			return fmt.Errorf("db get on Key '%x' fail:'%w'", key, err)
		}

		//Getting the value from the Item
		if err = item.Value(func(val []byte) error {
			value = val
			return nil
		}); err != nil {
			return fmt.Errorf("db value get on the key '%x' fail:%w", key, err)
		}
		return nil

	})
	return
}

//SetEntry to  store the data in the database
func (Db *database) SetEntry(key, value []byte) error {
	//Define an Update transaction the database
	return Db.Client.Update(func(view *badger.Txn) error {
		//Attempt to set the key-value pair to the database
		if err := view.Set(key, value); err != nil {
			return fmt.Errorf("db set for Key '%x' failed: %w", key, err)
		}
		return nil
	})
}
