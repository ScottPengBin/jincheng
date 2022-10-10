package base

type ReqPaginateParam interface {
	GetPageNum() int
	GetPageSize() int
}
