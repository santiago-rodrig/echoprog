// Prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(os.Args[0])
	for _, arg := range os.Args[1:] {
		fmt.Print(" ", arg)
	}
	fmt.Println()
}
