package storage

import (
	"context"
)

// TODO update types (string -> date)
type CreateConsultRequest struct {
	Name          string
	Phone         string
	Role          string
	AvailableTime string
	ConsultDate   string
}

func (s *Storage) CreateConsultRequest(ctx context.Context, i CreateConsultRequest) error {
	_, err := s.conn.Exec("INSERT INTO consult_requests(name, phone, role, available_time, consult_date) VALUES($1, $2, $3, $4, $5)", i.Name, i.Phone, i.Role, i.AvailableTime, i.ConsultDate)
	return err
}
