package base

type ReqPaginateParam struct {
	Current int `json:"current"`
	Size    int `json:"size"`
}
