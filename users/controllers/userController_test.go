// userController_test.go
package controllers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/mmorejon/cinema/users/common"
)

func TestGetUsers(t *testing.T) {
    // setup the app
    common.StartUp()

    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetUsers)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"data":[{"id":"58b1ec0ccd02960006e0fcde","name":"Jane","lastname":"Doe"},{"id":"58b1ec79cd02960006e0fcdf","name":"John","lastname":"Doe"},{"id":"58b1ebf7cd02960006e0fcdd","name":"Billy","lastname":"Bob"}]}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
