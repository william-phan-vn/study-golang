package main

import "log"

var a int        // declare int variable 'a'
a = 2            // assign 'a' the value of 2

// The zero value of pointers is nil because a nil pointer holds no memory address.
var p *int
p = &a // the variable 'p' contains the memory address of 'a'

// --------------------- DEREFERENCING ------------------------
// The operation *p fetches the value stored at the memory address stored in p.
// This operation is often called "dereferencing".
// Always make sure a pointer is not nil before dereferencing it.
var b int
b = *p // b == 2

var pa *int
pa = &a          // 'pa' now contains to the memory address of 'a'
*pa = *pa + 2    // increment by 2 the value at memory address 'pa'

fmt.Println(a)   // Output: 4
                 // 'a' will have the new value that was changed via the pointer!

// -------------------- STRUCT ------------------
// We could have also created a new Person and immediately stored a pointer to it:
var p *Person
p = &Person{Name: "Peter", Age: 22}


// When we have a pointer to a struct, we don't need to dereference the pointer
// before accessing one of the fields:
var p *Person
p = &Person{Name: "Peter", Age: 22}

fmt.Println(p.Name) // Output: "Peter"
                    // Go automatically dereferences 'p' to allow
                    // access to the 'Name' field

// ------------------ Slices and maps are already pointers ------------------



func main() {
	myString := "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("myString is set to", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to ", s)
	var newValue = "Red"
	*s = newValue
}

// Note that slices and maps are exceptions to the above-mentioned rule. 
// When we pass a slice or a map as arguments into a function, 
// they are treated as pointer types even though there is no explicit * in the type