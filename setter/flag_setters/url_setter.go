package flag_setters

import (
	"fmt"
	"strings"

	"github.com/Kurler3/gurl/classes"
	"github.com/Kurler3/gurl/utils"
)

func UrlSetter(g *classes.Gurl, val string) error {

	var protocolInUrl string

	if strings.HasPrefix(val, "http://") {
		protocolInUrl = "http"
	}

	if strings.HasPrefix(val, "https://") {
		protocolInUrl = "https"
	}

	// If there is no protocol in the url, save it as is.
	if protocolInUrl == "" {
		g.Url = val
		return nil
	}

	// If has protocol and there's already a protocol as well defined, check if the same.
	if g.Protocol != "" && g.Protocol != protocolInUrl {
		return fmt.Errorf("protocol specified in the flag \"%v\" not the same as the one specified in the url. Please either remove the one in the url or specify the same one in the \"%v\" flag", utils.ProtocolFlag, utils.ProtocolFlag)
	}

	g.Protocol = protocolInUrl

	// Strip the protocol from the url and save it.
	urlWithoutProtocol := strings.TrimPrefix(val, protocolInUrl)

	// Set the url without protocol.
	g.Url = urlWithoutProtocol

	return nil

}
