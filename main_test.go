package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHandler (t* testing.T) {
	req, err := http.NewRequest("GET", "", nil) 

	if err != nil {
		t.Fatal(err)
	}


	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler) 

	hf.ServeHTTP(recorder,req) 

	//Check the status code 

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status,http.StatusOK)
	}
	expected := `Hello, Sabin`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body:  got %v want %v", actual, expected)
	}

	//NOTE TO SELF: 
	// http.FileServer() returns a handler that serves HTTP reqs 
	// the contents of the file system at root 

}


//GO USES A CONVENTION TO ASCERTAINS A TEST FILE WHEN IT HAS A PATTERN
//*_test.go 


