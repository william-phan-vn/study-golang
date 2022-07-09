// https://pkg.go.dev/fmt
// %v	the value in a default format
// 	when printing structs, the plus flag (%+v) adds field names
// %#v	a Go-syntax representation of the value
// %T	a Go-syntax representation of the type of the value
// %%	a literal percent sign; consumes no value
import "fmt"

food := "taco"
fmt.Sprintf("Bring me a %s", food)
// Returns: Bring me a taco
