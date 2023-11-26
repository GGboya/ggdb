package main

import (
	"GG_DB"
	"fmt"
)

func main() {
	opts := GG_DB.DefaultOptions
	opts.DirPath = "tempdata"

	db, err := GG_DB.Open(opts)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("name"), []byte("bitcask"))
	if err != nil {
		panic(err)
	}

	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}

	fmt.Println("val = ", string(val))

	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}
}
