package person

// create a structure and methods to maintain an object for the cards

type Person struct {
    GivenName   string
    URLList     map[string]map[string]string
}

func NewPerson(curGivenName string) *Person {
    p := new(Person)
    p.GivenName = curGivenName
    p.URLList = make(map[string]map[string]string)
    return p
}

func (p *Person) NewPersonURLGroup(curGroupName string) {
    p.URLList[curGroupName] = make(map[string]string)
}

func (p *Person) AddPersonURL(curGroupName string, curURLLabel string, curURLHREF string) {
    p.URLList[curGroupName][curURLLabel] = curURLHREF
}
