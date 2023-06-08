// apiserver/apiserver.go
package apiserver

import (
	"context"
	"errors"
	"html/template"
	"net/http"
	"time"

	"github.com/andreev1024/ndiploma/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var defaultStopTimeout = time.Second * 30

type APIServer struct {
	addr    string
	storage *storage.Storage
}

func NewAPIServer(addr string, storage *storage.Storage) (*APIServer, error) {
	if addr == "" {
		return nil, errors.New("addr cannot be blank")
	}

	return &APIServer{
		addr:    addr,
		storage: storage,
	}, nil
}

// Start starts a server with a stop channel
func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
	}

	go func() {
		logrus.WithField("addr", srv.Addr).Info("starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStopTimeout)
	defer cancel()

	logrus.WithField("timeout", defaultStopTimeout).Info("stopping server")
	return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultRoute)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.Methods("POST").Path("/consult-requests").Handler(Endpoint{s.createConsultRequest})
	router.Methods("GET").Path("/admin").Handler(Endpoint{s.adminMainPage})
	router.Methods("GET").Path("/admin/calendar").Handler(Endpoint{s.calendarPage})
	router.Methods("GET").Path("/admin/consult-request/{id:[0-9]+}").Handler(Endpoint{s.adminItemPage})

	return router
}

func (s *APIServer) defaultRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tmpl, _ := template.ParseFiles("templates/index.go.tmpl")
	tmpl.Execute(w, nil)
}

type Endpoint struct {
	handler EndpointFunc
}

type EndpointFunc func(w http.ResponseWriter, req *http.Request) error

func (e Endpoint) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := e.handler(w, req); err != nil {
		logrus.WithError(err).Error("could not process request")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
}
