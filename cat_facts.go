package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// custom data type
type CatFact struct {
	Fact   string
	Length int
}

func main() {
	// fmt.Println("hello world")
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()
	// loose connection to the url

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var fact CatFact

	json.Unmarshal(body, &fact)

	// fmt.Println(resp)
	// fmt.Printf("Resp: %#v", string(body))
	// fmt.Printf("Resp: %#v", fact.Fact)
	fmt.Printf(fact.Fact)
	// decode the JSON

}
