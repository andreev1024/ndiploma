package apiserver

import (
	"net/http"

	"github.com/andreev1024/ndiploma/storage"
)

// TODO validation
func (s *APIServer) createConsultRequest(w http.ResponseWriter, req *http.Request) error {
	err := s.storage.CreateConsultRequest(req.Context(), storage.CreateConsultRequest{
		Name:          req.PostFormValue("name"),
		Phone:         req.PostFormValue("phone"),
		Role:          req.PostFormValue("role"),
		AvailableTime: req.PostFormValue("available-time"),
		ConsultDate:   req.PostFormValue("consult-date"),
	})

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}
