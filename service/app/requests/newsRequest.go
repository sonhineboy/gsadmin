package requests

type DeleteNewsRequest struct {
	Ids []int `binding:"required"`
}

type NewsRequest struct {
	Title          string    `json:"title" binding:"required"`
	Author         string    `json:"author" binding:"required"`
	Content        string    `json:"content" binding:"required"`
	
}