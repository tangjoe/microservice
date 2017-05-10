package main

import (
   "os"
   "encoding/json"
   "net/http"
   "fmt"
)

type Hello struct {
   Message string
}

func main() {
   port := os.Getenv("PORT")
   if len(port) == 0 {
      port = "8080"
   }

   fmt.Println("Listening on " + port + " ...")
   http.HandleFunc("/hello", hello)
   http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
   m := Hello {"Welcome to Microservice with Go --- Joe Tang."}
   b, err := json.Marshal(m)

   if err != nil {
      panic(err)
   }

   w.Header().Add("Content-Type", "application/json;charset=utf-8")
   w.Write(b)
}
