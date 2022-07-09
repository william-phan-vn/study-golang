// An interface type is effectively a set of method signatures. Any type that
// defines those methods "implements" the interface implicitly. There is no
// implements keyword in Go. Here is what an interface definition might look
// like:

type InterfaceName interface {
    MethodOne() MethodOneReturnType
    MethodTwo(paramOne ParamOneType, paramTwo ParamTwoType) MethodTwoReturnType
    ...
}

// There is one very special interface type in Go: the empty interface type that
// contains zero methods. The empty interface type is written like this:
// interface{}. Since it has no methods, every type implements the empty
// interface type. This is helpful for defining a function that can generically
// accept any value. In that case, the function parameter uses the empty
// interface type.