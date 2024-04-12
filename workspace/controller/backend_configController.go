package controller

import (
	"config"
	"conn"
	"context"
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"req"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ConfigController struct{}

func (c ConfigController) GetConfig(r *gin.Context) {
	//获取配置信息
	//获取邮箱配置信息
	//先从redis缓存中读取
	defer func() {
		if err := recover(); err != nil {
			r.JSON(500, gin.H{"code": 500, "msg": "未知错误"})
		}
	}()
	listname := r.Param("listname") // Fix: Access the parameter from the context
	//判断listname是否存在
	if listname == "" {
		r.JSON(200, gin.H{"code": 400, "msg": "listname不能为空"})
		return
	}
	res := req.ReqMapFunc("CodelkupWithoutRef")
	value, err := conn.RedisPool.Get(context.Background(), "listname-"+listname).Result()
	if err == nil {
		resNew, ok := res.(*req.CodelkupWithoutRef)
		if ok {
			resNew.Listname = listname
			resNew.Value = value
			r.JSON(200, gin.H{"code": 200, "data": res})
			return
		}

	}
	codelkup := []models.Codelkup{}
	//先从redis缓存中读取

	conn.DB.Debug().Select("Listname,Value").Where("listname = ?", listname).Find(&codelkup).Scan(res)
	//序列化res.value并存储到redis中
	resNew, ok := res.(*req.CodelkupWithoutRef)
	if !ok {
		r.JSON(500, gin.H{"code": 500, "msg": "类型转换失败"})
		return
	}

	conn.RedisPool.Set(context.Background(), "listname-"+listname, resNew.Value, 2*time.Hour)
	r.JSON(200, gin.H{"code": 200, "data": res})
}

// 更新配置
func (c ConfigController) UpdateConfig(r *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			r.JSON(500, gin.H{"code": 500, "msg": "未知错误"})
		}
	}()

	listname := r.Param("listname") // Fix: Access the parameter from the context
	//判断listname是否存在
	if listname == "" {
		r.JSON(200, gin.H{"code": 400, "msg": "listname不能为空"})
		return
	}

	//使listname首字母大写

	err := conn.DB.Transaction(func(tx *gorm.DB) error {
		reqConfig := req.ReqMapFunc("update_" + listname)

		err := r.ShouldBindJSON(&reqConfig)
		if err != nil {

			fmt.Println("err:", err)
			return err
		}

		validate := validator.New()
		err = validate.Struct(reqConfig)
		if err != nil {
			fmt.Println("验证失败:", err)
			return err
		}
		//emailConfig 转换成字符串
		jsonBytes, err := json.Marshal(reqConfig)
		if err != nil {
			return err
		}

		jsonString := string(jsonBytes)
		fmt.Println("jsonString:", jsonString)
		//将jsonString 反序列化到emailConfig中 因为是外部的所以要枷锁

		err = config.SetConfigMap(listname, jsonString)

		if err != nil {
			return err
		}
		// 将jsonString 更新到 codelkup 中的listname email 的value中
		err = tx.Model(&models.Codelkup{}).Where("listname = ?", listname).Update("value", jsonString).Error
		//将新的配置信息存储到redis中
		conn.RedisPool.Set(context.Background(), "listname-"+listname, jsonString, 2*time.Hour)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {

		r.JSON(200, gin.H{"code": 500, "msg": "更新失败,请检查参数"})

	} else {
		r.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})

	}
}
