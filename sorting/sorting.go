package main
import (
	"sort"
	"fmt"
	)

type Person struct {
	name string
	age int
}

type People []Person

// make People type conform to the sort interface
// reciever = People type p
// method = Len()
// return type = int
func (p People) Len() int { return len(p) }
func (p People) Less(i, j int) bool {return p[i].age < p[j].age }
func (p People) Swap(i,j int) { p[i], p[j] = p[j], p[i] }

type PeopleByName struct {People}
func (p PeopleByName) Less(i, j int) bool { return p.People[i].name < p.People[j].name }

func main() {
	unsorted := []int{1,0,2,3,6,19,16,25,31,21,55}
	fmt.Println(unsorted)
	
	sort.Ints(unsorted)
	fmt.Println(unsorted)

	a := People{Person{"Pete", 20},Person{"Dave",22}, Person{"Sarah", 19}}
	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)

	b := PeopleByName{a}
	sort.Sort(b)
	fmt.Println(b.People)
}