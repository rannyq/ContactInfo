package ContactInfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type FileStorage struct {
}

type ContactInfo struct {
	ID     int
	Name   string
	Street string
	City   string
	Zip    string
	Phone  string
}

type Contacts struct {
	Contacts []ContactInfo
}

type Storage interface {
	WriteFile(entityID string, contactinfo ContactInfo) error
}

func (f *FileStorage) WriteFile(entityID string, contactinfo ContactInfo) error {

	filename := fmt.Sprintf("%s.json", entityID)

	jsonBytes, err := json.Marshal(contactinfo)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonBytes, 0644)

	if err != nil {
		return err
	}

	return nil
}

