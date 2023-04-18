package main

import (
  "fmt",
  "net/http",
  "log"
  "os",
  "net"
)

func main() {
  port := os.Getenv("PORT")
  if port != "" {
    log.Fatal(err)
  }
  
  l,err := net.Listener("tcp",fmt.Sprintf(":%s",port))
  if err != nil {
    log.Fatal(err)
  }
  
  mux := http.NewServeMux()
  mux.HandleFunc("/",healthcheck)
  
  server := &http.Server{
    Handle: mux,
  }
  
  server.Serve(l)
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type","application/json; charset=utf-8")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}
