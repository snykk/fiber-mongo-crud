package datatransfer

type TodoRequest struct {
	Id       string `json:"id"`
	Activity string `json:"activity" validate:"required"`
	Priority string `json:"priority" validate:"required"`
}

type TodoUpdateRequest struct {
	Activity string `json:"activity"`
	Priority string `json:"priority"`
	IsDone   *bool  `json:"is_done"`
}
