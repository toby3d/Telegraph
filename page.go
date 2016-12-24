package telegraph

import (
	"encoding/json"
	"fmt"
	"github.com/kirillDanshin/dlog"
	"github.com/valyala/fasthttp"
	"strconv"
)

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (account *Account) CreatePage(page *Page, returnContent bool) (*Page, error) {
	var args fasthttp.Args

	// Required. Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken)

	// Required. Page title.
	args.Add("title", page.Title)

	// Author name, displayed below the article's title.
	args.Add("author_name", page.AuthorName)

	// Profile link, opened when users click on the author's name below the title. Can be any
	// link, not necessarily to a Telegram profile or channel.
	args.Add("author_url", page.AuthorURL)

	// If true, a content field will be returned in the Page object.
	args.Add("return_content", strconv.FormatBool(returnContent))

	content, err := json.Marshal(page.Content)
	if err != nil {
		return nil, err
	}

	// Required. Content of the page.
	args.Add("content", string(content))

	url := fmt.Sprintf(APIEndpoint, "createPage")
	dlog.Ln(url + "?" + args.String())
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EditPage edit an existing Telegraph page. On success, returns a Page object.
func (account *Account) EditPage(update *Page, returnContent bool) (*Page, error) {
	var args fasthttp.Args

	// Required. Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken)

	// Required. Page title.
	args.Add("title", update.Title)

	// Author name, displayed below the article's title.
	args.Add("author_name", update.AuthorName)

	// Profile link, opened when users click on the author's name below the title. Can be any
	// link, not necessarily to a Telegram profile or channel.
	args.Add("author_url", update.AuthorURL)

	// If true, a content field will be returned in the Page object.
	args.Add("return_content", strconv.FormatBool(returnContent))

	content, err := json.Marshal(update.Content)
	if err != nil {
		return nil, err
	}

	// Required. Content of the page.
	args.Add("content", string(content))

	url := fmt.Sprintf(PathEndpoint, "editPage", update.Path)
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetPage get a Telegraph page. Returns a Page object on success.
func GetPage(path string, returnContent bool) (*Page, error) {
	var args fasthttp.Args

	// If true, content field will be returned in Page object.
	args.Add("return_content", strconv.FormatBool(returnContent))

	url := fmt.Sprintf(PathEndpoint, "getPage", path)
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetPageList get a list of pages belonging to a Telegraph account. Returns a PageList object,
// sorted by most recently created pages first.
func (account *Account) GetPageList(offset int, limit int) (*PageList, error) {
	var args fasthttp.Args

	// Required. Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken)

	// Sequential number of the first page to be returned.
	args.Add("offset", strconv.Itoa(offset))

	// Limits the number of pages to be retrieved.
	args.Add("limit", strconv.Itoa(limit))

	url := fmt.Sprintf(APIEndpoint, "getPageList")
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp PageList
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetViews get the number of views for a Telegraph article. By default, the total number of page
// views will be returned. Returns a PageViews object on success.
func GetViews(path string, year int, month int, day int, hour int) (*PageViews, error) {
	var args fasthttp.Args

	if hour > -1 {
		// If passed, the number of page views for the requested hour will be returned.
		args.Add("hour", strconv.Itoa(hour))

		if day > 0 {
			// Required if hour is passed. If passed, the number of page views for the requested day
			// will be returned.
			args.Add("day", strconv.Itoa(day))

			if month > 0 {
				// Required if day is passed. If passed, the number of page views for the requested month
				// will be returned.
				args.Add("month", strconv.Itoa(month))

				if year >= 2000 && year <= 2100 {
					// Required if month is passed. If passed, the number of page views for the requested year
					// will be returned.
					args.Add("year", strconv.Itoa(year))
				}
			}
		}
	}

	url := fmt.Sprintf(PathEndpoint, "getViews", path)
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp PageViews
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
