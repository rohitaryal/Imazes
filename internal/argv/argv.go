// Package is a simple utility function to parse argument values

package argv

import (
	"fmt"
	"os"
	"strings"
)

type argument struct {
	longFlag     string
	shortFlag    string
	helpMsg      string
	defaultValue string
}

var (
	_author             = ""
	_programName        = ""
	_programDescription = ""
	arguments           []argument
)

// Initialize author names, program names, etc
func Init(author, programName, programDescription string) {
	_author = author
	_programName = programName
	_programDescription = programDescription
}

// Add argument that you want to parse
func AddArg(longFlag, shortFlag, helpMsg, defaultValue string) {
	longFlag = removeDash(longFlag)
	shortFlag = removeDash(shortFlag)

	newArgument := argument{
		longFlag:     longFlag,
		shortFlag:    shortFlag,
		helpMsg:      helpMsg,
		defaultValue: defaultValue,
	}

	// Add it to our arguments
	arguments = append(arguments, newArgument)
}

// Get argument value for specified argument name
func GetArg(argumentName string) (string, error) {
	argumentName = removeDash(argumentName)
	receivedArgs := os.Args[1:] // Revieved args by program

	selectedIndex := -1 // Index where the prompt matched in arguments
	argumentValue := "" // Default value of arg

	for index, argument := range arguments {
		if argument.longFlag == argumentName || argument.shortFlag == argumentName {
			selectedIndex = index
			argumentValue = argument.defaultValue
		}
	}

	// If nothing found then return an empty string
	if selectedIndex == -1 {
		return "", fmt.Errorf("%v: NOT A REGISTERED ARGUMENT", argumentName)
	}

	for index, userArgs := range receivedArgs {
		if removeDash(userArgs) != arguments[selectedIndex].longFlag && removeDash(userArgs) != arguments[selectedIndex].shortFlag {
			continue
		}
		// If current arg is a flag and theres at least one item after it
		if removeDash(userArgs) != userArgs && index+1 < len(receivedArgs) {
			// If next item is not a flag
			if removeDash(receivedArgs[index+1]) == receivedArgs[index+1] {
				argumentValue = receivedArgs[index+1]
			}

			break
		}
	}

	return argumentValue, nil
}

// Prints help banner for program
func PrintHelp() {
	fmt.Printf("%v:\n\t%v\n\n", _programName, _programDescription)
	fmt.Printf("Usage: %v [OPTION...]\n\n", _programName)

	for _, argument := range arguments {
		fmt.Printf("  -%v,  --%v\t %v\n", argument.shortFlag, argument.longFlag, argument.helpMsg)
	}

	fmt.Printf("\nProvided By: @%v", _author)
}

// Remove dashes from a string
func removeDash(flag string) string {
	if strings.HasPrefix(flag, "--") { // Remove the '--' from a longFlag
		flag = flag[2:]
	} else if strings.HasPrefix(flag, "-") { // Remove the '-' from shortFlag
		flag = flag[1:]
	}

	return flag
}
