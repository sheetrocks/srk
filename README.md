# srk - the SheetRocks Command Line Interface

---

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

## Create your function

For the purposes of this tutorial, there is a completed example in the ```formulas``` folder
called ```sum.go```. 

Every formula accepts in input of ```[]values.Value``` which represents a list of arguments
that the formula needs. Each ```values.Value``` passed to the function can have a type of
```EMPTY```, ```NUMBER```, ```DATE```, ```BOOLEAN```, ```TEXT```, or ```ARRAY```.

For our example SUM function, we will need an ```ARRAY``` representing multiple values that
we can return the sum of. This function would only have one argument passed, but other functions may need
more than one argument. Whether or not the function has multiple arguments, the input is always
an array of values.

For example, if we have 3 numbers in cell A1, A2, and A3, we input them into the formula as ```SUM(A1:A3)```
rather than ```SUM(A1, A2, A3)```.

For our function, the user will input a list of cells they want to get the sum of, similar to the
example above. They will not type in the cell of each individual number, rather they will use ```SUM(START:FINISH)```
which will pass a ```values.Value``` of type ```ARRAY``` inside of a ```[]values.Value``` where it is the
only item because the function only accepts one argument.

Follow along the ```formulas/sum.go``` file to see how the formula is created. There are comments in the file that 
will guide you through the steps that were taken to create the formula and the thought process behind it. The purpose
of the function is to help you understand the inputs and outputs and how they relate to SheetRocks.

To create your formula, you can use the template ```myformula.go``` file. Rename the file by replacing ```myformula``` 
with the actual name of your function in all lowercase. This file includes 
the package and import declarations at the top, which are required. In the file, there is a placeholder function
called ```Myfunction```. Change the name of this to the name of your function, with the first letter being capital
and the rest lowercase.

## Create your .md file

For every function that is to be uploaded, there needs to be an additional file titled ```yourfunction.md``` which
gives information about the function and the arguments it accepts. This file needs to be in the same folder as your ```myfunction.go``` 
file. This information is what pops up when you start 
typing a function in the spreadsheet. Open the ```sum.md``` file and you can see the example for the SUM function.
There is a template file you can use for your function called ```myformula.md``` where you can replace the file name, function name,
arguments, and description with what is appropriate for your function. Make sure the function name inside of the file is in all caps.

## Push the function to Sheetrocks

At this point, you should have two files. For example, if the name of your function was ```Average```, you should have a
```average.go``` with a function ```func Average(v []values.Value) values.Value``` and an ```average.md``` markdown file.

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
C:\path\to\srk>./srk.exe push ./formulas/yourformulaname.go
```

Linux:
```
$ ./srk push ./formulas/yourformulaname.go
```

If successful, it will display a success message. If not, it will display an error.


