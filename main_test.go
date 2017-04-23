package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestReadConfig(t *testing.T) {
    var curConf SrvConfig
    err := curConf.read("conf-demo.yaml")
    if err != nil {
        t.Error(err)
    }

    if curConf.DocumentRoot != "/var/www/html" {
        t.Error(`Wrong document root, was expecting "/var/www/html" but got`, curConf.DocumentRoot)
    }

    if curConf.Listen != "127.0.0.1:8080" {
        t.Error(`Wrong listening address, was expecting "127.0.0.1:8080" but got`, curConf.Listen)
    }

    if curConf.TLS != false {
        t.Error(`Wrong TLS flag, was expecting false but got`, curConf.TLS)
    }    

    if curConf.TLSListen != "127.0.0.1:8443" {
        t.Error(`Wrong TLS listening address, was expecting "127.0.0.1:8443" but got`, curConf.TLSListen)
    }    

    if curConf.TLSCert != "/etc/ssl/certs/cert.pem" {
        t.Error(`Wrong TLS cert path, was expecting "/etc/ssl/certs/cert.pem" but got`, curConf.TLSCert)
    }    

    if curConf.TLSKey != "/etc/ssl/private/key.pem" {
        t.Error(`Wrong TLS key path, was expecting "/etc/ssl/private/key.pem" but got`, curConf.TLSKey)
    }

}

func TestBasicHTTP(t *testing.T) {
    // ensure there is an index.html in /var/www/html

    var curConf SrvConfig
    err := curConf.read("conf-demo.yaml")
    if err != nil {
        t.Error(err)
    }

    mux := http.NewServeMux()
    files := http.FileServer(http.Dir(curConf.DocumentRoot))
    mux.Handle("/", http.StripPrefix("/", files))

    writer := httptest.NewRecorder()
    request := httptest.NewRequest("GET", "/", nil)
    mux.ServeHTTP(writer, request)

    if writer.Code != 200 {
        t.Errorf("Server did not return index.html, response code is %v", writer.Code)
    }

    if testing.Verbose() {

        // show the header and the body if we are verbose
        resp := writer.Result()
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(resp.Header.Get("Content-Type"))
        fmt.Println(string(body))

    }

}


