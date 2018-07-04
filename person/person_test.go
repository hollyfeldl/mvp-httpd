package person

import (
	"fmt"
	"testing"
    "github.com/satori/go.uuid"
)

func TestPopulatePerson (t *testing.T) {

	// A test for the person struct and generator functions

    curPerson, err := NewPerson("prsn@gmail.com", "Prsn", "Prsn Name", "mypic.jpg")
    if err != nil {
        t.Errorf("Error creating new person", err)
    }

    curPerson.NewPersonURLGroup("Grp")
    curPerson.AddPersonURL("Grp", "url1","http://url1.com/")
    curPerson.AddPersonURL("Grp", "url2","http://url2.com/")

    // test UUID for a type 4 version
    if theUUID, _ := uuid.FromString(curPerson.UID); theUUID.Version() != 4 {
        t.Errorf("UUID - Expected V4 UUID but got %b", theUUID.Version())
    }

    // test email
    if curPerson.EMail != "prsn@gmail.com" {
        t.Errorf("EMAIL - Expected '%s' but got '%s'", "prsn@gmail.com", curPerson.EMail)
    }

    // test persona
    if curPerson.Persona != "Prsn" {
        t.Errorf("PERSONA - Expected '%s' but got '%s'", "Prsn", curPerson.Persona)
    }

    // test display name
    if curPerson.DisplayName != "Prsn Name" {
        t.Errorf("DISPLAYNAME - Expected '%s' but got '%s'", "Prsn Name", curPerson.DisplayName)
    }

    // test background
    if curPerson.Background != "mypic.jpg" {
        t.Errorf("BACKGROUND - Expected '%s' but got '%s'", "mypic.jpg", curPerson.Background)
    }


    // test the URL list
    expectOutput := "map[Grp:map[url1:http://url1.com/ url2:http://url2.com/]]"
    if curOutput := fmt.Sprint(curPerson.URLList); curOutput != expectOutput {
        t.Errorf("URLLIST - Expected '%s' but got '%s'", expectOutput, curOutput)
    }

}
