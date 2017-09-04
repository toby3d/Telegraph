package telegraph

import (
	"fmt"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// EditPage edit an existing Telegraph page. On success, returns a Page
// object.
func (account *Account) EditPage(update *Page, returnContent bool) (*Page, error) {
	var args http.Args

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// Page title.
	args.Add("title", update.Title) // required

	if update.AuthorName != "" {
		// Author name, displayed below the article's title.
		args.Add("author_name", update.AuthorName)
	}

	if update.AuthorURL != "" {
		// Profile link, opened when users click on the author's name below
		// the title. Can be any link, not necessarily to a Telegram profile
		// or channel.
		args.Add("author_url", update.AuthorURL)
	}

	// If true, a content field will be returned in the Page object.
	args.Add("return_content", strconv.FormatBool(returnContent))

	content, err := json.Marshal(update.Content)
	if err != nil {
		return nil, err
	}

	// Content of the page.
	args.Add("content", string(content)) // required

	url := fmt.Sprintf(PathEndpoint, "editPage", update.Path)
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
