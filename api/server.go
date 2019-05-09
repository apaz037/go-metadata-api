package api

import (
	"context"
	"crypto/tls"
	"github.com/apaz037/go-metadata-api/api/utils"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	tlsKey  = "key.pem"
	tlsCert = "cert.pem"
	tlsCfg  = &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	utils.GenerateTLSKeys()

	api := New() // create an instance of our API

	var addr string

	port := viper.GetString("port")

	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:           addr,
		Handler:        api,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      tlsCfg,
		TLSNextProto:   make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	return &Server{&srv}
}

func (srv *Server) Start() {
	go func() {
		if err := srv.ListenAndServeTLS(tlsCert, tlsKey); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	sig := <-quit
	log.Println("Shutting down server: ", sig)

	// shut down server
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server shut down")
}
