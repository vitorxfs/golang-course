package main

import "fmt"

func main() {
	age := 26; // regular variable

	var agePointer *int;
	agePointer = &age;

	fmt.Println("Age: ", *agePointer);

	editAgeToAdultYears(agePointer);
	fmt.Println(*agePointer);
}

func editAgeToAdultYears(age *int) {
	*age -= 18;
}
