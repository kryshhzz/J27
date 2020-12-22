package core

import (
	"fmt"
)

type Details struct {
	Target   string
	PostData string
	File     string
	Word     string
	Timeout  int
}

func CheckErr(err error) {
	Red.Println("Error")
	fmt.Println(err)
}
