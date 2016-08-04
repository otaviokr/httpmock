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
        URL1: {
            Body: fmt.Sprintf("Here's your answer for %s!", URL1),
            HeaderValues: map[string][]string{"Content-Type": {"meh"}},
            Code: 200},
        URL2: {
            Body: fmt.Sprintf("Here's your answer for %s!", URL2),
            HeaderValues: map[string][]string{"Content-Type": {"meh"}},
            Code: 200},
        URL3: {
            Body: fmt.Sprintf("Here's your answer for %s!", URL3),
            HeaderValues: map[string][]string{"Content-Type": {"meh"}},
            Code: 200}}

    Server, Client := httpmock.ServeMulti(Dummies)
    defer Server.Close()

    for _, URL := range URLList {
        // Always remember to send the request to the server you created!
        Response, Err := Client.Get(Server.URL + URL)

        if Err != nil {
            fmt.Printf("Error on server processing request: %s", Err.Error())
        }

        DataInBytes, Err := ioutil.ReadAll(Response.Body)
        if Err != nil {
            fmt.Printf("Error processing response body: %s", Err.Error())
        }

        fmt.Printf("%s ** result: (%d) ==> %s", URL, Response.StatusCode, string(DataInBytes))
        Response.Body.Close()
    }
}
