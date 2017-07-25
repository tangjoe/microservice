package main

import (
   "net/http"
   "net/http/httptest"
   "testing"
)

func TestHandler(t *testing.T) {
   req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
   if err != nil {
      t.Fatal(err)
   }

   rr := httptest.NewRecorder()
   handler := http.HandlerFunc(hello)

   handler.ServeHTTP(rr, req)

   if status := rr.Code; status != http.StatusOK {
      t.Errorf("handler returned wrong status code: go %v want %v",
         status, http.StatusOK)
   }

   expected := `{"Message":"Hello Go."}`
   if rr.Body.String() != expected {
      t.Errorf("handler returned unexpected body: got %v want %v",
         rr.Body.String(), expected)
   }
}
