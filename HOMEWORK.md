# Homework


1. Clone this repo to your local machine. Get golang setup (VS Code recommended)
2. Create a new formula, YEARFRAC, following the conventions of the sum.go / sum.md file in this repo. This means it should have an interface like `YEARFRAC(args []values.Value) value.Value`
3. The YEARFRAC formula should be based on [this google sheet spec](https://support.google.com/docs/answer/3092989?hl=en)
4. The solution will consist of yearfrac.go file and a yearfrac.md file.
5. Once complete, zip up the repo and email it to me at tony@sheet.rocks.

### Other notes 
- [This link](https://en.wikipedia.org/wiki/360-day_calendar) might also be helpful.
- Follow the convention in sum.go, specifically the input should be a slice of ([]values.Value) and the output should be a value.Value
- Conceptually a `values.Value` represents a spreadsheet-native type (text, number, date, boolean, or range). Each elements in `args` is an argument into the spreadsheet formula. The output values.Value is what would get outputted into the cell. 
- Feel free to ask questions if you have them! Part of this exercise is to get a sense of what it's like to work together.
