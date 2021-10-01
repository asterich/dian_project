package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Job  string `json:"job"`
	Age  int    `json:"age"`
}

func (p Person) printAll() {
	fmt.Println(p.Name, p.Job, p.Age)
}

type PersonList struct {
	List []Person `json:"list"`
}

func (l PersonList) printAll() {
	for index, p := range l.List {
		fmt.Print(index, " ")
		p.printAll()
	}
}

func main() {
	var president Person = Person{
		"Joe Biden",
		"American President",
		78,
	}

	var str, err1 = json.Marshal(president)
	if err1 != nil {
		fmt.Printf("Marshal failed, err = %v", err1)
		return
	}
	fmt.Printf("%s\n", str)

	var p2 Person
	var jsonstr = `{"name":"Donald Trump", "job":"merchant","age":75}`
	json.Unmarshal([]byte(jsonstr), &p2)
	p2.printAll()

	var Yes PersonList = PersonList{
		[]Person{
			{
				"anderson", "singer", 34,
			},
			{
				"howe", "guitarist", 33,
			},
			{
				"wakeman", "keyboardist", 32,
			},
			{
				"bruford", "drummer", 31,
			},
		},
	}
	Yes.printAll()

	var jsonstr2, err2 = json.Marshal(Yes)
	if err2 != nil {
		fmt.Println("Marshal failed, err: ", err2)
	}
	fmt.Println(string(jsonstr2))
	var l PersonList
	json.Unmarshal([]byte(jsonstr2), &l)
	l.printAll()
	fmt.Println("Yes, yyds!")
}
