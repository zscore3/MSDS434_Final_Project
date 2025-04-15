package main

import( 
  "fmt"
  "net/http"
  "time"
  )

func main(){
  http.HandleFunc("/", index)
  fmt.Println("Now Serving localhost:8000")
  http.ListenAndServe(":8000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, this is a web page! %s", time.Now())
}

