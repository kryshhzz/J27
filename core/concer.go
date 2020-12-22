package core

import (
	"github.com/zenthangplus/goccm"
	"io/ioutil"
	"strings"
)

var list []string

func (input Details) Concer() {
	fyl, err := ioutil.ReadFile(input.File)

	if err != nil {
		CheckErr(err)
	} else {

		list = strings.Split(string(fyl), "\n")

		c := goccm.New(50)

		nbJobs := len(list)

		Heading()

		for i := 0; i < nbJobs; i++ {
			c.Wait()

			go func(input Details, i int) {
				Run(input, i)
				c.Done()
			}(input, i)
		}

		c.WaitAllDone()
		Red.Println("\nPassword not found\n")
	}
}

func Run(input Details, i int) {
	resp, err := Attack(input, list[i])

	if err != nil {
		CheckErr(err)
	} else {
		clr := Check(resp)
		Output(clr, resp, list[i], input.Word)
	}
}
