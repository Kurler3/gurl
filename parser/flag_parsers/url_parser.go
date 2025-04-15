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

	u, err := url.ParseRequestURI(urlToCheck)

	if err != nil {
		return "", fmt.Errorf("eror while validating url: %v", err)
	}

	if u.Scheme != utils.HTTP && u.Scheme != utils.HTTPS {
		return "", errors.New("url has invalid protocol")
	}

	host := u.Host

	if host == "" || !strings.Contains(host, ".") {
		return "", errors.New("invalid host name")
	}

	return value, nil
}
