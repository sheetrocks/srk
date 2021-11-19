# srk - the SheetRocks Command Line Interface

## How to make and push your own function to SheetRocks

This tutorial will walk you through how to create a custom formula and upload it to SheetRocks.
You should have at least a basic understanding of programming.

## Download the repository

If you are familiar with git, clone this repository. If not, download the zip.

To download the zip, click the green button labeled "Code" and click "Download Zip."
Extract the downloaded zip file to a destination of your choice on your computer.

## Open the repository

Open the repository folder in your text editor of choice. If you do not have one,
[Visual Studio Code](https://code.visualstudio.com/) is widely used and recommended.

To open the folder in vscode, click ```File > Open Folder...``` and navigate to the location
where you chose to clone the repository or unzip the file.

## Creating a formula

For the purposes of this tutorial, there is a completed example in the ```examples/python``` folder. 
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

Please follow the supplied convention for help files. For reference see the `help_format.md` file in this directory and the completed example `sumplusone.md` in `/examples/sumplusone`. 

- All `#` shown are required for headers. `---` is used to separate sections and is also required.

- The description for a formula must start with `Returns`.

- Optional arguments are enclosed in brackets [] and should specify default value if not given.

- In the `## Examples` section, you may include an `EXAMPLE_RANGE` that will format to appear as a spreadsheet snippet, with the top left most value appearing in cell `A1`. You can then reference these cells in your examples to show how the formula would be used in a sheet.

- In the `## Related Formulas` section, urls should be formatted as `formulas/relatedformulaname`.

## Creating a config file

You must include with your formula and help document a `config.json` file. This file includes metadata about your formula and the paths to the required files. See `config_template.json` in this directory and the completed example `config.json` in `/examples/python/sumplusone`

## Push the function to Sheetrocks

We recommend pushing the example located at `/examples/sumplusone` to make sure you understand how to upload new formulas before you try uploading your own code.

First, you'll need to log in and create an API key (open SheetRocks, log in, create a workbook, click profile, and click create API key).
Copy the key to your clipboard.

You then need to push the formula using the terminal/command prompt.

First, navigate to the directory where you originally cloned the repository or unzipped the file you downloaded.

On windows:
```
C:\>cd C:\path\to\srk
```

On Linux:
```
$ cd /path/to/srk
```

Next, you need to set your API key.

On windows:
```
C:\path\to\srk>setx SRK_TOKEN "your_api_key_here"
```

On Linux:
```
$ export SRK_TOKEN="your_api_key_here"
```

Then, push the formula to SheetRocks.

Windows:
```
C:\path\to\srk>srk.exe push /path/to/config.json
```

Linux:
```
$ ./srk push /path/to/config.json
```

If successful, it will display a success message. If not, it will display an error.
Once your formula is loaded, you can immediately visit a SheetRocks sheet and use your new formula.
If there is an error running your new code with a certain input, it should display in the cell. 


