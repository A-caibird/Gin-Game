package handler

import (
	"github.com/TheAlgorithms/Go/structure/set"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/locales/de"
)

var set1 = set.New[int]()
var set2 = set.New[int]()
var set3 = set.New[int]()
var sli1 = make([]int, 3)
var sli2 = make([]int, 2)
var sli3 = make([]int, 4)

func GameMatch(c *gin.Context) {
	type body struct {
		GameType string
		UserId   int
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		return
	}
	//
	switch rby.GameType {
	case "斗地主":
		if set1.In(rby.UserId) {
			c.AbortWithStatus(409)
			return
		}
		sli1 = append(sli1, rby.UserId)
		match(c, set1, 3, rby.UserId)
		break
	case "象棋":
		if set2.In(rby.UserId) {
			c.AbortWithStatus(409)
			return
		}
		sli2 = append(sli2, rby.UserId)
		match(c, set2, 2, rby.UserId)
		break
	case "麻将":
		if set3.In(rby.UserId) {
			c.AbortWithStatus(409)
			return
		}
		sli3 = append(sli3, rby.UserId)
		match(c, set1, 4, rby.UserId)
	}
}

func match(c *gin.Context, Set set.Set[int], num int, id int) {
	Set.Add(id)
	if Set.Len() == num+1 {
		for _, v := range sli1 {
			if v == id {
				continue
			}
			Set.Delete(v)
		}
	}
	for {
		//color.Red("%d \n", Set.Len())
		if Set.Len() == num {
			c.JSON(200, Set.GetItems())
			return
		}
	}
}
func CancelMatch(c *gin.Context) {
	type body struct {
		GameType string
		UserId   int
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		return
	}
	//
	switch rby.GameType {
	case "斗地主":
		if set1.In(rby.UserId) {
			set1.Delete(rby.UserId)
			c.AbortWithStatus(200)
			return
		}
		break
	case "象棋":
		if set2.In(rby.UserId) {
			set2.Delete(rby.UserId)
			c.AbortWithStatus(200)
			return
		}
		break
	case "麻将":
		if set3.In(rby.UserId) {
			set3.Delete(rby.UserId)
			c.AbortWithStatus(200)
			return
		}
		break
	}
	c.AbortWithStatus(400)
}
