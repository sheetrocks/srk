package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type FormulaBody struct {
	ScriptText string
}

func main() {
	command := os.Args[1]
	filename := os.Args[2]
	token := os.Getenv("SRK_TOKEN")
	baseUrl := os.Getenv("SRK_BASE_URL")

	if baseUrl == "" {
		baseUrl = "https://sheet.rocks/api/v1"
	}

	if token == "" {
		fmt.Println("You must provide your API token in the SRK_TOKEN environment variable.")
	}

	if command != "push" {
		fmt.Println("Command not found. Supported commands: push")
		return
	}

	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	formulaBody := FormulaBody{string(dat)}

	body, _ := json.Marshal(formulaBody)

	r := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/formula", baseUrl), r)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Success! You have pushed your new formula to SheetRocks.")
	} else {
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("SheetRocks error: ", string(bytes))
	}
}
