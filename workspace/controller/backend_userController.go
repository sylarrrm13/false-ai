package controller

import (
	"conn"
	"errors"
	"fmt"
	"models"
	"net/http"
	"req"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserController struct{}

func (u UserController) GetUserInfoList(c *gin.Context) {
	//获取用户信息
	defer func() {
		if err := recover(); err != nil {
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

	var userInfoList []req.GetUserInfoListResp

	conn.DB.Debug().
		Table("User").
		Joins("INNER JOIN user_bill_info ON user_bill_info.userid = User.id").
		Select("User.Id,User.username,User.status,User.email,User.phone,user_bill_info.coins,user_bill_info.used_coins,user_bill_info.coins_expired,user_bill_info.expired").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).Scan(&userInfoList)
	//获取总数
	var total int64
	conn.DB.Model(&models.User{}).Count(&total)
	c.JSON(200, gin.H{"code": 200, "data": gin.H{"list": userInfoList, "total": total, "pageNum": pageNum, "pageSize:": pageSize}})

}

func (u UserController) DeleteUser(c *gin.Context) {
	//删除用户
	err := conn.DB.Transaction(func(tx *gorm.DB) error {
		userId := c.Query("id")
		if userId == "" {
			return errors.New("id is empty")
		}

		user := models.User{}
		userBillInfo := models.UserBillInfo{}

		if err := tx.Where("id = ?", userId).Delete(&user).Error; err != nil {
			return err
		}

		if err := tx.Where("UserId = ?", userId).Delete(&userBillInfo).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "删除用户失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
	}
}

func (u UserController) UpdateUser(c *gin.Context) {
	//更新用户信息

	var req req.GetUserInfoListResp
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "参数错误"})
		return
	}

	err = conn.DB.Transaction(func(tx *gorm.DB) error {
		if req.Id == 0 {
			return errors.New("id is empty")
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			fmt.Println("验证失败:", err)
			return err
		}
		//更新user
		user := models.User{}
		updates := map[string]interface{}{
			"Username": req.Username,
			"Status":   req.Status,
			"Email":    req.Email,
			"Phone":    req.Phone,
		}
		result := tx.Model(&user).Debug().Where("id = ?", req.Id).Updates(updates)
		if result.Error != nil {
			return result.Error
		}

		err = service.UserService{}.UpdateUserBill(req.Id, req.Coins, req.Expired, req.CoinsExpired, false, tx)

		if err != nil {
			return err
		}
		//返回异常

		return nil
	})
	if err != nil {

		c.JSON(200, gin.H{"code": 500, "msg": "更新失败,请检查参数"})

	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})

	}

}
