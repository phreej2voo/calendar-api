package request

type CreateBaby struct {
	Name     string     `json:"name" validate:"required"`
	Birthday CustomDate `json:"birthday" validate:"required"`
	Gender   string     `json:"gender" validate:"required"`
}

type GetBaby struct {
	BabyID int `query:"babyId" validate:"required"`
}

type UpdateBaby struct {
	Name     string     `json:"name"`
	Birthday CustomDate `json:"birthday"`
	Gender   string     `json:"gender"`
}
