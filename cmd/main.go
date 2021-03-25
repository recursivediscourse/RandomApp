package main

import (
	"fmt"

	"github.com/recursivediscourse/RandomApp/pkg/fairwells"
	"github.com/recursivediscourse/RandomApp/pkg/greetings"
)

func main() {
	fmt.Println("Hello World!")
	greetings.Hello()
	fairwells.Goodbye()
}
