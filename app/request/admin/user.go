package admin

type UserSearchParam struct {
	Account  string
	Name     string
	Mobile   string
	PageNum  int
	PageSize int
}

func (up *UserSearchParam) GetPageNum() int {
	return up.PageNum
}
func (up *UserSearchParam) GetPageSize() int {
	return up.PageSize
}
