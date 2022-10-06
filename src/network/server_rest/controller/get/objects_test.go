package get

import (
	"D7024E/node/id"
	"D7024E/node/stored"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Test
func TestObjects(t *testing.T) {
	// Create value
	value := stored.Value{Data: "this is data"}
	value.ID = *id.NewKademliaID(value.Data)
	jsonValue, err := json.Marshal(value)
	if err != nil {
		t.FailNow()
	}

	// Store value
	stored.GetInstance().Store([]stored.Value{value})

	// Create request
	route := "/objects/" + value.ID.String()
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.FailNow()
	}

	// Set request parameters
	vars := map[string]string{
		"hash": value.ID.String(),
	}
	req = mux.SetURLVars(req, vars)

	// Test the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Objects)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := string(jsonValue)
	body := strings.ReplaceAll(rr.Body.String(), "\n", "")
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expected)
	}
}
