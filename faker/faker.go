package faker

import (
	"io/ioutil"
	"net/http"

	"github.com/weslenng/dorothy/config"
)

func GetProductName() (string, error) {
	r, err := http.Get(config.Default.FakerAddr)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
