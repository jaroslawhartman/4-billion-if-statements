package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	tmpl_beg := `
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
		`

	tmpl_mid := `
} else if number == %d {
	fmt.Println("odd")
} else if number == %d {
	fmt.Println("even")
	`

	tmpl_end := `
	} else {
		fmt.Println("Number overflow!")
	}
}
	`

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <number to check>", os.Args[0])
	}

	number, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalf("Can't parse: %s to number", os.Args[1])
	}

	fmt.Print(tmpl_beg)
	for i := 1; i < number; i++ {
		fmt.Printf(tmpl_mid, i, i+1)
	}
	fmt.Print(tmpl_end)
}
