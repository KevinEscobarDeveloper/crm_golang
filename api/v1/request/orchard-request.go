package request

type CreateOrchardRequest struct {
	Name string `json:"name" validate:"required"`
}
