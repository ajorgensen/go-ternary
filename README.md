# Ternary [![Build Status](https://secure.travis-ci.org/ajorgensen/go-ternary.png?branch=master)](http://travis-ci.org/ajorgensen/go-ternary) 

The ternary operator has been excluded from the go language syntax. The [official](http://golang.org/doc/faq#Does_Go_have_a_ternary_form) documentation recommends using the long form if statement instead but somtimes 5 lines is just too many. This library allows you to write simple ternary statements to replace larger if statements.

```go
if expr {
    n = trueVal

} else {
    n = falseVal

}
```
could become
```go
n := ternary.If(expr, trueVal, falseVal)
```

# Basic Usage

Save the following as `foo.go`
```go
package main

import (
	"fmt"
	"github.com/ajorgensen/go-ternary"
)

func main() {
	foo := ternary.If(true, "true", "false")
	fmt.Println(foo)
}

```

Install the go-ternary package: 
```
go get github.com/ajorgensen/go-ternary
```


Then try it out:
```
go run foo.go
```

You should see the string `true` printed in the terminal output.

# Nesting statements

You can nest the ternary statements to create more complex expressions

```go
package main

import (
	"fmt"
	"github.com/ajorgensen/go-ternary"
)

func main() {
	bar := 1
	foo := ternary.If(true, ternary.If(bar == 1, 1, 2), 2)
	fmt.Println(foo)
}
```

Running this you should see `1` printed out.

# Mixed types

You can have mixed types in the ternary expression as well.

```go
package main

import (
	"fmt"
	"github.com/ajorgensen/go-ternary"
)

func main() {
	bar := true
	foo := ternary.If(true, ternary.If(bar, "true", 0), false)
	fmt.Println(foo)
}
```

You should see the string `true` printed out.
