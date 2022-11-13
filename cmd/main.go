package main

import (
	"fmt"
	"github.com/lgzzzz/LGZ/pkg"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	if os.Args[1] == "newpassword" {
		str := pkg.NewPassword()
		fmt.Println(str)
	}

}
