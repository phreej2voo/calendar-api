package request

type Knowledges struct {
	LabelID  int    `query:"labelId" validate:"required"`
	BabyID   int    `query:"babyId" validate:"required"`
	Category string `query:"category" validate:"required"`
	Days     int    `query:"days"`
	Page     int    `query:"page"`
	Size     int    `query:"size"`
}

func NewKnowleges() *Knowledges {
	return &Knowledges{
		Page: 1,
		Size: 10,
		Days: -1,
	}
}
