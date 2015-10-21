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
	CheckHeaderValuesTestable []Testable
)

func init() {
	SingleServingTestable = []Testable{
		Testable{
			U: "/sendTest",
			BT: "application/json",
			B: "not importante",
			I: DummyResponse{
				Body: "{Result:True}",
				Code: 200,
				HeaderValues: map[string][]string{"Content-Type": {"application/json"}}},
			E: "{Result:True}\n"}}

	MultipleServingTestable = []MultiTestable{
		MultiTestable{
			U: []string{"/test.html", "/test.json", "/other.php", "/generic"},
			BT: "not important",
			B: "not important",
			I: map[string]DummyResponse{
				"/test.html": DummyResponse{Body: "Result from HTML test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200},
				"/test.json": DummyResponse{Body: "Result from JSON test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200},
				"/other.php": DummyResponse{Body: "Result from PHP test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200},
				"/generic": DummyResponse{Body: "Result from generic test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200}},
			E: map[string]string{
				"/test.html": "Result from HTML test\n",
				"/test.json": "Result from JSON test\n",
				"/other.php": "Result from PHP test\n",
				"/generic": "Result from generic test\n"}}}

	CheckHeaderValuesTestable = []Testable {
		Testable{
			U: "/sendTestHeader",
			BT: "not important",
			B: "not important",
			I: DummyResponse{
				Body: "{Result:True}",
				Code: 200,
				HeaderValues: map[string][]string{"Content-Type": {"application/json"}, "Test-Key": {"TestValue1", "TestValue2"}}},
			E: map[string][]string{"Content-Type": {"application/json"}, "Test-Key": {"TestValue1", "TestValue2"}}}}
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

func TestSingleServingHeaders(T *testing.T) {
	for _, testCase := range CheckHeaderValuesTestable {
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

		header := response.Header
		expected := testCase.E.(map[string][]string)
		for key, values := range expected {
			for i, value := range values {
				if header[key][i] != value {
					T.Fatalf(fmt.Sprintf("Mismatch - %s!\nExpected: %+v\nFound: %+v\n", key, expected, header))
				}
			}
		}
	}
}
