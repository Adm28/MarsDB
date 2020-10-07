package store

import (
	"log"
	"github.com/syndtr/goleveldb/leveldb"
)

func LoadStore() (*leveldb.DB,error){
  db,err := leveldb.OpenFile("Node/LevelDB",nil)
  if err != nil {
    log.Println(err)
  }
  return db,err
}

func GetValueForKey(key string,db* leveldb.DB) (string,error) {
  value,err := db.Get([]byte(key),nil)
  if err != nil {
    log.Println(err)
  }
  return string([]byte(value)),err
}

func CreateKeyValue(key string,value string,db *leveldb.DB) error {
  err := db.Put([]byte(key),[]byte(value),nil)
  return err
}

func UpdateKeyValue(key string,value string,db *leveldb.DB) error { 
  err := db.Put([]byte(key),[]byte(value),nil)
  return err
}



