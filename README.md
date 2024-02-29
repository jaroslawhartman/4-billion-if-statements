# Intro

This program is inspired by [4 billion if statements](https://andreasjhkarlsson.github.io//jekyll/update/2023/12/27/4-billion-if-statements.html) blog article:

> I recently stumbled upon this screenshot while researching social media on the train. Of course, it was followed by a cascade of spiteful comments, criticizing this fresh programmer’s attempt to solve a classical problem in computer science. The modulus operation.

![image](https://andreasjhkarlsson.github.io/assets/images/GCPVDa1WYAAoBut.jpg)

This was also discussed in a [Youtube video](https://andreasjhkarlsson.github.io/assets/images/GCPVDa1WYAAoBut.jpg).

As `Go` is famous of its blazing fast compilation time, why not give a try?

# Generator


Let's create a small generator to create our `modulo`` program.

```shell
% go run cmd/generator/generator.go 3 > modulo.go
```

It should draft something like this:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <number to check>", os.Args[0])
	}

	number, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalf("Can't parse: %s to number", os.Args[1])
	}

	fmt.Println("your number ", number)

	if number == 0 {
		fmt.Println("even")
	} else if number == 1 {
		fmt.Println("odd")
	} else if number == 2 {
		fmt.Println("even")
	} else if number == 2 {
		fmt.Println("odd")
	} else if number == 3 {
		fmt.Println("even")
	}
}
```

NB: `gofmt` gave up when fixing 100.000 file....

```bash
% gofmt bzz.go
bzz.go:250009:19: exceeded max nesting depth
```

# Build

If you wish to change supported range, change the nuber in the `Makefile`, then run it:

```shell
% make
go run cmd/generator/generator.go 50000 > cmd/oddeven/oddeven.go
time go build cmd/oddeven/oddeven.go
     1003.71 real      1025.70 user        62.37 sys
```

The above were results with `50000` range.

Failed with `100000`:

```shell
go run cmd/generator/generator.go 100000 > cmd/oddeven/oddeven.go
time go build cmd/oddeven/oddeven.go
command-line-arguments: /opt/homebrew/Cellar/go/1.21.6/libexec/pkg/tool/darwin_arm64/compile: signal: killed
     5010.63 real      4809.05 user       331.20 sys
```

# Modulo

...then run it:

```shell
% ./oddeven 49999
your number  49999
odd
```
