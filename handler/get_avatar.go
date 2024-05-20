package handler

import (
	"Game/tools"
	"github.com/gin-gonic/gin"
)

func GetAvatar(c *gin.Context) {
	id := c.Param("id")
	path := tools.Conf.RootPath.Path + "/public/avatar/" + id + ".png"
	c.File(path)
}
