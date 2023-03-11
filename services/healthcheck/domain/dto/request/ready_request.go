package request

type ReadyRequest struct {
	BaseReq
	Shout string `json:"shout" binding:"required"`
}

type BaseReq struct{}
