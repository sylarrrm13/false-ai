package controller

import (
	"fmt"
	"models"
	"net/http"
	"req"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CardController struct{}

func (CardController) GetCardList(c *gin.Context) {
	// Code to get card list
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err})
			return
		}
	}()
	pn := c.Query("pageNum")
	ps := c.Query("pageSize")

	pageNum, err := strconv.Atoi(pn)
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		panic(err)
	}
	//查询数据库

	var cardList []models.CardUsage
	var total int64
	//分页查询ModelCate
	err = service.CardService{}.GetCardList(&cardList, &total, pageSize, pageNum)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "查询失败"})
		return

	} else {
		c.JSON(200, gin.H{"code": 200, "data": gin.H{"list": cardList, "total": total, "pageNum": pageNum, "pageSize:": pageSize}})
	}
}

//增加卡密

func (CardController) GenCard(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err})
			return
		}
	}()
	var cardGen = req.CardGen{}
	err := c.ShouldBindJSON(&cardGen)
	if err != nil {
		panic(err)
	}
	//新增
	addr := new(string)
	err = service.CardService{}.GenCard(cardGen, addr)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0, "addr": *addr}})
	}
}

//删除卡密

func (CardController) DeleteCard(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err})
			return
		}
	}()
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"code": 500, "msg": "id不能为空"})
		return
	}
	result := service.CardService{}.DelCard(id)
	if result != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "删除失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
	}
}
