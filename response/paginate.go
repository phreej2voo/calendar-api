package response

type Paginate struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalPage  int `json:"totalPage"`
	TotalCount int `json:"totalItemCount"`
}

func NewPaginate(page, size, totalCount int) Paginate {
	return Paginate{
		Page:       page,
		Size:       size,
		TotalPage:  int(totalCount/size) + 1,
		TotalCount: totalCount,
	}
}
