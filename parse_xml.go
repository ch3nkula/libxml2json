/**
 * Parse XML program will get an XML file or XML data and parse it to
 * get the structure and understand it. Also, it can rebuild this data
 * after parsing into different formats.
 */

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// the print element function
func printElmt(s string, depth int) {
	for n := 0; n < depth; n++ {
		fmt.Print(" ")
	}

	fmt.Println(s)
}

// checking of errors
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// main function definition
func main() {

	// check the command syntax
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "file")
		os.Exit(1)
	}

	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	checkError(err)
	r := strings.NewReader(string(bytes))

	// using the xml parser
	parser := xml.NewDecoder(r)
	depth := 0

	for {
		token, err := parser.Token()
		if err != nil {
			break
		}

		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
			depth++
		case xml.EndElement:
			depth--
			elmt := xml.EndElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
		case xml.CharData:
			bytes := xml.CharData(t)
			printElmt("\""+string([]byte(bytes))+"\"", depth)
		case xml.Comment:
			printElmt("Comment", depth)
		case xml.ProcInst:
			printElmt("ProcInst", depth)
		case xml.Directive:
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}
