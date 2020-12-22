package core

import (
	"fmt"
)

type Hue interface {
	Print()
	Println()
}

type Color struct {
	Code string
}

var Red Color = Color{"\033[1;31m"}
var Blue Color = Color{"\033[0;34m"}
var Green Color = Color{"\033[1;32m"}
var Cyan Color = Color{"\033[1;36m"}
var Gray Color = Color{"\033[1;30m"}
var Reset Color = Color{"\033[0m"}

func (clr Color) Print(txt string) {
	fmt.Print(clr.Code + txt + Reset.Code)
}

func (clr Color) Println(txt string) {
	fmt.Println(clr.Code + txt + Reset.Code)
}

func Heading() {
	fmt.Println()
	fmt.Print(Blue.Code)
	fmt.Printf(" %-18s", "Status")
	fmt.Printf(" %-8s", "Length")
	fmt.Printf(" %-14s\n","Load")

	fmt.Print(Reset.Code)

	fmt.Print(Gray.Code)
	fmt.Printf(" %-16s", "-----------")
	fmt.Printf("%-8s", "-----------")
	fmt.Printf("%-13s\n","-----------")
}

func Banner() {

	Red.Println(`
  
    ::::::::::: ::::::::  ::::::::::: 
        :+:    :+:    :+: :+:     :+: 
        +:+          +:+         +:+  
        +#+        +#+          +#+   
        +#+      +#+           +#+    
    #+# #+#     #+#           #+#     
     #####     ##########     ###     
  `)

	fmt.Print(Blue.Code)
	fmt.Printf("%30s\n", " A neonified tool  ")
	fmt.Print(Gray.Code)
	fmt.Printf("%28s\n", "------------------")
	fmt.Print(Reset.Code)
}
