package form

type AddMomentForm struct {
	UserId  int    `form:"uid" json:"uid" binding:"required"`
	Content string `form:"content" json:"content" binding:"required,max=2048"`
}

type AddCommentForm struct {
	UserId   int    `form:"uid" json:"uid" binding:"required"`
	MomentId int    `form:"mid" json:"mid" binding:"required"`
	ParentId int    `form:"pid" json:"pid" binding:"gte=0"`
	Content  string `form:"content" json:"content" binding:"required,max=1024"`
}
