package ContactInfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type FileStorage struct {
}

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
}

type Address struct{
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
	Country string
}

type Phone struct{
	AreaCode string
	Number   string
}

type ContactInfo struct {
	ID     int
	Person Person `json:"Person,omitempty"`
	Address Address `json:"Address,omitempty"`
	Phone Phone `json:"Phone,omitempty"`
	EmailAddress  string
}

type Contacts struct {
	Contacts []ContactInfo
}

type Storage interface {
	WriteFile(entityID string, contactinfo ContactInfo) error
}

func getAddresses(s string) Address {
	var a Address
	return a
}

func getPerson(s string) Person {
	var p Person
	return p
}

func getPhoneNumbers(s string) Phone {
	var p Phone
	return p
}

func (f *FileStorage) WriteFile(entityID string, contactinfo ContactInfo) error {

	filename := fmt.Sprintf("%s.json", entityID)

	//create json from struct
	jsonBytes, err := json.Marshal(contactinfo)

	if err != nil {
		return err
	}

	//write file to disk
	err = ioutil.WriteFile(filename, jsonBytes, 0644)

	if err != nil {
		return err
	}

	return nil
}

