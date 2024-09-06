package requests

type DeleteDemo2Request struct {
	Ids []int `binding:"required"`
}

type Demo2Request struct {
	Name           string    `json:"name" binding:"required"`
	Age            int8      `json:"age" binding:"required"`
	
}