package server

import (
    "net/http"
    "net/http/httptest"
	"testing"
    "bytes"

    "gitlab.com/andersph/docker-master-class/api/internal"
)

func TestAlertAPI(t *testing.T) {
	
	// load configuration
    c := internal.LoadConfig()

    var jsonStr = []byte(`{"name":"Max Mustermann", "address":"01234 Musterhausen, Musterstraße 1a", "email":"max@mustermann.org", "birth":"1989-11-11", "department":"Musterdepartment", "job_title":"Musterjob"}`)

	req, err := http.NewRequest("POST", c.Endpoints.CreateEmployee, bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
	}

    // ResponseRecorder (which satisfies http.ResponseWriter) to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(newEmployee)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

    // Check the response body is what we expect.
    if internal.GetEnv("IS_STDOUT", c.Flags.Is_StdOut) != "true" {
        expected := `Employee entry was successfully created on database: {Name:Max Mustermann Address:01234 Musterhausen, Musterstraße 1a Mail:max@mustermann.org DateOfBirth:1989-11-11 Department:Musterdepartment JobTitle:Musterjob}`
        if rr.Body.String() != expected {
            t.Errorf("handler returned unexpected body: got %v want %v",
                rr.Body.String(), expected)
        }
    } else {
        t.Errorf("Test not implemented for std_out option")
    }
}
