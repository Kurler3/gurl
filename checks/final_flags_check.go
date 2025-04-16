package checks

import (
	"net/http"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/utils"
)

var REQUIRED_FLAGS = map[string]struct{}{
	utils.UrlFlag:    {},
	utils.MethodFlag: {},
}

var DEFAULT_FLAG_VALUES = map[string]string{
	utils.MethodFlag:   http.MethodGet,
	utils.ProtocolFlag: utils.HTTPS,
}

func FinalFlagsCheck(g *gurl.Gurl) error {

	//TODO For each flag in the default flag map => set it on the g instance.

	//TODO Check if every required flag is defined.

	return nil
}

// Have a single function that handles all the final flag checks.
// 	For each flag that has default values, need to set them. How to access the correct attribute on the g instance?
//	For each flag in the required flags, we need to check whether they exist.
