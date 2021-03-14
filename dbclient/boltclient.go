package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/yigithanbalci/amazon-clone-account-service/model"
)

type BoltClient struct {
	boltdb *bolt.DB
}

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountID string) (model.Account, error)
	Seed()
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltdb, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initializeBucket() {
	bc.boltdb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("accountbucket"))
		if err != nil {
			return fmt.Errorf("Creating bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000 + i)

		acc := model.Account{
			ID:   key,
			Name: "person_" + strconv.Itoa(i),
		}

		jsonbytes, _ := json.Marshal(acc)

		bc.boltdb.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("accountbucket"))
			err := b.Put([]byte(key), jsonbytes)
			return err
		})
	}
	fmt.Printf("seed %v fake accounts...\n", total)
}

// QueryAccount returns account based on accountID
func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("accountbucket"))

		accountBytes := b.Get([]byte(accountID))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}

		json.Unmarshal(accountBytes, &account)

		return nil
	})

	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}
