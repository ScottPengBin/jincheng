package meber

type MemRequest struct {
	PageNum   int
	PageSize  int
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	CreatedAt string `json:"created_at"`
}

func (m *MemRequest) GetPageNum() int {
	return m.PageNum
}
func (m *MemRequest) GetPageSize() int {
	return m.PageSize
}
