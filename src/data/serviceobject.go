package data

type ApiRequest struct {
	Api  *string
	Data *string
}

type ApiResponse struct {
	Status int
	Data   interface{}
	Error  string
}
