package examples

import (
    "fmt"
    "io/ioutil"
    "github.com/otaviokr/httpmock"
)

func ServeGeneric() {
    // These are the addresses we will test. They'll be appended to the server URL when the client makes the request.
    URLList := []string{"/", "/anythingGoes", "/another/example/to/test.html"}

    // This is the "page" served by the server.
    dummyResponse := httpmock.DummyResponse{
        Body: "Here's your answer!",
        HeaderValues: map[string][]string{"Arbitrary-Key": {"they all answer the same!"}},
        Code: 200}

    // Creating the server and client instances. ServeGeneric expects a generic response that will be served to ANY
    // request made. Don't forget to close the server when you're done!
    Server, Client := httpmock.ServeGeneric(dummyResponse)
    defer Server.Close()

    for _, URL := range URLList {
        // All requests must be made to our server instance!
        Response, Err := Client.Get(Server.URL + URL)

        if Err != nil {
            fmt.Printf("Error on server processing request: %s\n", Err.Error())
        }

        DataInBytes, Err := ioutil.ReadAll(Response.Body)
        if Err != nil {
            fmt.Printf("Error processing response body: %s", Err.Error())
        }

        fmt.Printf("%s ** result: (%d) %s => %s",
            URL, Response.StatusCode, Response.Header["Arbitrary-Key"][0], string(DataInBytes))

        Response.Body.Close()
    }
}
