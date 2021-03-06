package utils

import (
	"fmt"
	"io"
)

// SingleOrListIntf function converts map[string]interface{} objects into an interface slice
func MapToIntfSlice(fileMap map[string]interface{}, key string) []interface{} {
	var interfaceSlice []interface{}
	if fileMap[key] == nil {
		interfaceSlice = append(interfaceSlice, fileMap)
	} else {
		interfaceSlice = fileMap[key].([]interface{})
	}
	return interfaceSlice
}

// Print a message to an io.Writer
func PrintMessage(w io.Writer, message string) {
	fmt.Fprintln(w, message)
}
