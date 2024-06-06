package handler

import (
	"github.com/TheAlgorithms/Go/structure/deque"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var Queue1 = deque.New[int]()
var Queue2 = deque.New[int]()
var Queue3 = deque.New[int]()

func GameMatch(c *gin.Context) {
	type body struct {
		GameType string
		UserId   int
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		return
	}
	//Subscriber1

	switch rby.GameType {
	case "斗地主":
		Queue1.EnqueueRear(rby.UserId)
		match(c, Queue1, 3)
		break
	case "象棋":
		Queue2.EnqueueRear(rby.UserId)
		break
	case "麻将":
		Queue3.EnqueueRear(rby.UserId)
	}
}

func match(c *gin.Context, queue *deque.DoublyEndedQueue[int], num int) {
	var res []int
	if queue.Length() == num+1 {
		for range num {
			_, _ = queue.DequeueFront()
		}
	}
	for {
		color.Red("%d \n", queue.Length())
		if queue.Length() == num {
			for range num {
				v, _ := queue.DequeueFront()
				res = append(res, v)
				queue.EnqueueRear(v)
			}
			c.JSON(200, res)
			return
		}
	}
}
