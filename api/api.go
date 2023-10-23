package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/palp1tate/MultiLevelCommentDemo/form"
	"github.com/palp1tate/MultiLevelCommentDemo/global"
	"github.com/palp1tate/MultiLevelCommentDemo/model"
)

func AddMoment(c *gin.Context) {
	addMomentForm := form.AddMomentForm{}
	if err := c.ShouldBind(&addMomentForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	moment := model.Moment{
		UserId:  addMomentForm.UserId,
		Content: addMomentForm.Content,
	}
	global.DB.Create(&moment)
	var user model.User
	global.DB.First(&user, addMomentForm.UserId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"mid":     moment.ID,
			"content": moment.Content,
			"author":  gin.H{"uid": user.ID, "nickname": user.Nickname, "avatar": user.Avatar},
		},
	})
	return
}

func AddComment(c *gin.Context) {
	addCommentForm := form.AddCommentForm{}
	if err := c.ShouldBind(&addCommentForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	comment := model.Comment{
		UserId:   addCommentForm.UserId,
		MomentId: addCommentForm.MomentId,
		Content:  addCommentForm.Content,
	}
	if addCommentForm.ParentId != 0 {
		comment.ParentId = &addCommentForm.ParentId
	}
	global.DB.Create(&comment)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func GetComments(c *gin.Context) {
	var commentTrees []model.CommentTree
	momentId := c.Query("mid")
	if momentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "mid不能为空",
		})
		return
	}
	mid, _ := strconv.Atoi(momentId)
	commentTrees = GetMomentComment(mid)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": commentTrees,
	})
}

func GetMomentCommentChild(pid int, commentTree *model.CommentTree) {
	var comments []model.Comment
	global.DB.Where("parent_id = ?", pid).Find(&comments)

	// 查询二级及以下的多级评论
	for i := 0; i < len(comments); i++ {
		var user model.User
		cid := int(comments[i].ID)
		uid := comments[i].UserId
		global.DB.Where("id = ?", uid).First(&user)
		child := model.CommentTree{
			CommentId: cid,
			Content:   comments[i].Content,
			Author:    gin.H{"uid": user.ID, "nickname": user.Nickname, "avatar": user.Avatar},
			CreatedAt: comments[i].CreatedAt.Format("2006-01-02 15:04"),
			Children:  []*model.CommentTree{},
		}
		commentTree.Children = append(commentTree.Children, &child)
		GetMomentCommentChild(cid, &child)
	}
}

func GetMomentComment(mid int) []model.CommentTree {
	var commentTrees []model.CommentTree
	var comments []model.Comment
	global.DB.Where("moment_id = ? AND parent_id IS NULL", mid).Order("id desc").Find(&comments)
	for _, comment := range comments {
		var user model.User
		cid := int(comment.ID)
		uid := comment.UserId
		global.DB.Where("id = ?", uid).First(&user)
		commentTree := model.CommentTree{
			CommentId: cid,
			Content:   comment.Content,
			Author:    gin.H{"uid": uid, "nickname": user.Nickname, "avatar": user.Avatar},
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04"),
			Children:  []*model.CommentTree{},
		}
		GetMomentCommentChild(cid, &commentTree)
		commentTrees = append(commentTrees, commentTree)
	}
	return commentTrees
}
