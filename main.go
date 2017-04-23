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

func main() {

    // get a configuration 
    var curConf SrvConfig

    err := curConf.read("conf.yaml")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

    fmt.Println("Configuration Read")

    fmt.Println(curConf)

    mux := http.NewServeMux()
    files := http.FileServer(http.Dir(curConf.DocumentRoot))
    mux.Handle("/", http.StripPrefix("/", files))

    if curConf.TLS {
        fmt.Println("Start TLS Server")
        t_server := http.Server{
            Addr:   curConf.TLSListen,
            Handler: mux,
        }
        t_server.ListenAndServeTLS(curConf.TLSCert, curConf.TLSKey)

    } else {    
        fmt.Println("Start Server")
        h_server := http.Server{
            Addr:       curConf.Listen,
            Handler:    mux,
        }
        h_server.ListenAndServe()
    }

}

