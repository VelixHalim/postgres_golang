package model

import (
	"time"
)

type Profile struct {
	ID       string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
	UpdateAt time.Time
}

type Profiles []Profile

func NewProfile() *Profile {
	return &Profile{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}
