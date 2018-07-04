package person

import (
	"errors"
	"fmt"
    "github.com/satori/go.uuid"
)

// create a structure and methods to maintain an object for the cards

type Person struct {
	UID			string
	EMail		string
	Persona		string
    DisplayName   string
    Background	string
    URLList     map[string]map[string]string
}

// function to create a new person
func NewPerson(curEMail string, curPersona string, curDisplayName string, curBackground string ) (*Person , error) {
    p := new(Person)

    // create a RFC 4122 V4 UUID
    curUUID, err := uuid.NewV4()
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Failed to get a UUID: %s", err))
    }

    p.UID = curUUID.String()
    p.EMail = curEMail
    p.Persona = curPersona
    p.DisplayName = curDisplayName
    p.Background = curBackground
    p.URLList = make(map[string]map[string]string)
    return p, nil
}

// add a group of URL links 
func (p *Person) NewPersonURLGroup(curGroupName string) {
    p.URLList[curGroupName] = make(map[string]string)
}

// add a URL to a group
func (p *Person) AddPersonURL(curGroupName string, curURLLabel string, curURLHREF string) {
    p.URLList[curGroupName][curURLLabel] = curURLHREF
}
