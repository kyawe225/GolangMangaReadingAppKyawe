package dtos

type responseDto[T any] struct {
	Status  string
	Message string
	Model   T
}

func NewResponseDto[T any](status string, message string, model T) responseDto[T] {
	return responseDto[T]{
		Status:  status,
		Message: message,
		Model:   model,
	}
}
