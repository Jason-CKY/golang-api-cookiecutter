package schemas

type SortTaskRequestParams struct {
	TaskIds []string `form:"task_ids"`
	Status  string   `param:"status"`
}
