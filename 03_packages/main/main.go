package main

import (
	"fmt"

	"github.com/AndrewThian/udemyGolang/03_packages/otherpackage"
	"github.com/AndrewThian/udemyGolang/03_packages/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,tahW"))
	fmt.Println(stringutil.MyName)
	fmt.Println(chickendinnerpackage.Winner)
}
