package ui

import "fmt"

func ShowPreview(originalMessage string, improvedMessage string) {
	fmt.Println("ORIGINAL MESSAGE:")
	fmt.Println(originalMessage)

	fmt.Println()
	fmt.Println("AI SUGGESTION:")
	fmt.Println(improvedMessage)
	fmt.Println()
}
