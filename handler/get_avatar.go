package handler

import (
	"Game/tools"
	"github.com/gin-gonic/gin"
	"regexp"
)

// GetAvatar preview file or download file as attachment
func GetAvatar(c *gin.Context) {
	id := c.Param("id")
	path := tools.Conf.RootPath.Path + "/public/avatar/" + id + ".png"
	//
	c.File(path)
	// download file or preview file
	url := c.Request.URL.Path
	patter := "download"
	regex := regexp.MustCompile(patter)
	matches := regex.FindAllString(url, -1)
	//color.Red("%d", len(matches))
	if len(matches) == 1 {
		c.Header("Content-Disposition", "attachment;filename="+id+".png")
	}
}
