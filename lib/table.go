package lib

import "os"

// Prints an array of titles and a 2D-array of data as a table
func PrintAsTable(titles []string, data [][]string) {
	// check for matching dimensions
	for _, element := range data {
		if len(element) != len(titles) {
			println("Table not printable.")
			os.Exit(1)
		}
	}

	// set maximum width of all rows
	var rowWidth = make([]int, len(titles))
	for index, title := range titles {
		rowWidth[index] = len(title)
	}

	// set actual width of all rows
	for _, element := range data {
		for indexData, elementData := range element {
			if len(elementData) > rowWidth[indexData] {
				rowWidth[indexData] = len(elementData)
			}
		}
	}

	// print heading
	for index, title := range titles {
		print(title)
		if len(title) != rowWidth[index] {
			for i := len(title); i < rowWidth[index]; i++ {
				print(" ")
			}
		}
		if index != len(rowWidth)-1 {
			print(" | ")
		}
	}
	println()

	// print dash-row
	for index, spaces := range rowWidth {
		if index > 0 && index < len(rowWidth)-1 {
			spaces = spaces + 1
		}

		for i := 0; i < spaces+1; i++ {
			print("-")
		}
		if index != len(rowWidth)-1 {
			print("|")
		}
	}
	println()

	// print data
	for index, element := range data {
		for indexData, elementData := range element {
			print(elementData)
			if len(elementData) != rowWidth[indexData] {
				for i := len(elementData); i < rowWidth[indexData]; i++ {
					print(" ")
				}
			}
			if indexData != len(rowWidth)-1 {
				print(" | ")
			}
		}
		if index != len(data)-1 {
			println()
		}
	}
}
