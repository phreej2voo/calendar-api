package request

type AddPhotos struct {
	Links []string `json:"links,omitempty" validate:"required"`
}
