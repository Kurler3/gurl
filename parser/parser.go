package parser

import (
	"fmt"
	"strings"

	"github.com/Kurler3/gurl/parser/flag_parsers"
	"github.com/Kurler3/gurl/utils"
)

// Init a map between the cmd arg flag and the function that parses that flag
type FlagParser = map[string]func(string) (string, error)

//TODO - Find a way to do this without modifying the return type of the other parsers to any.

var flagParser = FlagParser{
	utils.MethodFlag: flag_parsers.ParseMethod,
	utils.UrlFlag:    flag_parsers.ParseUrl,
}

// Function that parses each cmd arg string.
func ParseCmdArg(arg string) (string, any, error) {

	// - Split the arg by =
	splitArg := strings.Split(arg, "=")

	var flag string

	// Could be a boolean flag because no need to set a value.
	if len(splitArg) >= 1 {

		var found bool

		// Get flag by removing the first 2 hiffens from the first value of the split array.
		flag, found = strings.CutPrefix(splitArg[0], "--")

		if !found {
			return "", "", fmt.Errorf("flag \"%v\" is invalid. Flags need to start with the prefix \"--\"", flag)
		}

		// If flag specified is a short version of a flag, set the flag as the long version of it.
		if _, hasLongVersion := utils.SHORT_FLAG_TO_LONG_FLAG_MAP[flag]; hasLongVersion {
			flag = utils.SHORT_FLAG_TO_LONG_FLAG_MAP[flag]
		}

		// If flag is a bool flag, then return true directly.
		if _, ok := utils.BOOL_FLAGS[flag]; ok {
			return flag, true, nil
		}

	}

	// If less than 2 items => error
	if len(splitArg) < 2 {
		return "", "", fmt.Errorf("argument \"%v\" not valid because it doesn't contain the pair flag value. Valid format: --[flag]=[value]", arg)
	}

	// Check if flag is available.
	if _, isAvailable := utils.AVAILABLE_FLAGS[flag]; !isAvailable {

		availableFlagsArr := utils.GetMapKeysAsArray(utils.AVAILABLE_FLAGS)

		return "", "", fmt.Errorf("flag \"%v\" is not available. Available flags are: \"%v\"", flag, strings.Join(availableFlagsArr, ", "))

	}

	// Check if theres an available parser for this flag.
	parserFn, ok := flagParser[flag]

	// If no parser found, simply return it without parsing
	if !ok {
		return flag, splitArg[1], nil
	}

	parsedValue, errorWhileParsing := parserFn(splitArg[1])

	if errorWhileParsing != nil {
		return "", "", fmt.Errorf("error: \"%v\" while parsing value \"%v\" for flag \"%v\"", errorWhileParsing, splitArg[1], flag)
	}

	return flag, parsedValue, nil

}
