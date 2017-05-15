package main

import (
    "fmt"
    "github.com/kylelemons/go-gypsy/yaml"
    "net/http"
)

type SrvConfig struct {
    DocumentRoot    string
    Listen          string
    TLS             bool
    TLSListen       string
    TLSCert         string
    TLSKey          string
}

func (c *SrvConfig) read(fName string) (err error) {
    // read in a yaml formated configuration

    config, err := yaml.ReadFile(fName)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

    c.DocumentRoot, _ = config.Get("DocumentRoot") 
    c.Listen, _ = config.Get("Listen")
    c.TLS, _ = config.GetBool("TLS")
    c.TLSListen, _ = config.Get("TLSListen")
    c.TLSCert, _ = config.Get("TLSCert")
    c.TLSKey, _ = config.Get("TLSKey")

    return
}

func startMainServer(c SrvConfig) (err error) {
    // The main server 

    fmt.Println("Start Server")

    h_mux := http.NewServeMux()
    h_files := http.FileServer(http.Dir(c.DocumentRoot))
    h_mux.Handle("/", http.StripPrefix("/", h_files))

    h_server := http.Server{
        Addr:       c.Listen,
        Handler:    h_mux,
    }
    h_server.ListenAndServe()

    err = nil

    return

}

func startTLSServer(c SrvConfig) (err error) {
    // A TLS server 
    fmt.Println("Start TLS Server")

    t_mux := http.NewServeMux()
    t_files := http.FileServer(http.Dir(c.DocumentRoot))
    t_mux.Handle("/", http.StripPrefix("/", t_files))

    t_server := http.Server{
        Addr:   c.TLSListen,
        Handler: t_mux,
    }
    t_server.ListenAndServeTLS(c.TLSCert, c.TLSKey)

    err = nil

    return

}


func main() {

    // get a configuration 
    var curConf SrvConfig

    err := curConf.read("mvp-httpd.yaml")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

    fmt.Println("Configuration Read")

    fmt.Println(curConf)

    if curConf.TLS {
        
        go startTLSServer(curConf)

    }

    startMainServer(curConf)

}

