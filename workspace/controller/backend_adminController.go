package controller

import (
	"config"
	"conn"
	"fmt"
	"models"
	"net/http"
	"req"
	"time"
	"tools"
	"utils"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

// 查询所有管理员
func (a AdminController) GetAdmins(c *gin.Context) {
	var admin []models.Admin
	conn.DB.Debug().Preload("AdminRole").Find(&admin)
	fmt.Println(admin)
	c.JSON(http.StatusOK, gin.H{"result": admin})

}

func (a AdminController) Login(c *gin.Context) {
	// 获取username 和 password json数据

	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err})
			return
		}
	}()

	var req req.CreateAdminReq
	if err := c.BindJSON(&req); err != nil {
		panic("参数错误")
	}
	// 查询数据库
	admin := models.Admin{}
	result := conn.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&admin)
	if result.Error != nil {
		panic("用户名或密码错误")

	}
	// 生成token
	access_token, _, _ := utils.CreateToken("access", config.ConfigList.Jwt.Secret, admin.Id, admin.Role, 0)
	refresh_token, _, _ := utils.CreateToken("refresh", config.ConfigList.Jwt.Secret, admin.Id, admin.Role, 0)
	// 保存token models.user_token
	userToken := models.AdminUserToken{}
	userToken.UserId = admin.Id
	userToken.Token = refresh_token
	userToken.CreateAt = utils.LocalTime(time.Now())
	userToken.ExpiredAt = utils.LocalTime(time.Now())
	userToken.Expired = 0
	conn.DB.Create(&userToken)
	// 将refresh_token放入redis 并设置2小时到期一次
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"access_token": access_token, "refresh_token": refresh_token}})
}

func (a AdminController) Logout(c *gin.Context) {
	fmt.Println("logout")
}

// 更新密码
func (a AdminController) UpdatePassword(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err})
			return
		}
	}()
	var req req.UpdateAdminReq
	if err := c.BindJSON(&req); err != nil {
		panic("参数错误")
	}
	// 查询数据库
	admin := models.Admin{}
	//数据库中查找当前用户的role

	result := conn.DB.Joins("AdminRole").Where("username = ? ", req.Username).First(&admin)
	if result.Error != nil {
		panic("用户名错误")
	}
	//判断角色
	if admin.AdminRole.Name == "admin" {
		//判断原密码是否正确
		fmt.Println(admin.Password, req.OriPass)
		if admin.Password != req.OriPass {
			panic("原密码错误")
		}
		conn.DB.Model(&admin).Update("password", req.Password)
		//删除之前数据库中的refresh_token
		conn.DB.Where("userid = ?", admin.Id).Delete(&models.AdminUserToken{})
		//生成新的 access_token和 refresh_token
		access_token, refresh_token := tools.GenerateToken(admin.Id, admin.Role, 0)
		//创建新的refresh_token

		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"access_token": access_token, "refresh_token": refresh_token}})
	} else {
		//如果要修改的用户使guest 则直接修改 admin和guest为内置类型无法创建修改属性
		conn.DB.Model(&admin).Update("password", req.Password)
		//删除对应的
		conn.DB.Where("userid = ?", admin.Id).Delete(&models.AdminUserToken{})
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok"})
	}
	//判断原密码是否正确
	//获取原密码和新密码

	//更新密码
}

func (a AdminController) RefreshToken(c *gin.Context) {
	//获取refresh_token
	var req req.RefreshToken
	err := c.ShouldBindJSON(&req)
	//验证通过字段验证
	fmt.Println(err)
	claims, err := utils.ParseToken(req.RefreshToken, config.ConfigList.Jwt.Secret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "token invalid1"})
		return
	}

	//判断是否为refresh token
	if claims["type"] != "refresh" {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "token invalid2"})
		return
	}

	if int(claims["role"].(float64)) != 1 {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "token invalid3"})
		return
	}
	//对比 claims中的生成日期 和 user表中的修改日期，如果修改日期小于 生成日期 则生成新的token否则返回错误

	//查询数据库
	admin := models.Admin{}
	conn.DB.Where("id = ?", claims["userId"]).First(&admin)
	if admin.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "token invalid4"})
	}

	// Convert admin.UpdatedAt to Unix timestamp
	adminUpdatedAt := time.Time(admin.UpdatedAt).Unix()
	// Convert claims["iat"] to Unix timestamp

	claimsIat := int64(claims["iat"].(float64))
	// Compare the two timestamps
	if adminUpdatedAt > claimsIat {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "token invalid5"})
		return
	} else {
		access_token := tools.GenerateAccessToken(admin.Id, admin.Role, 0)
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"access_token": access_token}})
		return
	}

}
