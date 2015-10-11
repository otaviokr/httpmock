package httpmock
import (
	"testing"
	"io/ioutil"
	"fmt"
)

type Testable struct {
	U string
	BT string
	B string
	I DummyResponse
	E interface{}
}

type MultiTestable struct {
	U []string
	BT string
	B string
	I map[string]DummyResponse
	E interface{}
}

var (
	SingleServingTestable []Testable
	MultipleServingTestable []MultiTestable
)

func init() {
	SingleServingTestable = []Testable{
		Testable{
			U: "/sendTest",
			BT: "application/json",
			B: "not importante",
			I: DummyResponse{Body: "{Result:True}", ContentType: "application/json", Code: 200},
			E: "{Result:True}\n"}}

	MultipleServingTestable = []MultiTestable{
		MultiTestable{
			U: []string{"/test.html", "/test.json", "/other.php", "/generic"},
			BT: "not important",
			B: "not important",
			I: map[string]DummyResponse{
				"/test.html": DummyResponse{Body: "Result from HTML test", ContentType: "text/html", Code: 200},
				"/test.json": DummyResponse{Body: "Result from JSON test", ContentType: "text/html", Code: 200},
				"/other.php": DummyResponse{Body: "Result from PHP test", ContentType: "text/html", Code: 200},
				"/generic": DummyResponse{Body: "Result from generic test", ContentType: "text/html", Code: 200}},
			E: map[string]string{
				"/test.html": "Result from HTML test\n",
				"/test.json": "Result from JSON test\n",
				"/other.php": "Result from PHP test\n",
				"/generic": "Result from generic test\n"}}}
}

func TestSingleServing(T *testing.T) {
	for _, testCase := range SingleServingTestable {
		server, client := ServeGeneric(testCase.I)
		defer server.Close()

		response, err := client.Get(server.URL+testCase.U)
		defer response.Body.Close()
		if err != nil {
			T.Fatal(fmt.Sprintf("Error on server processing request: %s", err.Error()))
		}

		expectedCode := testCase.I.Code
		actualCode := response.StatusCode
		if actualCode != expectedCode {
			T.Fatal(fmt.Sprintf("Unexpected status code - expected %d, but received %d", expectedCode, actualCode))
		}

		dataInBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			T.Fatal(fmt.Sprintf("Error processing response body: %s", err.Error()))
		}

		data := string(dataInBytes)
		expected := testCase.E.(string)
		if expected != data {
			T.Fatal(fmt.Sprintf("Mismatch response:\nExpected: %s\nFound: %s", expected, data))
		}
	}
}

func TestMultipleService(T *testing.T) {
	for _, testCase := range MultipleServingTestable {
		server, client := ServeMulti(testCase.I)
		defer server.Close()

		for _, url := range testCase.U {
			response, err := client.Get(server.URL + url)
			if err != nil {
				T.Fatal(fmt.Sprintf("Error on server processing request: %s", err.Error()))
			}

			dataInBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				T.Fatal(fmt.Sprintf("Error processing response body: %s", err.Error()))
			}

			data := string(dataInBytes)
			expected := testCase.E.(map[string]string)[url]
			if expected != data {
				T.Fatal(fmt.Sprintf("Mismatch response:\nExpected: %s\nFound: %s", expected, data))
			}
		}
	}
}
