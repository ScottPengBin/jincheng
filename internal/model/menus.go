package model

type Menus struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parentId" gorm:"column:parentId"`
	Title     string `json:"title" binding:"required"`
	Icon      string `json:"icon"`
	BasePath  string `json:"basePath" gorm:"column:basePath"`
	Path      string `json:"path"`
	Target    string `json:"target"`
	Sort      int    `json:"sort"`
	Type      int    `json:"type" binding:"required"`
	Enabled   bool   `json:"enabled"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Entry     string `json:"entry"`
	CreatedAt MyTime `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt MyTime `json:"updatedAt" gorm:"column:updatedAt"`
}
