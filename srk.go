package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type Runtime struct {
	Runtime string
}

type WebConfig struct {
	Name       string
	Runtime    string
	Help       string
	Parameters []Parameter
	Assets     []string
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Error: not enough arguments.\nExpected command format: './srk push ./path/to/config.json'")
	}

	command := os.Args[1]
	filepath := os.Args[2]
	token := os.Getenv("SRK_TOKEN")
	baseUrl := os.Getenv("SRK_BASE_URL")

	if baseUrl == "" {
		baseUrl = "https://sheet.rocks/api/v1"
	}

	if token == "" {
		log.Fatal("You must provide your API token in the SRK_TOKEN environment variable.")
	}

	if command != "push" {
		log.Fatal("Command not found. Supported commands: push")
	}

	_, err := os.Stat(filepath)

	if err != nil {
		log.Fatal("Error: configuration file path was not recognized. Please check the file path.\n\n", err)
		return
	}

	dir := path.Dir(filepath)

	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	runtime := Runtime{}
	err = json.Unmarshal(dat, &runtime)
	if err != nil {
		log.Fatal(err)
	}

	if runtime.Runtime == "" {
		log.Fatal("Error: must specify runtime.")
	} else if !(runtime.Runtime == "python" || runtime.Runtime == "web") {
		log.Fatal(`Error: invalid runtime. Accepted runtimes: "python", "web"`)
	}

	if runtime.Runtime == "python" {
		config := Config{}
		err = json.Unmarshal(dat, &config)

		if err != nil {
			log.Fatal("Error: could not parse configuration file.\n\n", err)
		}

		if config.Name == "" {
			log.Fatal("Error: config file must have name field.")
		}

		if config.Runtime != "python" {
			log.Fatal(`Error: invalid runtime field specified. Accepted runtimes: "python"`)
		}

		scriptText, err := ioutil.ReadFile(path.Join(dir, config.Formula))

		if err != nil {
			log.Fatal("Error: could not find formula script located at %s\n", config.Formula)
		}

		helpText, err := ioutil.ReadFile(path.Join(dir, config.Help))

		if err != nil {
			log.Fatal("Error: could not find help file located at %s\n", config.Help)
		}

		dependenciesText, err := ioutil.ReadFile(path.Join(dir, config.Dependencies))

		if err != nil {
			log.Fatal("Error: could not find dependencies file located at %s\n", config.Dependencies)
		}

		s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
		s.Color("fgBlack")
		s.FinalMSG = ""
		s.Start()

		formulaBody := FormulaBody{Name: strings.ToUpper(config.Name), Runtime: config.Runtime, Help: string(helpText), ScriptText: string(scriptText), Dependencies: string(dependenciesText)}

		body, _ := json.Marshal(formulaBody)

		r := bytes.NewReader(body)

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/formula", baseUrl), r)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Unhandled error submitting to SheetRocks: ", err)
		}

		s.Stop()
		if resp.StatusCode == 200 {
			fmt.Printf("🎉 Success! You have pushed your formula \"%s\" to SheetRocks 🎉\nThe formula is available for immediate use in a SheetRocks sheet.\n", formulaBody.Name)
		} else {
			bytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("SheetRocks error: ", string(bytes))
		}
	}

	if runtime.Runtime == "web" {
		webConfig := WebConfig{}
		err = json.Unmarshal(dat, &webConfig)

		if err != nil {
			log.Fatal("Error: could not read configuration file for chart.")
		}

		if webConfig.Name == "" {
			log.Fatal(`Error: configuration file must have a "name" field which specifies the name of the chart.`)
		}

		if webConfig.Help == "" {
			log.Fatal("Error: must specify path of help file.")
		}

		helpText, err := ioutil.ReadFile(path.Join(dir, webConfig.Help))

		if err != nil {
			log.Fatal("Error: could not read help file: ", err)
		}

		hasIndex := false
		for _, filename := range webConfig.Assets {
			if filename == "index.html" {
				hasIndex = true
			}
		}

		if !hasIndex {
			log.Fatal(`Error: there must be an "index.html" file listed in assets.`)
		}

		assets := []Asset{}

		for _, assetPath := range webConfig.Assets {
			assetDir := path.Dir(assetPath)

			if assetDir != "." {
				log.Fatal("Error: assets must be in same directory as config file. Asset not in config directory: ", assetPath)
			}
			dat, err := ioutil.ReadFile(path.Join(dir, assetPath))

			if err != nil {
				log.Fatal("Error: could not read asset located at ", assetPath)
			}
			mimeType := ""

			switch path.Ext(assetPath) {
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
				assets = append(assets, Asset{assetPath, mimeType, string(dat)})
			}
		}
		chartSubmission := ChartTypeSubmission{Title: webConfig.Name, Help: string(helpText), Parameters: webConfig.Parameters}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/chart-type", baseUrl), nil)

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Unhandled error encountered while submitting to SheetRocks: ", err)
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
				log.Fatal("Unhandled error encountered while submitting to SheetRocks: ", err)
			}

			if resp.StatusCode == 200 {
				fmt.Printf("🎉 Successfully %s new chart \"%s\" 🎉\n", outputType, chartSubmission.Title)
			} else {
				fmt.Printf("Encountered unexpected status code: %d\n", resp.StatusCode)
			}
		}
	}
}
