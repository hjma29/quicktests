package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hjma29/ovcli/ovextra"
)

func main() {
	var cliOVClientPtr = ovextra.NewCLIOVClient()

	icl, err := cliOVClientPtr.GetInterconnect("", "")
	if err != nil {
		log.Fatal(err)
	}

	jsonIcl, err := json.MarshalIndent(icl, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonIcl))

}
