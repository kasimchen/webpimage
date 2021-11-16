package main

import (
	"fmt"
	"net/http"
)

func main() {


	//c := &http.Client{
	//	Timeout: 5 * time.Second,
	//}

	resp, err := http.Get("https://secure.gravatar.com/avatar/d41d8cd98f00b204e9800998ecf8427e?s=65&amp;r=G&amp;d=")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Header)




}
