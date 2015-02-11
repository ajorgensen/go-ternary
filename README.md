# Ternary

Somtimes 5 lines is too many

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

As long as your statement returns a boolean value you can nest the ternary statements to create compound expressions.

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
