# srk - the SheetRocks Command Line Interface

## How to make and push your own function to SheetRocks

This tutorial will walk you through how to create a custom function and upload it to SheetRocks.
Hopefully you have at least a basic understanding of programming, but I will try to make this
as accessible as possible to any skill level.

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

For the purposes of this tutorial, there are some completed examples in the ```examples/python``` folder. 
To make a new formula, follow the conventions of the examples:
- A config.json file is required. Please use that file to define your formula name and the file path to the required files.
- The python formula must be called `calculate`
- The python function will be sent a list of arguments to you formula, cast as a spreadsheet castable type:
    * None (for empty cells)
    * float
    * string
    * datetime
    * bool
    * A two dimensional list of the primative types above (this represents a range such as A1:C3).
- The return value must be a spreadsheet castable type as defined above. If the output is a 2D list, it will overwrite
cells.

## Creating a help file

Please follow the supplied convention for help files:
- First line should be signature of the formula, e.g., SUM(arg1, arg2, ..argN). This will display in autocommplete.
- Second line should be 1-sentence description of the formula. 
- Next paragraph(s) can be long form explanation of the formula.

## Push the function to Sheetrocks

We recommend pushing the example located at `/examples/echo` to make sure you understand how to upload new formulas, before you try uploading your own code.

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


