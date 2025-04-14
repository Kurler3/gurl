package flag_parsers

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/Kurler3/gurl/utils"
)

// Parser of the method flag.
func ParseUrl(value string) (string, error) {

	urlToCheck := value

	// If no protocol in the url, default to https.
	if !strings.Contains(urlToCheck, "://") {
		urlToCheck = utils.HTTPS + "://" + urlToCheck
	}

	fmt.Println("url to check: ", urlToCheck)

	u, err := url.ParseRequestURI(urlToCheck)

	fmt.Println("Check of url: ", u, err)

	if err != nil || u.Host == "" {
		return "", errors.New("url is not valid")
	}

	return value, nil
}
