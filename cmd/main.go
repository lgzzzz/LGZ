package main

import (
	"fmt"
	"github.com/lgzzzz/lgzkit/pkg"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	if os.Args[1] == "newpassport" {
		str := pkg.NewPassport()
		fmt.Println(str)
	}

}
