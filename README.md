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

| Attribute    | Description |
|--------------|-------------|
| Code         | It's the Response Code; usually 200 for a successful request |
| HeaderValues | Key-value pairs stored in the response header. Previous ContentType attribute should now be stored here, under *Content-Type* key |
| Body         | The body response content |

If you only want a generic one-for-all response, use `ServeGeneric`. It will ignore the URL requested and provide the 
same response every time - obviously, it won't recognize an invalid address. 

If you want the server it to respond to an 
specific URL or set of URLs, the best option is `ServeMulti`. The keys in the parameter map are the URLs that should be 
answered with the dummy response defined as its value. I think the examples below explain better.

You can also make requests to real-world sites (like Reddit, Youtube, your favorite blog etc.) and the Mock client 
won't redirect the request to the mock server; instead, it will work as a normal HTTP client.

## Examples

You can find some usage examples on folder `/examples`. They are complete programs, so you can run each script to see 
them in action. More details about the each script can be found in `/examples/README.md`.

## Future Features

This is my wishlist of things to include in this lib. I'm not making any promises and most of them are just random 
thoughts that I still need to polish before having something interesting.

- [ ] **Add form/URL processing** - process values passed as POST or GET to influence the output response;
- [ ] **Work with templates** - add templates to answers;
