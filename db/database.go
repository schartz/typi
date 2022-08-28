package db

import (
	"fmt"
	"github.com/ostafen/clover/v2"
)

const DATABASE = "_DB_GOTYPI"

func DatabaseTest() {
	db, _ := clover.Open(DATABASE)
	err := db.CreateCollection("myCollection")
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert a new document inside the collection
	doc := clover.NewDocument()
	doc.Set("hello", "clover!")

	// InsertOne returns the id of the inserted document
	docId, _ := db.InsertOne("myCollection", doc)
	fmt.Println(docId)
}
