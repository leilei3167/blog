package v1

import (
	"blog/model"
	"blog/pkg/er"
	"net/http"

	"github.com/gin-gonic/gin"
)

//对于tag的增删改查

//获取标签列表(考虑分页)
func GetTags(c *gin.Context) {
	//通过name来查询,获取查询字段
	name := c.Query("name")
	data := model.GetTags(name)
	c.JSON(http.StatusOK, gin.H{
		"code": er.SUCCESS,
		"msg":  er.GetMsg(er.SUCCESS),
		"data": data,
	})

}

//新增标签
func AddTag(c *gin.Context) {

}

//改
func UpdateTag(c *gin.Context) {

}

//查
func DeleteTag(c *gin.Context) {

}
