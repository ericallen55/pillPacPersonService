package main

var people = []Person{
	{FirstName: "Eric", LastName: "Allen", Age: "Old", Id: 0},
	{FirstName: "Stephanie ", LastName: "Fillion", Age: "Young", Id: 1},
}
var id = 2

func getAllPeopleDb() []Person {
	return people
}

func addPersonDb(person Person) int {
	person.Id = id
	id++
	people = append(people, person)
	return person.Id
}

func getPersonDb(id int) Person {
	for _, n := range people {
		if id == n.Id {
			return n
		}
	}
	return Person{}
}

func deletePersonDb(id int) {
	for i, n := range people {
		if id == n.Id {
			people[i] = people[len(people)-1]
			people = people[:len(people)-1]
		}
	}
}

func updatePersonDb(id int, person Person) Person {
	for i, n := range people {
		if id == n.Id {
			people[i].Age = person.Age
			people[i].FirstName = person.FirstName
			people[i].LastName = person.LastName
			return people[i]
		}
	}
	return Person{}
}
