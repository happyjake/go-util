package store

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/gorm"

	// for postgress drive
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// Store the storage for everything
type Store struct {
	DB *gorm.DB
}

var instance *Store
var once sync.Once

// Close close db connection
func (store Store) Close() {
	store.DB.Close()
}

// GetStore the store singleton
func GetStore() *Store {
	once.Do(func() {
		db, err := gorm.Open("postgres", viper.GetString("database"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// dont close it now.
		// defer db.Close()

		instance = &Store{DB: db}
	})
	return instance
}
