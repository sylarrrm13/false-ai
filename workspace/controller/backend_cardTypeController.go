package controller

import (
	"conn"
	"fmt"
	"models"
	"net/http"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CardTypeController struct{}

// GetCardList is a function to get card list
func (CardTypeController) GetCardTypeList(c *gin.Context) {
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
	filter := c.Query("filter")

	if filter == "1" {
		var cardTypeNameList []models.CardTypeNameList
		service.CardService{}.GetCardTypeNameList(&cardTypeNameList) // Remove the argument from the function call
		c.JSON(200, gin.H{"code": 200, "data": gin.H{"data": cardTypeNameList}})
		return
	}

	pageNum, err := strconv.Atoi(pn)
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		panic(err)
	}
	//查询数据库

	var cardTypeList []models.CardType
	var total int64
	//分页查询ModelCate
	err = service.CardService{}.GetCardTypeList(&cardTypeList, &total, pageSize, pageNum)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "查询失败"})
		return

	} else {
		c.JSON(200, gin.H{"code": 200, "data": gin.H{"list": cardTypeList, "total": total, "pageNum": pageNum, "pageSize:": pageSize}})
	}

}

// AddCardType is a function to add card type
func (CardTypeController) AddCardType(c *gin.Context) {
	// Code to add card type
	var cardType models.CardType
	err := c.ShouldBindJSON(&cardType)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "参数错误"})
		return
	}
	result := conn.DB.Create(&cardType)
	if result.Error != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "添加失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
	}
}

// UpdateCardType is a function to update card type
func (CardTypeController) UpdateCardType(c *gin.Context) {
	// Code to update card type
	var cardType models.CardType
	//获取ID

	err := c.ShouldBindJSON(&cardType)

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 500, "msg": "参数错误"})
		return
	}
	//通过ID 更新字段 而不是新增
	//更新

	result := conn.DB.Model(&models.CardType{}).Where("id = ?", cardType.ID).Updates(cardType)
	if result.Error != nil {
		c.JSON(200, gin.H{"code": 404, "msg": "更新失败"})
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
}

// DeleteCardType is a function to delete card type
func (CardTypeController) DeleteCardType(c *gin.Context) {
	// Code to delete card type
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"code": 500, "msg": "id不能为空"})
		return
	}
	//删除
	conn.DB.Where("id = ?", id).Delete(&models.CardType{})
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
}
