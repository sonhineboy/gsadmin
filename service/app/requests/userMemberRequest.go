package requests

type DeleteUserMemberRequest struct {
	Ids []int `binding:"required"`
}

type UserMemberRequest struct {
	NickName       string    `json:"nick_name" binding:"required"`
	RealName       string    `json:"real_name"`
	Age            int32     `json:"age"`
	Status         int8      `json:"status"`
	Online         string    `json:"online"`
	
}