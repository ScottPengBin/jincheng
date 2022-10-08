package model

import (
	"time"
)

type Menus struct {
	Id        int       `json:"id"`
	ParentId  int       `json:"parentId"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	BasePath  string    `json:"basePath"`
	Path      string    `json:"path"`
	Target    string    `json:"target"`
	Sort      int       `json:"sort"`
	Type      int       `json:"type"`
	Enabled   int       `json:"enabled"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Entry     string    `json:"entry"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
