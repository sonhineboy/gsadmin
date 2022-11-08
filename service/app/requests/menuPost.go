package requests

type MenuPost struct {
	Meta      map[string]interface{} `json:"meta"`
	Component string                 `json:"component"`
	Name      string                 `json:"name"`
	ParentId  uint                   `json:"parentId"`
	Path      string                 `json:"path"`
	Redirect  string                 `json:"redirect"`
	Id        string                 `json:"id"`
	ApiList   []map[string]string    `json:"apiList"`
	Sort      int                    `json:"sort"`
}

type MenuDel struct {
	Ids []uint `json:"ids"`
}
