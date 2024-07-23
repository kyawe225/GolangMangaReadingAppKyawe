package dtos

type ResponseDto[T any] struct {
	Status  string
	Message string
	Model   T
}

func NewResponseDto[T any](status string, message string, model T) ResponseDto[T] {
	return ResponseDto[T]{
		Status:  status,
		Message: message,
		Model:   model,
	}
}
