package main

import (
	"fmt"
	"slices"
	"maps"
)

func sum(a, b int)int{
	return a + b
}

func main(){
	fmt.Println("hello world")

	//declaring variable
	var fname = "ian"
	fmt.Println(fname)

	//shorthand for declaring a variable
	lname := "aule"
	fmt.Println(lname)

	//declaring a constant
	const mywoman = "kalule"
	fmt.Println(mywoman)

	//loops --go only uses a for loop
	j := 0
	for j < 3{
		fmt.Println(j)
		j++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} 

	//modern way
	for k := range 5{
		fmt.Println(k)
	}

	//control flow
	l := 4
	if l%2 == 0 {
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}

	//else if
	gpa := 4.5
	if gpa > 4.4 {
		fmt.Println("First class")
	}else if gpa > 3.6{
		fmt.Println("Second class upper")
	}else if gpa > 2.6{
		fmt.Println("Second class lower")
	}else{
		fmt.Println("Pass degree")
	}

	//switch statement
	cgpa := 3
	switch cgpa {
	case 1:
		fmt.Println("Pass degree")
	case 2:
		fmt.Println("Second class lower")
	case 3:
		fmt.Println("Second class upper")
	}

	//arrays
	var names [2]string
	names[1] = "ian"
	names[0] = "aule"

	fmt.Println(names)

	//declare and intialize
	z :=[3]int{1, 2, 3}
	fmt.Println(z)

	//unspecified length
	a := [...]int{4,5}
	fmt.Println(a)

	//slices --special feature for go
	

	s := make([]string, 3)
	s[0] = "ian"
	s[1] = "aule"
	s[2] = "ayaan"

	t := make([]string, 3)
	t[0] = "Aban"
	t[1] = "skylar"
	t[2] = "Immy"
	fmt.Println(t)

	s = append(s, "shamirah")
	s = append(s, "jaziirah", "kalule")
	fmt.Println(s)

	if slices.Equal(s, t){
		fmt.Println("s == t")
	}else{
		fmt.Println("Not equal")
	}

	//maps
	m := make(map[string]string)

	m["husband"] = "Ayaan"
	m["wife"] = "shamirah"
	m["brother"] = "Anyone"

	delete(m, "brother")

	n := make(map[string]string)
	n["older"] = "Ayaan"
	n["brighter"] = "Shamirah"

	if maps.Equal(m, n){
		fmt.Println("Equal")
	}else{
		fmt.Println("not equal")
	}


	//functions
	result := sum(2, 2)
	fmt.Println(result)

	//
}