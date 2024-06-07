package handler

import (
	"github.com/TheAlgorithms/Go/structure/set"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/locales/de"
)

var set1 = set.New[int]()
var set2 = set.New[int]()
var set3 = set.New[int]()
var sli1 = make([]int, 0)
var sli2 = make([]int, 0)
var sli3 = make([]int, 0)

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
		if set1.In(rby.UserId) && set1.Len() < 3 {
			c.AbortWithStatus(409)
			return
		}
		if set1.Len() >= 3 {
			for _, v := range sli1 {
				set1.Delete(v)
			}
			sli1 = sli1[:0]
		}
		sli1 = append(sli1, rby.UserId)
		set1.Add(rby.UserId)
		for {
			if set1.Len() == 3 {
				c.JSON(200, set1.GetItems())
				return
			}
		}
	case "象棋":
		if set2.In(rby.UserId) && set2.Len() < 3 {
			c.AbortWithStatus(409)
			return
		}
		if set2.Len() >= 3 {
			for _, v := range sli2 {
				set2.Delete(v)
			}
			sli2 = sli1[:0]
		}
		sli2 = append(sli2, rby.UserId)
		set2.Add(rby.UserId)
		for {
			if set2.Len() == 3 {
				c.JSON(200, set2.GetItems())
				return
			}
		}
		break
	case "麻将":
		if set3.In(rby.UserId) && set3.Len() < 3 {
			c.AbortWithStatus(409)
			return
		}
		if set3.Len() >= 3 {
			for _, v := range sli3 {
				set3.Delete(v)
			}
			sli3 = sli1[:0]
		}
		sli3 = append(sli3, rby.UserId)
		set3.Add(rby.UserId)
		for {
			if set3.Len() == 3 {
				c.JSON(200, set3.GetItems())
				return
			}
		}
	}
}

// match has bug
func match(c *gin.Context, Set set.Set[int], num int, id int, slice []int) {
	if Set.Len() >= num {
		color.Red("%#v", sli1)
		for _, v := range slice {
			Set.Delete(v)
		}
		slice = slice[:0]
		color.Red("%#v", sli1)
		color.Red("%#v", Set.GetItems())
	}
	//
	slice = append(slice, id)
	color.Red("%#v", sli1)
	Set.Add(id)
	//
	for {
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
			//
			j := 0
			for _, v := range sli1 {
				if v != rby.UserId {
					sli1[j] = v
					j++
				}
			}
			sli1 = sli1[:j]
			//
			c.AbortWithStatus(200)
			return
		}
		break
	case "象棋":
		if set2.In(rby.UserId) {
			set2.Delete(rby.UserId)
			//
			set1.Delete(rby.UserId)
			j := 0
			for _, v := range sli2 {
				if v != rby.UserId {
					sli2[j] = v
					j++
				}
			}
			sli2 = sli2[:j]
			//
			c.AbortWithStatus(200)
			return
		}
		break
	case "麻将":
		if set3.In(rby.UserId) {
			set3.Delete(rby.UserId)
			//
			j := 0
			for _, v := range sli3 {
				if v != rby.UserId {
					sli3[j] = v
					j++
				}
			}
			sli3 = sli3[:j]
			//
			c.AbortWithStatus(200)
			return
		}
		break
	}
	c.AbortWithStatus(400)
}
