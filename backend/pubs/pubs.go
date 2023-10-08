package pubs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blueskyxi3/pillow/pkg/pillow"
)

func DBConnect(dsn string) *pillow.Client {
	client, err := pillow.New(dsn)

	if err != nil {
		panic(err)
	}

	return client
}

func writeFile(file []byte) {
	if err := os.WriteFile("file.txt", file, 0666); err != nil {
		log.Fatal(err)
	}
}

func GetAllEntries(dsn string) map[string]interface{} {
	//get all database records
	client := DBConnect(dsn)
	db := client.Database(context.TODO(), "docs")

	list, err := db.ListDocuments(context.TODO())
	if err != nil {
		panic(err)
	}

	return list

}

func CreateEntry(dsn string, hash string, file []byte) string {
	//add a record to the database

	client := DBConnect(dsn)
	db := client.Database(context.TODO(), "docs")

	document := map[string]interface{}{
		"_id":        hash,
		"filename":   "filename_sample",
		"bytestring": file,
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
	return "File created"

}

func DeleteEntry(dsn string, hash string) string {
	//removes file record and deletes local copy
	client := DBConnect(dsn)
	db := client.Database(context.TODO(), "docs")

	_, err := db.DeleteDocument(context.TODO(), hash)
	if err != nil {
		panic("Error deleting document")
	}

	//TODO: delete local copy

	return "File deleted"
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
