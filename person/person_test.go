package person

import (
	"fmt"
	"testing"
)

func TestPopulatePerson (t *testing.T) {

	// A test for the person struct and generator functions

    curPerson := NewPerson("Prsn")
    curPerson.NewPersonURLGroup("Grp")
    curPerson.AddPersonURL("Grp", "url1","http://url1.com/")
    curPerson.AddPersonURL("Grp", "url2","http://url2.com/")
    fmt.Println(curPerson)

    // Output:
    // &{Prsn map[Grp:map[url1:http://url1.com/ url2:http://url2.com/]]}

}
