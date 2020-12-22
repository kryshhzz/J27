package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var exit bool = false

func Output(clr Color, resp *http.Response, pwd, word string) {

	cont, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		CheckErr(err)
	} else {

		if strings.Contains(string(cont), word) {
			clr = Green
			defer func(){
			  Green.Println("\n Password :: "+pwd+"\n")
			  os.Exit(0)
			}()
		}

		fmt.Print(clr.Code)
		fmt.Printf(" %-20s", resp.Status)
		fmt.Printf(" %-6d", len(cont))
		fmt.Printf(" %-15s",pwd)
		fmt.Println(Reset.Code)

	}

}
