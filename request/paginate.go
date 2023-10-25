package request

type Paginate struct {
	Page int `query:"page"`
	Size int `query:"size"`
}
