# srk - the SheetRocks Command Line Interface

## How to make and push your own function to SheetRocks

This tutorial will walk you through how to connect your dev environment to SheetRocks and how to push a pre-existing custom formula. This task is not for beginners—it will be helpful if you have at least a basic understanding of programming! At the end of this doc there is some bonus content about the nuances of writing your own custom formulas.

## 1. Download the repository

If you are familiar with git, clone this repository. If not, download the zip.

To download the zip, click the green button labeled "Code" and click "Download Zip." Extract the downloaded zip file to a destination of your choice on your computer. In this tutorial the folder will be downloaded to the computer’s desktop.

<img src="img/download.png" width=25% height=25%>

## 2. Open the repository

Open the repository folder in your text editor of choice. If you do not have one,
[Visual Studio Code](https://code.visualstudio.com/) is widely used and recommended.

Open the zipped folder to unzip it and then open the folder in VSCode. To do this, click `File > Open Folder` and select the entire folder that was unzipped.

![vscode](img/vscode.png)

## 3. Generate an API key

First, you'll need to log in and create an API key (open SheetRocks, log in, create a workbook, hover on Profile, click API keys, and click Create API key). Copy the key to your clipboard.

![sheet](img/sheet.png)

![api key](img/apikey.png)

## 4. Set your API key

On Mac:
```
export SRK_TOKEN="your_api_key_here"
```

On Windows (Command Prompt):
```
setx SRK_TOKEN "your_api_key_here"
```

On Windows (PowerShel):
```
$env:SRK_TOKEN="your_api_key_here"
```

On Linux:
```
export SRK_TOKEN="your_api_key_here"
```

## 5. Setup your terminal

Navigate to the directory where you originally cloned the repository or unzipped the file you downloaded.

On Mac:
```
cd /path/to/srk
```

On Windows:
```
cd C:\path\to\srk
```

On Linux:
```
cd /path/to/srk
```

## 6. Push the function to SheetRocks

Push the example formula located at `/examples/python/sumplusone` to SheetRocks.

On Mac:
```
./srk.mac push ./examples/python/sumplusone/config.json
```

On Windows (Command Prompt):
```
srk.exe push ./examples/python/sumplusone/config.json
```

On Windows (PowerShell):
```
./srk.exe push ./examples/python/sumplusone/config.json
```

On Linux:
```
./srk push /path/to/config.json
```

## 7a. Success

If successful, you will see
```
🎉 Success! You have pushed your formula "SUMPLUSONE" to SheetRocks 🎉
```
Once your formula is loaded, you can immediately visit a SheetRocks sheet and use your new formula. 

## 7b. Troubleshooting
On a Mac you might receive the following error message. 

![error message](img/error1.png)

To get around this, you will need to allow the download from the system preferences.
```
Apple menu>System Preferences>Security & Privacy>General>Allow Anyway
```

![system preferences](img/error2.png)
![security and privacy](img/error3.png)


Run the `push` command again.
```
./srk.mac push ./examples/python/sumplusone/config.json
```

Hit `open` when asked if you want to open `srk.mac`.

![allow](img/error4.png)

---

# More Information About Writing Custom Formulas

## Creating a formula

For the purposes of this tutorial, there is a completed example in the `examples/python` folder. 
To make a new formula, follow the conventions of the example:
- The python formula must be called `calculate`
- The python function will be sent a list of arguments, cast as a spreadsheet castable type:
    * None (for empty cells)
    * float
    * string
    * datetime
    * bool
    * A two dimensional list of the primative types above (this represents a range such as A1:C3).
- The return value must be a spreadsheet castable type as defined above. If the output is a 2D list, it will overwrite
cells.

## Creating a help file

Please follow the supplied convention for help files. For reference see the `help_format.md` file in this directory and the completed example at `/examples/sumplusone/sumplusone.md`. 

- All `#` shown are required for headers. `---` is used to separate sections and is also required.

- The description for a formula must start with `Returns`.

- Optional arguments are enclosed in brackets [] and should specify default value if not given.

- In the `## Examples` section, you may include an `EXAMPLE_RANGE` that will format to appear as a spreadsheet snippet, with the top left most value appearing in cell `A1`. You can then reference these cells in your examples to show how the formula would be used in a sheet.

- In the `## Related Formulas` section, urls should be formatted as `formulas/relatedformulaname`.

## Creating a config file

You must include with your formula and help document a `config.json` file. This file includes metadata about your formula and the paths to the required files. See `/templates/config_template.json` and the completed example at `/examples/python/sumplusone/config.json`.

