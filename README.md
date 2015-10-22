# HTTPMock 
## A Go lib to mock HTTP responses.

[![Coverage Status](https://coveralls.io/repos/otaviokr/httpmock/badge.svg?branch=master&service=github)](https://coveralls.io/github/otaviokr/httpmock?branch=master)
[![Build Status](https://travis-ci.org/otaviokr/httpmock.svg)](https://travis-ci.org/otaviokr/httpmock)

This is an auxiliary lib to help you simulate HTTP requests or redirect your requests - for example, if you're running
tests and don't want to hit the actual destination.

Each response is represented by a `DummyResponse` instance, that should be defined by the tester:

```go
type DummyResponse struct {
	Code int
	HeaderValues map[string][]string
	Body string
}
```

where:
- **Code**: the Response Code; usually 200 for a successful request;
- **HeaderValues**: key-value pairs stored in the response header. Previous ContentType attribute should now be stored here, under *Content-Type* key;
- **Body**: the body response content.

If you only want a generic one-for-all response, use `ServeGeneric`. It will ignore the URL requested and provide the 
same response every time - obviously, it won't recognize an invalid address. If you want the server it to respond to an 
specific URL or set of URLs, the best option is `ServeMulti`. The keys in the parameter map are the URLs that should be 
answered with the dummy response defined as its value. I think the examples below explain better.

## Examples

The examples below are complete, meaning you can execute them individually.

This example shows how `ServeGeneric` will answer always the same dummy response to any request sent.

```go
package main
import (
	"io/ioutil"
	"fmt"
	"github.com/otaviokr/httpmock"
)

func main() {
	URLList := []string{"/", "/anythingGoes", "/another/example/to/test.html"}
	Dummy := httpmock.DummyResponse{
		Body: "Here's your answer!",
		HeaderValues: map[string][]string{"Arbitrary-Key": {"does not matter"}},
		Code: 200}

	Server, Client := httpmock.ServeGeneric(Dummy)
	defer Server.Close()

	for _, URL := range URLList {
		// Always remember to send the request to the server you created!
		Response, Err := Client.Get(Server.URL + URL)
		defer Response.Body.Close()
		if Err != nil {
			fmt.Printf("Error on server processing request: %s\n", Err.Error())
		}

		DataInBytes, Err := ioutil.ReadAll(Response.Body)
		if Err != nil {
			fmt.Printf("Error processing response body: %s", Err.Error())
		}

		fmt.Printf("%s result: (%d) %s => %s\n",
			URL, Response.StatusCode, Response.Header["Arbitrary-Key"][0], string(DataInBytes))
	}
}
```

This example shows how `ServeMulti` will answer each specified URL with the designated response.

```go
package main
import (
	"fmt"
	"io/ioutil"
	"github.com/otaviokr/httpmock"
)

func main() {
	URL1 := "/"
	URL2 := "/fyeo"
	URL3 := "/another/example/to/test.html"
	URLList := []string{URL1, URL2, URL3}
	Dummies := map[string]httpmock.DummyResponse {
		URL1: httpmock.DummyResponse{
			Body: fmt.Sprintf("Here's your answer for %s!", URL1),
			HeaderValues: map[string][]string{"Content-Type": {"meh"}},
			Code: 200},
		URL2: httpmock.DummyResponse{
			Body: fmt.Sprintf("Here's your answer for %s!", URL2),
			HeaderValues: map[string][]string{"Content-Type": {"meh"}},
			Code: 200},
		URL3: httpmock.DummyResponse{
			Body: fmt.Sprintf("Here's your answer for %s!", URL3),
			HeaderValues: map[string][]string{"Content-Type": {"meh"}},
			Code: 200}}

	Server, Client := httpmock.ServeMulti(Dummies)
	defer Server.Close()

	for _, URL := range URLList {
		// Always remember to send the request to the server you created!
		Response, Err := Client.Get(Server.URL + URL)
		defer Response.Body.Close()
		if Err != nil {
			fmt.Printf("Error on server processing request: %s", Err.Error())
		}

		DataInBytes, Err := ioutil.ReadAll(Response.Body)
		if Err != nil {
			fmt.Printf("Error processing response body: %s", Err.Error())
		}

		fmt.Printf("%s result: (%d)%s\n", URL, Response.StatusCode, string(DataInBytes))
	}
}
```

## Future Features

This is my wishlist of things to include in this lib. I'm not making any promises and most of them are just random 
thoughts that I still need to polish before having something interesting.

- [ ] **Add form/URL processing** - process values passed as POST or GET to influence the output response;
- [ ] **Work with templates** - add templates to answers;
