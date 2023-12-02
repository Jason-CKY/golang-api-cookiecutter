package schemas

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskSort struct {
	Id            string   `json:"id"`
	Status        string   `json:"status"`
	Sorting_order []string `json:"sorting_order"`
}
