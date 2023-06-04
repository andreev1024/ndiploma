package storage

import (
	"context"
	// "database/sql"
)

// TODO update types (string -> date)
type CreateConsultRequest struct {
	Name          string
	Phone         string
	Role          string
	AvailableTime string
	ConsultDate   string
}

type ConsultRequest struct {
	Id            int
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

func (s *Storage) GetLast100ConsultRequest() ([]ConsultRequest, error) {
	rows, err := s.conn.Query("SELECT id, name, phone, role, available_time, consult_date FROM consult_requests ORDER BY created_at DESC LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consultRequests []ConsultRequest

	for rows.Next() {
		var consulRequest ConsultRequest
		if err := rows.Scan(&consulRequest.Id, &consulRequest.Name, &consulRequest.Phone, &consulRequest.Role, &consulRequest.AvailableTime, &consulRequest.ConsultDate); err != nil {
			return consultRequests, err
		}
		consultRequests = append(consultRequests, consulRequest)
	}
	if err = rows.Err(); err != nil {
		return consultRequests, err
	}
	return consultRequests, nil
}

func (s *Storage) FindConsultRequest(id int) (ConsultRequest, error) {
	var consulRequest ConsultRequest
	row := s.conn.QueryRow("SELECT id, name, phone, role, available_time, consult_date FROM consult_requests WHERE id = $1", id)
	err := row.Scan(&consulRequest.Id, &consulRequest.Name, &consulRequest.Phone, &consulRequest.Role, &consulRequest.AvailableTime, &consulRequest.ConsultDate)
	return consulRequest, err
}
