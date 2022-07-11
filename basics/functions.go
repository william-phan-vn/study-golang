import "fmt"

// No parameters, no return value
func PrintHello() {
    fmt.Println("Hello")
}
// Called like this:
PrintHello()

// One parameter, one return value
func Hello(name string) string {
  return "Hello " + name
}
// Called like this:
greeting := Hello("Dave")

// Multiple parameters, multiple return values
func SumAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
// Called like this:
aplusb, atimesb := SumAndMultiply(a, b)

// Named Return Values and Naked Return
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
  sum, mult = a+b, a*b
  sum -= c
  mult -= c
  return
}

// ------------------ Variadic function ----------------------

// The variadic parameter must be the last parameter of the function.
func find(num int, nums ...int) {
  fmt.Printf("type of nums is %T\n", nums)

  for i, v := range nums {
      if v == num {
          fmt.Println(num, "found at index", i, "in", nums)
          return
      }
  }

  fmt.Println(num, "not found in ", nums)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
  fmt.Printf("type of nums is %T\n", value)
return append(value, slice...)
}

func main() {
  list := []int{1, 2, 3}
  // find(1, list...) // "find" defined as shown above
  fmt.Println(PrependItems(list, 1))
}


// ---------------- First class functions ------------------- 

// In Go, functions are first-class values. This means that you can do with
// functions the same things you can do with all other values - assign functions
// to variables, pass them as arguments to other functions or even return
// functions from other functions.
import "fmt"

func engGreeting(name string) string {
	return fmt.Sprintf("Hello %s, nice to meet you!", name)
}

func espGreeting(name string) string {
	return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := engGreeting			// greeting is a variable of type func(string) string
fmt.Println(greeting("Alice"))	// Hello Alice, nice to meet you!

greeting = espGreeting
fmt.Println(greeting("Alice")) 	// ¡Hola Alice, mucho gusto!

// Function values provide an opportunity to parametrize functions not only with
// data but with behavior too. In the following example, we are passing
// behaviour to the dialog function via the greetingFunc parameter:
func dialog(name string, greetingFunc func(string) string) {
	fmt.Println(greetingFunc(name))
	fmt.Println("I'm a dialog bot.")
}

func espGreeting(name string) string {
	return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := espGreeting
dialog("Alice", greeting)
// Output:
// ¡Hola Alice, mucho gusto!
// I'm a dialog bot.

// -------------- Function type ------------------

// Using function values is possible thanks to the function types in Go. A
// function type denotes the set of all functions with the same sequence of
// parameter types and the same sequence of result types. User-defined types can
// be declared on top of function types. For instance, the dialog function from
// the previous examples can be updated as following:
type greetingFunc func(string) string

func dialog(name string, f greetingFunc) {
	fmt.Println(f(name))
	fmt.Println("I'm a dialog bot.")
}

// ---------------- Anonymous functions ------------------

// Another powerful tool that is available thanks to first-class functions
// support is anonymous functions. Anonymous functions are defined at their
// point of use, without a name following the func keyword. Such functions have
// access to the variables of the enclosing function.
func fib() func() int {
	var n1, n2 int

	return func() int {
		if n1 == 0 && n2 == 0 {
			n1 = 1
		} else {
			n1, n2 = n2, n1 + n2
		}
		return n2
	}
}

next := fib()
for i := 0; i < N; i++ {
  fmt.Printf("F%d\t= %4d\n", i, next())
}
// A call to fib declares the variables n1 and n2 and returns an anonymous
// function that, in turn, changes the values of these variables each time the
// function is called. Nth calls of the anonymous function return the Nth number
// of the Fibonacci sequence starting from 0. The anonymous inner function has
// access to the local variables (n1 and n2) of the enclosing function fib. This
// is a great way to have function values keep state between calls. We say that
// the anonymous function is a closure of the variables n1 and n2. Closures are
// widely used in programming and you might see other languages supporting them.
