// Error handling is not done via exceptions in Go. Instead, errors are normal
// values of the interface type error.

// Creating and Returning Errors
// You do not have to always implement the error interface yourself. To create a
// simple error, you can use the errors.New() function that is part of the
// standard library package errors. These error variables are also called
// sentinel errors and by convention their names should start with Err or err
// (depending on whether they are exported or not). You should use error
// variables instead of directly writing errors.New in cases where you use an
// error multiple times or where you want the consumer of your package to be
// able to check for the specific error.

import "errors"

var ErrSomethingWrong = errors.New("something went wrong")
ErrSomethingWrong.Error() // returns "something went wrong"

// An error is by convention the last value returned in a function with multiple
// return values. If the function returns an error, it should always return the
// zero value for other returned values:
import "errors"

// Do this:
func GoodFoo() (int, error) {
  return 0, errors.New("Error")
}

// Not this:
func BadFoo() (int, error) {
  return 10, errors.New("Error")
}

// Return nil for the error when there are no errors:
func Foo() (int, error) {
	return 10, nil
  }
  