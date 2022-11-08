package requests

type Role struct {
	Alias  string `json:"alias"`
	Label  string `json:"label"`
	Remark string `json:"remark"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
	Id     uint   `json:"id"`
}

type RoleDel struct {
	Ids []uint `json:"id"`
}

type RoleUpMenus struct {
	Id    uint   `json:"id" binding:"required"`
	Menus []uint `json:"menus" binding:"required"`
}
