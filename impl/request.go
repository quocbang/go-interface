package impl


type GetUserInfoRequest struct {
	UserID string 
}

type GetUserInfoReply struct {
	LastName string
	FisrtName string
	UserAge int
}