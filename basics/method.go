// A method is a function with a special receiver argument. The receiver appears
// in its own argument list between func keyword and the name of the method.
func (receiver type) MethodName(parameters) (returnTypes) {
}

// You can only define a method with a receiver whose type is defined in the
// same package as the method.
type Person struct {
	Name string
}

func (p Person) Greetings() string {
	return fmt.Sprintf("Welcome %s !", p.Name)
}

p := Person{Name: "Bronson"}
fmt.Println(p.Greetings())
// Output: Welcome Bronson !

// Remember: a method is just a function with a receiver argument. Methods help
// to avoid naming conflicts - since a method is tied to a particular receiver
// type, you can have the same method name on different types.
import "math"

type rect struct {
	width, height int
}
func (r rect) area() int {
	return r.width * r.height
}

type circle struct {
	radius int
}
func (c circle) area() float64 {
	return 2*c.radius*math.Pi
}

// ----------------- RECEIVER TYPE ------------------------------
// There are two types of receivers, value receivers, and pointer receivers.
type rect struct {
	width, height int
}
func (r *rect) squareIt() {
	r.height = r.width
}

r := rect{width: 10, height: 20}
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 20

r.squareIt()
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 10

