package apiserver

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"

	"github.com/andreev1024/ndiploma/storage"
)

const SessionName = "session"
const SessionKey = "authenticated"
const SessionMaxAge = 86400 // 1 day

func (s *APIServer) adminMainPage(w http.ResponseWriter, req *http.Request) error {
	//TODO move to middleware
	session, _ := s.sessions.Get(req, SessionName)
	_, sessionExist := session.Values[SessionKey]
	if !sessionExist {
		http.Redirect(w, req, AdminLoginPage, http.StatusSeeOther)
		return nil
	}

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
		Title:           "Admin Main Page",
		ConsultRequests: consultRequests,
	}

	return t.ExecuteTemplate(w, "base", data)
}

func (s *APIServer) calendarPage(w http.ResponseWriter, req *http.Request) error {
	session, _ := s.sessions.Get(req, SessionName)
	_, sessionExist := session.Values[SessionKey]
	if !sessionExist {
		http.Redirect(w, req, AdminLoginPage, http.StatusSeeOther)
		return nil
	}

	t, err := template.ParseFiles("templates/admin/calendar.go.tmpl", "templates/admin/layout.go.tmpl")
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, "base", nil)
}

func (s *APIServer) adminItemPage(w http.ResponseWriter, req *http.Request) error {
	session, _ := s.sessions.Get(req, SessionName)
	_, sessionExist := session.Values[SessionKey]
	if !sessionExist {
		http.Redirect(w, req, AdminLoginPage, http.StatusSeeOther)
		return nil
	}

	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	t, err := template.ParseFiles("templates/admin/consult-request.go.tmpl", "templates/admin/layout.go.tmpl")
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
		Title:          consultRequest.Name,
		ConsultRequest: consultRequest,
	}

	return t.ExecuteTemplate(w, "base", data)
}

func (s *APIServer) loginPage(w http.ResponseWriter, req *http.Request) error {
	t, err := template.ParseFiles("templates/admin/login.go.tmpl")
	if err != nil {
		return err
	}

	_, error := req.URL.Query()["error"]

	errorMsg := ""
	if error {
		errorMsg = "Invalid Login or Password"
	}

	data := struct {
		ErrorMsg string
	}{
		ErrorMsg: errorMsg,
	}

	return t.Execute(w, data)
}

func (s *APIServer) login(w http.ResponseWriter, req *http.Request) error {
	login := req.FormValue("login")
	password := req.FormValue("password")
	hashFromLoginAndPassword := md5.Sum([]byte(fmt.Sprintf("%s%s", login, password)))

	if fmt.Sprintf("%x", hashFromLoginAndPassword) == "f6fdffe48c908deb0f4c3bd36c032e72" {
		session, _ := s.sessions.Get(req, SessionName)
		session.Options.MaxAge = SessionMaxAge
		session.Values[SessionKey] = true
		session.Save(req, w)
		http.Redirect(w, req, AdminMainPage, http.StatusSeeOther)
	} else {
		http.Redirect(w, req, fmt.Sprintf("%s?error=1", AdminLoginPage), http.StatusSeeOther)
	}

	return nil
}

func (s *APIServer) logout(w http.ResponseWriter, req *http.Request) error {
	session, _ := s.sessions.Get(req, SessionName)
	session.Options.MaxAge = -1
	session.Save(req, w)

	http.Redirect(w, req, AdminLoginPage, http.StatusSeeOther)

	return nil
}
