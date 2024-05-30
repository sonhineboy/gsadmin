package requests

type DeleteOrderRequest struct {
	Ids []int `binding:"required"`
}

type OrderRequest struct {
	UserName string `json:"userName" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}
