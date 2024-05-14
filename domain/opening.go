package domain

import (
	"time"

	"github.com/google/uuid"
)

type Opening struct {
	Id        string    `json:"id" gorm:"primarykey"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    *bool     `json:"remote"`
	Link      string    `json:"link"`
	Salary    string    `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *Opening) Validate() error {
	if err := uuid.Validate(s.Id); string(s.Id) != "" && err != nil {
		return err
	}
	if s.Role == "" {
		return RequiredParamError("role", "string")
	}
	if s.Company == "" {
		return RequiredParamError("company", "string")
	}
	if s.Location == "" {
		return RequiredParamError("location", "string")
	}
	if s.Remote == nil {
		return RequiredParamError("remote", "bool")
	}
	if s.Link == "" {
		return RequiredParamError("link", "string")
	}
	if s.Salary == "" {
		return RequiredParamError("salary", "string")
	}
	return nil
}
