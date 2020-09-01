package main

import (
	"fmt"
	"os"
	"io"
)

// Greet writes a Hello-Message into a given object that implements the Writer-Interface (os.Stdout, bytes.Buffer, ...)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}