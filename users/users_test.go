// userController_test.go
// See https://elithrar.github.io/article/testing-http-handlers-go/
// https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "os"
    "github.com/gorilla/mux"
    "github.com/mmorejon/cinema/common"
    "github.com/mmorejon/cinema/users/routers"
)

var router *mux.Router

func TestMain(m *testing.M) {
    // setup the app
    common.StartUp()
    router = routers.InitRoutes()
    retCode := m.Run()
    os.Exit(retCode)
}

func TestGetUsers(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    //handler := http.HandlerFunc(controllers.GetUsers)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    router.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect.
    actual := rr.Body.String()
    expected := `{"data":[{"id":"58b1ebf7cd02960006e0fcdd","name":"Billy","lastname":"Bob"},{"id":"58b1ec0ccd02960006e0fcde","name":"Jane","lastname":"Doe"},{"id":"58b1ec79cd02960006e0fcdf","name":"John","lastname":"Doe"}]}`

    if actual != expected {
      t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }
}

func TestGetUserById_Exists(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/users/58b1ec0ccd02960006e0fcde", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    //handler := http.HandlerFunc(controllers.GetUserById)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    router.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect.
    actual := rr.Body.String()
    expected := `{"data":{"id":"58b1ec0ccd02960006e0fcde","name":"Jane","lastname":"Doe"}}`

    if actual != expected {
      t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }
}

func TestGetUserById_404(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/users/11b1ec0ccd02960006e0fcde", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    router.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
    }

    // Check the response body is empty
    actual := rr.Body.String()
    if actual != "" {
      t.Errorf("handler returned unexpected body: got [%v] want no body", actual)
    }
}
