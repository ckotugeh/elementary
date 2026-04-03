package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
	
	fmt.Println(FirstWord("   hello world")) // outputs: hello
    fmt.Println(FirstWord("Go is awesome")) // outputs: Go
    fmt.Println(FirstWord("    "))          // outputs: (empty)
    fmt.Println(FirstWord(""))              // outputs: (empty)
	fmt.Println(LastWord("hello world"))       // outputs: world
    fmt.Println(LastWord("Go is awesome"))     // outputs: awesome
    fmt.Println(LastWord("   multiple spaces ")) // outputs: spaces
    fmt.Println(LastWord(""))                  // outputs: (empty))

}
