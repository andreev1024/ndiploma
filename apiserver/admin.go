package apiserver

import (
	"database/sql"
	"net/http"
	"strconv"

	"text/template"

	"github.com/gorilla/mux"

	"github.com/andreev1024/ndiploma/storage"
)

func (s *APIServer) adminMainPage(w http.ResponseWriter, req *http.Request) error {
	t, err := template.ParseFiles("templates/admin/main.go.tmpl", "templates/admin/layout.go.tmpl")
	if err != nil {
		return err
	}

	consultRequests, err := s.storage.GetLast100ConsultRequest()
	if err != nil {
		return err
	}

	data := struct {
		Title           string
		ConsultRequests []storage.ConsultRequest
	}{
		Title:           "My page",
		ConsultRequests: consultRequests,
	}

	return t.ExecuteTemplate(w, "base", data)
}

func (s *APIServer) adminItemPage(w http.ResponseWriter, req *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	t, err := template.ParseFiles("templates/admin/item.go.tmpl", "templates/admin/layout.go.tmpl")
	if err != nil {
		return err
	}

	consultRequest, err := s.storage.FindConsultRequest(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found"))
			return nil
		}
		return err
	}

	data := struct {
		Title          string
		ConsultRequest storage.ConsultRequest
	}{
		Title:          "My page",
		ConsultRequest: consultRequest,
	}

	return t.ExecuteTemplate(w, "base", data)
}
