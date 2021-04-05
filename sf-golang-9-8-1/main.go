package main

import "fmt"

const (
	GenderMale    = "муж."
	GenderFemale  = "жен."
	NoSuspectInfo = "В базе данных нет информации по запрошенным подозреваемым"
)

type Person struct {
	Name     string
	LastName string
	Age      uint
	Gender   string
	Crimes   uint
}

func (p *Person) String() string {
	return fmt.Sprintf(
		"Имя:          %s\nФамилия:      %s\nВозраст:      %d\nПол:          %s\nПреступлений: %d\n",
		p.Name, p.LastName, p.Age, p.Gender, p.Crimes,
	)
}

type Persons map[string]Person

func main() {

	if p := MostDangerousSuspect(suspectsA, personsData); p != nil {
		fmt.Println(p)
	} else {
		fmt.Println(NoSuspectInfo)
	}

	if p := MostDangerousSuspect(suspectsB, personsData); p != nil {
		fmt.Println(p)
	} else {
		fmt.Println(NoSuspectInfo)
	}

}

// MostDangerousSuspect - returns a Person with max number of crimes within suspects.
func MostDangerousSuspect(suspects []string, persons Persons) *Person {
	var maxCrimes uint
	var result *Person
	for _, name := range suspects {
		if p, ok := persons[name]; ok {
			if p.Crimes > maxCrimes {
				maxCrimes = p.Crimes
				result = &p
			}
		}
	}
	return result
}

// Sample data
var suspectsA = []string{"Ольга", "Геннадий", "Сергей", "Александр"}
var suspectsB = []string{"Джон", "Майкл", "Дональд", "Марио"}

var personsData = Persons{
	"Иван": {
		Name:     "Иван",
		LastName: "Петров",
		Age:      19,
		Gender:   GenderMale,
		Crimes:   2,
	},
	"Николай": {
		Name:     "Николай",
		LastName: "Потапов",
		Age:      43,
		Gender:   GenderMale,
		Crimes:   11,
	},
	"Ольга": {
		Name:     "Ольга",
		LastName: "Никитина",
		Age:      34,
		Gender:   GenderFemale,
		Crimes:   7,
	},
	"Василий": {
		Name:     "Василий",
		LastName: "Сидоров",
		Age:      53,
		Gender:   GenderMale,
		Crimes:   19,
	},
	"Елена": {
		Name:     "Елена",
		LastName: "Николаева",
		Age:      21,
		Gender:   GenderFemale,
		Crimes:   5,
	},
	"Сергей": {
		Name:     "Сергей",
		LastName: "Иванов",
		Age:      42,
		Gender:   GenderMale,
		Crimes:   16,
	},
	"Борис": {
		Name:     "Борис",
		LastName: "Боренко",
		Age:      54,
		Gender:   GenderMale,
		Crimes:   25,
	},
	"Максим": {
		Name:     "Максим",
		LastName: "Деркач",
		Age:      48,
		Gender:   GenderMale,
		Crimes:   1,
	},
	"Светлана": {
		Name:     "Светлана",
		LastName: "Тихомирова",
		Age:      42,
		Gender:   GenderFemale,
		Crimes:   9,
	},
	"Александр": {
		Name:     "Александр",
		LastName: "Павлов",
		Age:      27,
		Gender:   GenderMale,
		Crimes:   13,
	},
}
