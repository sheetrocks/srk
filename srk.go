package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
		fmt.Println("Error: file path was not recognized. Please check the file path.")
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

	scriptText, err := ioutil.ReadFile(config.Formula)

	if err != nil {
		fmt.Printf("Error: could not find formula script located at %s\n", config.Formula)
		return
	}

	helpText, err := ioutil.ReadFile(config.Help)

	if err != nil {
		fmt.Printf("Error: could not find help file located at %s\n", config.Help)
		return
	}

	dependenciesText, err := ioutil.ReadFile(config.Dependencies)

	if err != nil {
		fmt.Printf("Error: could not find dependencies file located at %s\n", config.Dependencies)
		return
	}

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

	if resp.StatusCode == 200 {
		fmt.Printf("🎉 Success! You have pushed your formula \"%s\" to SheetRocks 🎉\n", formulaBody.Name)
	} else {
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("SheetRocks error: ", string(bytes))
	}
	/*
		if info.IsDir() {
			files, _ := ioutil.ReadDir(filepath)

			chartSubmission := ChartTypeSubmission{}
			chartSubmission.Assets = []Asset{}
			hasIndex := false
			hasHelp := false
			hasConfig := false

			_, f := path.Split(filepath)
			chartSubmission.Title = f
			for _, file := range files {
				if file.Name() == "index.html" {
					hasIndex = true
				}
				switch file.Name() {
				case "parameters.json":
					parameterPath := path.Join(filepath, file.Name())
					dat, err := ioutil.ReadFile(parameterPath)
					var parametersConfig ParametersConfig
					err = json.Unmarshal(dat, &parametersConfig)

					if err != nil {
						fmt.Println("Error encountered while trying to read parameters.json file")
						fmt.Println(err)
						return
					}
					chartSubmission.Parameters = parametersConfig.Parameters
					hasConfig = true
					break
				case "help.md":
					dat, _ := ioutil.ReadFile(path.Join(filepath, file.Name()))
					chartSubmission.Help = string(dat)
					hasHelp = true
					break
				default:
					dat, _ := ioutil.ReadFile(path.Join(filepath, file.Name()))
					mimeType := ""

					switch path.Ext(file.Name()) {
					case ".html":
						mimeType = "text/html"
						break
					case ".js":
						mimeType = "application/javascript"
						break
					case ".css":
						mimeType = "text/plain"
						break
					}

					if mimeType != "" {
						chartSubmission.Assets = append(chartSubmission.Assets, Asset{file.Name(), mimeType, string(dat)})
					}

				}
			}

			if !hasIndex {
				fmt.Println("Missing required file: index.html")
				return
			}

			if !hasConfig {
				fmt.Println("Missing required file: parameters.json")
				return
			}

			if !hasHelp {
				fmt.Println("Missing required file: help.md")
				return
			}


			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/chart-type", baseUrl), nil)

			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			req.Header.Add("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			if resp.StatusCode == 200 {
				chartTypes := []ChartType{}
				json.NewDecoder(resp.Body).Decode(&chartTypes)
				resp.Body.Close()

				body, _ := json.Marshal(chartSubmission)
				r := bytes.NewReader(body)
				foundChartID := ""

				for _, ct := range chartTypes {
					if ct.Title == chartSubmission.Title {
						foundChartID = ct.ID
					}
				}

				var outputType string
				if foundChartID == "" {
					outputType = "added"
					req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/chart-type", baseUrl), r)
				} else {
					outputType = "updated"
					req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("%s/chart-type/%s", baseUrl, foundChartID), r)
				}

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
				req.Header.Add("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					panic(err)
				}

				if resp.StatusCode == 200 {
					fmt.Printf("🎉 Successfully %s new chart \"%s\" 🎉\n", outputType, chartSubmission.Title)
				} else {
					fmt.Printf("Encountered unexpected status code: %d\n", resp.StatusCode)
				}
			}
		} else {
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
				fmt.Printf("🎉 Success! You have pushed your formula \"%s\" to SheetRocks 🎉\n", formulaBody.Name)
			} else {
				bytes, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("SheetRocks error: ", string(bytes))
			}
		}
	*/
}
