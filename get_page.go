package telegraph

import (
	"fmt"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetPage get a Telegraph page. Returns a Page object on success.
func GetPage(path string, returnContent bool) (*Page, error) {
	var args http.Args

	// If true, content field will be returned in Page object.
	args.Add("return_content", strconv.FormatBool(returnContent))

	url := fmt.Sprintf(PathEndpoint, "getPage", path)
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}