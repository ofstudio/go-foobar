package main

import (
	"fmt"

	"github.com/ofstudio/go-foobar"
	"github.com/ofstudio/go-foobar/internal/hello"
	"github.com/ofstudio/go-foobar/pkg/privet"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(hello.Hi())
	fmt.Println(privet.PrivetMir())
	fmt.Println(foobar.Aloha())
	fmt.Println("Bye!")
}
