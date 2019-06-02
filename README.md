kli
===

Command-Line Interface for Computer Science and Go course.

---

Usage
-----

```go
package main

import (
	"fmt"
	"github.com/kxnes/kli"
)

func main() {
	// flags
	op := &kli.RuneFlag{Name: "-o"}
	array := &kli.Float64SliceFlag{Name: "-a"}

	// args (will be os.Exit if error occurred)
	args := kli.Args{Flags: []kli.Flag{op, array}}
	args.Parse()

	// parsed
	fmt.Println(op.Value, array.Value)
}
```
