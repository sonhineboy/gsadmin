package global

type List struct {
	Page     int                    `json:"page" form:"page"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Where    map[string]interface{} `json:"where" form:"where"`
}

type Del struct {
	Ids []uint `json:"id"`
}
