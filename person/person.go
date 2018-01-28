package person

// create a structure and methods to maintain an object for the cards

type Person struct {
	UID			string
	EMail		string
    GivenName   string
    URLList     map[string]map[string]string
}

// function to create a new person
func NewPerson(curUID string, curEMail string, curGivenName string ) *Person {
    p := new(Person)
    p.UID = curUID
    p.EMail = curEMail
    p.GivenName = curGivenName
    p.URLList = make(map[string]map[string]string)
    return p
}

// add a group of URL links 
func (p *Person) NewPersonURLGroup(curGroupName string) {
    p.URLList[curGroupName] = make(map[string]string)
}

// add a URL to a group
func (p *Person) AddPersonURL(curGroupName string, curURLLabel string, curURLHREF string) {
    p.URLList[curGroupName][curURLLabel] = curURLHREF
}
