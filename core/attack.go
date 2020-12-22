package core

import (
	"net/http"
	"net/url"
	"strings"
	"time"
	"net"
)

func Attack(input Details, pwd string) (
	*http.Response, error) {

	var resp *http.Response
	var err error
	PD := url.Values{}

	Data := strings.Replace(input.PostData, "FUZZ", pwd, -1)

	InputData := strings.Split(Data, "&")
	for _, x := range InputData {
		two := strings.Split(x, "=")
		param := two[0]
		value := two[1]

		PD.Set(param, value)
	}


  timeout := time.Duration(time.Duration(input.Timeout) *time.Second)
  
  if err != nil{
    CheckErr(err)
  }
  
  tr := &http.Transport{
      Dial: (&net.Dialer{
      Timeout: timeout,
      }).Dial,
      TLSHandshakeTimeout: timeout,
      MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 500,
			MaxConnsPerHost:     500,
    }
  
  client := http.Client{Timeout:timeout,Transport : tr}
  
	resp, err = client.PostForm(input.Target, PD)

	return resp, err
}

func Check(resp *http.Response) Color {
	code := resp.StatusCode

	var clr Color

	switch {
	case code <= 299 && code >= 200:
		clr = Blue
	case code <= 399 && code >= 300:
		clr = Cyan
	case code <= 499 && code >= 400:
		clr = Red
	default:
		clr = Gray
	}

	return clr
}
