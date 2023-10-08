package pubs

import (
	"context"
	"fmt"

	"github.com/blueskyxi3/pillow/pkg/pillow"
)

func DBConnect(dsn string) *pillow.Client {
	client, err := pillow.New(dsn)

	if err != nil {
		panic(err)
	}

	return client
}

func GetAllEntries() string {
	//get all database records
	return "this works"

}

func CreateEntry(dsn string, hash string) string {
	//add a record to the database

	client := DBConnect(dsn)
	db := client.Database(context.TODO(), "docs")

	document := map[string]interface{}{
		"_id":      hash,
		"filename": "filename_sample",
		"metadata": map[string]any{
			"data 1": 1,
			"data 2": 2,
			"data 3": 3,
		},
		"devices": map[string]any{
			"device_id": "device_name",
		},
	}

	//TODO - if device does not already exist in the database, crease a devices document entry for the device

	_, err := db.CreateDocument(context.TODO(), document)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tenant created")
	return "this works"

}

func DeleteEntry() string {
	//removes file record and deletes local copy
	return "this works"
}

func SaveEntry() string {
	//fetch a copy of a file and save it on a device
	return "this works"
}

func UnsaveEntry() string {
	//remove local copy, update database saying file is no longer saved on this device
	return "this works"
}

func GetDeviceList() string {
	//returns a list of all devices
	return "this works"
}
