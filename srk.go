package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type FormulaBody struct {
	Name       string
	Help       string
	ScriptText string
}

func main() {
	command := os.Args[1]
	filepath := os.Args[2]
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

	dat, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error: Formula file was not found.")
		fmt.Println(err)
		return
	}

	scriptText := string(dat)

	dir, filename := path.Split(filepath)

	formulaName := strings.Split(filename, ".")[0]
	markdownPath := path.Join(dir, fmt.Sprintf("%s.md", formulaName))

	dat, err = ioutil.ReadFile(markdownPath)

	if err != nil {
		fmt.Println("Error: Help document is required.")
		fmt.Println(err)
		return
	}

	help := string(dat)

	formulaBody := FormulaBody{Name: strings.ToUpper(formulaName), Help: help, ScriptText: scriptText}

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
