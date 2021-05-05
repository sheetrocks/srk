package formulas

import (
	"github.com/sheetrocks/srk/values"
)

//                     Telling the formula what type our input will be
// 					    v
// Naming our input	    v			v < < Telling our formula what type the output will be
// 		 v              v  			v
func Sum(v []values.Value) values.Value {
	// Because we only have one input, the length of input will be 1.
	// Let's retrieve the first value of our input and assign it to a variable.

	list := v[0].Array
	// In most programming languages, 0 represents the first value
	// instead of 1.
	// To access the array that was passed, we use .Array.
	// If the value we were expected was a single number, we would use .Number, or if
	// it was text we would use .Text
	//
	// Our variable list represents a [][]values.Value. If someone input into our function
	// A1:A3, we would access those values as follows:
	//
	//       ROW  > > > > >    < < < < < Column
	//                     v  v
	// A1 would be at list[0][0]
	// A2 would be at list[1][0]
	// A3 would be at list[2][0]
	//
	// Another example, this time the user wants to calculate the sum of B4:B7
	//
	//       ROW  > > > > >    < < < < < Column
	//                     v  v
	// B4 would be at list[0][0]
	// B5 would be at list[1][0]
	// B6 would be at list[2][0]
	// B7 would be at list[3][0]
	//
	// Notice how accessing the values was similar regardless of what column or row
	// the user input.

	output := 0.0

	for i := range list {
		// Here, we are getting the length of our list, and for each row
		// we are getting the first value. This is essentially like iterating
		// down a column, item by item. at every list[i][0] there is a values.Value
		// representing a cell, and we use .Number to access the number in that cell.
		output = output + list[i][0].Number
	}

	// Here we convert our variable output, which is a float64, into a values.Value
	// to be returned. All outputs have to be in value.Value form. All numbers in
	// a values.Value object have to be float64. If we were writing another function
	// that returns a text value, our return statement might look something
	// like:
	//
	// return values.Value{
	// 		Type: values.TEXT,
	// 		Text: output,
	// }
	return values.Value{
		Type:   values.NUMBER,
		Number: output,
	}

}
