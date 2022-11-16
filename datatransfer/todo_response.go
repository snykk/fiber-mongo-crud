package datatransfer

type TodoResponse struct {
	Id       string `json:"id"`
	Activity string `json:"activity"`
	Priority string `json:"priority"`
	IsDone   bool   `json:"is_done"`
}
