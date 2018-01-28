package person

import (
	"fmt"
	"testing"
)

func TestPopulatePerson (t *testing.T) {

	// A test for the person struct and generator functions

    curPerson := NewPerson("1234", "prsn@gmail.com", "Prsn")
    curPerson.NewPersonURLGroup("Grp")
    curPerson.AddPersonURL("Grp", "url1","http://url1.com/")
    curPerson.AddPersonURL("Grp", "url2","http://url2.com/")
    fmt.Println(curPerson)

    expectOutput := "&{1234 prsn@gmail.com Prsn map[Grp:map[url1:http://url1.com/ url2:http://url2.com/]]}"

    if curOutput := fmt.Sprint(curPerson); curOutput != expectOutput {
        t.Errorf("Expected '%s' but got '%s'", expectOutput, curOutput)
    }

}
