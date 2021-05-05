# srk - the SheetRocks Command Line Interface


## Homework

1. Clone this repo to your local machine. Get golang setup (VS Code recommended)
2. Create a new formula, YEARFRAC, following the conventions of the echo.go / echo.md file in this repo.
3. The YEARFRAC formula should be based on [this google sheet spec](https://support.google.com/docs/answer/3092989?hl=en)
4. The solution will consist of yearfrac.go file and a yearfrac.md file.
5. Once complete, zip up the repo and email it to me.

### Other notes 
- [This link](https://en.wikipedia.org/wiki/360-day_calendar) might also be helpful.
- Follow the convention in echo.go, specifically the input should be a slice of ([]values.Value) and the output should be a value.Value


## Using srk - testers only
First, you'll need to log in and create an API key (you do this by creating a workbook, clicking profile, and clicking create API key).

Copy the key to your clipboard. On a Mac, you can run

```
$ SRK_TOKEN={your_api_token} ./srk push echo.go
```

If successful, it will display a success message. If not, it will display an error.

Using windows, run
```
C:\path\to\srk>setx SRK_TOKEN "your_api_token"
C:\path\to\srk>./srk.exe push echo.go 
```

## Creating your custom formula
- In order to work, your custom formula should keep the same func signature as the echo.go example, namely it should accept a values slice ([]values.Value) and return a values.Value.
- There is a little magic that happens with the name. In order to work, please name the method signature, filename, and help file the same thing (with func capitalized). For example to create a Multiply function:
    * Name of file is multiply.go
    * Name of help file is multiply.md
    * Function signature is `func Multiply(v []values.Value) values.Value { ... }`
- The help file will display in autocomplete. Please follow the convention of echo.md, with the first line containing the syntax of the formula, followed by a brief summmary of functionality in the next paragraph, and any additional information afterwards.
