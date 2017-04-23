package main

import (
    "fmt"
    "github.com/kylelemons/go-gypsy/yaml"
    "net/http"
)

func main() {

    // read in a yaml formated configuration
    config, err := yaml.ReadFile("conf.yaml")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

    fPath, _ := config.Get("DocumentRoot") 
    lServ, _ := config.Get("Listen")
    sServ, _ := config.Get("SSLListen")
    cert, _ := config.Get("SSLCert")
    key, _ := config.Get("SSLKey")

    fmt.Println("Configuration Read")

    fmt.Println(fPath, lServ, sServ, cert, key)

    mux := http.NewServeMux()
    files := http.FileServer(http.Dir(fPath))
    mux.Handle("/", http.StripPrefix("/", files))


    h_server := http.Server{
        Addr:       lServ,
        Handler:    mux,
    }
    h_server.ListenAndServe()

    t_server := http.Server{
        Addr:   sServ,
        Handler: mux,
    }
    t_server.ListenAndServeTLS(cert, key)

}

