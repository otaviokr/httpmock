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
	ExternalAddressesTestable []Testable
)

func init() {
	SingleServingTestable = []Testable{
		{
			U: "/sendTest",
			BT: "application/json",
			B: "not importante",
			I: DummyResponse{
				Body: "{Result:True}",
				Code: 200,
				HeaderValues: map[string][]string{"Content-Type": {"application/json"}}},
			E: "{Result:True}\n"}}

	MultipleServingTestable = []MultiTestable{
		{
			U: []string{"/test.html", "/test.json", "/other.php", "/generic"},
			BT: "not important",
			B: "not important",
			I: map[string]DummyResponse{
				"/test.html": {Body: "Result from HTML test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200},
				"/test.json": {Body: "Result from JSON test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200},
				"/other.php": {Body: "Result from PHP test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200},
				"/generic": {Body: "Result from generic test",
					HeaderValues: map[string][]string{"Content-Type": {"text/html"}}, Code: 200}},
			E: map[string]string{
				"/test.html": "Result from HTML test\n",
				"/test.json": "Result from JSON test\n",
				"/other.php": "Result from PHP test\n",
				"/generic": "Result from generic test\n"}}}

	CheckHeaderValuesTestable = []Testable {
		{
			U: "/sendTestHeader",
			BT: "not important",
			B: "not important",
			I: DummyResponse{
				Body: "{Result:True}",
				Code: 200,
				HeaderValues: map[string][]string{"Content-Type": {"application/json"}, "Test-Key": {"TestValue1", "TestValue2"}}},
			E: map[string][]string{"Content-Type": {"application/json"}, "Test-Key": {"TestValue1", "TestValue2"}}}}

	ExternalAddressesTestable = []Testable {
		{
			U: "http://www.google.com",
			BT: "not important",
			B: "WWWRRROOONNNGGGGG!!!",
			I: DummyResponse{
				Body: "not important",
				Code: 500,
				HeaderValues: map[string][]string{}},
			E: map[string][]string{}}}
}

func TestSingleServing(T *testing.T) {
	for _, testCase := range SingleServingTestable {
		server, client := ServeGeneric(testCase.I)

        response, err := client.Get(server.URL+testCase.U)
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

        response.Body.Close()
        server.Close()
	}
}

func TestMultipleService(T *testing.T) {
	for _, testCase := range MultipleServingTestable {
        server, client := ServeMulti(testCase.I)

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
        server.Close()
	}
}

func TestSingleServingHeaders(T *testing.T) {
	for _, testCase := range CheckHeaderValuesTestable {
        server, client := ServeGeneric(testCase.I)

		response, err := client.Get(server.URL+testCase.U)
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
        response.Body.Close()
        server.Close()
	}
}

func TestExternalAddressesGeneric(T *testing.T) {
    for _, testCase := range ExternalAddressesTestable {
        server, client := ServeGeneric(testCase.I)

        response, err := client.Get(testCase.U)
        if err != nil {
            T.Fatalf(fmt.Sprintf("Could not get response from request %s:\n%+v\n", testCase.U, err))
        }

        data, err := ioutil.ReadAll(response.Body)
        if err != nil {
            T.Fatalf(fmt.Sprintf("Error parsing the response body from %s:\n%+v\n", testCase.U, err))
        }

        if testCase.B == string(data) {
            T.Fatalf(fmt.Sprintf("External connection did not worked! from %s:\n%+v\n", testCase.U, err))
        }

        response.Body.Close()
        server.Close()
    }
}
