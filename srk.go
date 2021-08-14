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
	"time"

	"github.com/briandowns/spinner"

	v "github.com/sheetrocks/srk/values"
)

type FormulaBody struct {
	Name         string
	Runtime      string
	Help         string
	ScriptText   string
	Dependencies string
}

type Asset struct {
	Filename string
	MimeType string
	Content  string
}

type Parameter struct {
	Type  string
	Title string
}

type ParameterJSON struct {
	Type  v.ValueType
	Title string
}

type ChartTypeSubmission struct {
	Title      string
	Parameters []Parameter
	Assets     []Asset
	Help       string
}

func (p *Parameter) MarshalJSON() ([]byte, error) {
	var t v.ValueType

	switch p.Type {
	case "number":
		t = v.NUMBER
		break
	case "text":
		t = v.TEXT
	case "date":
		t = v.DATE
	case "boolean":
		t = v.BOOLEAN
	}
	return json.Marshal(ParameterJSON{t, p.Title})
}

type ParametersConfig struct {
	Version    string
	Parameters []Parameter
}

type ChartType struct {
	ID    string
	Title string
}

type Config struct {
	Name         string
	Runtime      string
	Formula      string
	Help         string
	Dependencies string
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

	_, err := os.Stat(filepath)

	if err != nil {
		fmt.Println("Error: configuration file path was not recognized. Please check the file path.")
		fmt.Println(err)
		return
	}

	dat, err := ioutil.ReadFile(filepath)

	config := Config{}
	err = json.Unmarshal(dat, &config)

	if err != nil {
		fmt.Println("Error: could not parse configuration file.")
		fmt.Println(err)
		return
	}

	if config.Name == "" {
		fmt.Println("Error: config file must have name field.")
		return
	}

	if config.Runtime != "python" {
		fmt.Println(`Error: invalid runtime field specified. Accepted runtimes: "python"`)
		return
	}

	dir := path.Dir(filepath)

	scriptText, err := ioutil.ReadFile(path.Join(dir, config.Formula))

	if err != nil {
		fmt.Printf("Error: could not find formula script located at %s\n", config.Formula)
		return
	}

	helpText, err := ioutil.ReadFile(path.Join(dir, config.Help))

	if err != nil {
		fmt.Printf("Error: could not find help file located at %s\n", config.Help)
		return
	}

	dependenciesText, err := ioutil.ReadFile(path.Join(dir, config.Dependencies))

	if err != nil {
		fmt.Printf("Error: could not find dependencies file located at %s\n", config.Dependencies)
		return
	}

	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Color("fgBlack")
	s.FinalMSG = ""
	s.Start()

	formulaBody := FormulaBody{Name: strings.ToUpper(config.Name), Runtime: config.Runtime, Help: string(helpText), ScriptText: string(scriptText), Dependencies: string(dependenciesText)}

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

	s.Stop()
	if resp.StatusCode == 200 {
		fmt.Printf("🎉 Success! You have pushed your formula \"%s\" to SheetRocks 🎉\n", formulaBody.Name)
	} else {
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("SheetRocks error: ", string(bytes))
	}
}
