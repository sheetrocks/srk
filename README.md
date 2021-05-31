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

For the purposes of this tutorial, there is a completed example in the ```formulas``` folder
called ```sum.go```. There are comments in the file that will guide you through the steps that 
were taken to create the formula and the thought process behind it. Your custom formula will need
to be in a file titled ```yourformulaname.go``` and follow the format that is described in the
```sum.go``` example. You can put your file in any directory you prefer, as long as it and the help
file (next section) are in the same folder. The command line instruction at the end of the tutorial
assume that the files are in the ```srk/formulas``` folder.

## Creating a help file

For every formula that is to be uploaded, there needs to be an additional file titled ```yourformulaname.md``` which gives information about the function and the arguments it accepts. This file needs to be in the same folder as your ```yourformulaname.go``` file. The information in the help file is what pops up when you start typing a formula name in the spreadsheet. Open the ```sum.md``` file and you can see the example for the SUM function. Make sure the function name inside of the file is in all caps.

## Push the function to Sheetrocks

Before uploading a formula, you should have two files. For example, if you created a formula called ```Average```, you should have an ```average.go``` file with a function that has the signature ```func Average(v []values.Value) values.Value```, and an ```average.md``` markdown file.

You'll need to log in and create an API key (open SheetRocks, log in, create a workbook, click profile, and click create API key).
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
C:\path\to\srk>srk.exe push formulas/yourformulaname.go
```

Linux:
```
$ ./srk push ./formulas/yourformulaname.go
```

If successful, it will display a success message. If not, it will display an error.
Note: It is currently possible to push a formula that doesn't meet the required specifications. This means that you could see a success message for a formula that does not function properly. This will be fixed in the future.


