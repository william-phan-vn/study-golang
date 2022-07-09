// map[KeyType]ElementType
 // Creation with map literal
 foo := map[string]int{}

 // Add a value in a map with the `=` operator:
 foo["bar"] = 42
 // Here we update the element of `bar`
 foo["bar"] = 73
 // To retrieve a map value, you can use
 baz := foo["bar"]
 // To delete an item from a map, you can use
 delete(foo, "bar")

 value, exists := foo["baz"]
  // If the key "baz" does not exist,
  // value: 0; exists: falsei