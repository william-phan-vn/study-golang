// The rune type in Go is an alias for int32. Given this underlying int32 type,
// the rune type holds a signed 32-bit integer value. However, unlike an int32
// type, the integer value stored in a rune type represents a single Unicode
// character

// Variables of type rune are declared by placing a character inside single quotes:

myRune := '¿'
fmt.Printf("myRune type: %T\n", myRune)
// Output: myRune type: int32

myString := "❗hello"
for index, char := range myString {
  fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\n", index, char, char)
}
// Output:
// Index: 0	Character: ❗		Code Point: U+2757
// Index: 3	Character: h		Code Point: U+0068
// Index: 4	Character: e		Code Point: U+0065
// Index: 5	Character: l		Code Point: U+006C
// Index: 6	Character: l		Code Point: U+006C
// Index: 7	Character: o		Code Point: U+006F