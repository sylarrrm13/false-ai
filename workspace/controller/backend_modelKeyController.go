package controller

import (
	"config"
	"conn"
	"fmt"
	"models"
	"net/http"
	"req"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModelKeyController struct{}

func (ModelKeyController) GetModelKeyList(c *gin.Context) {
	//
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

	var modeKeyList []models.ModelKeyCate

	//分页查询ModelKey
	//需要关联查询 获取对应model
	result := conn.DB.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&modeKeyList)
	//查询总数
	var total int64
	conn.DB.Model(&models.ModelKeyCate{}).Count(&total)
	if result.Error != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "查询失败"})
		return

	} else {
		c.JSON(200, gin.H{"code": 200, "data": gin.H{"list": modeKeyList, "total": total, "pageNum": pageNum, "pageSize:": pageSize}})
	}
}

func (ModelKeyController) DeleteModelKey(c *gin.Context) {
	//删除用户
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"code": 500, "msg": "id不能为空"})
		return
	}
	//删除
	//获取modelcateid
	var modelKey models.ModelKeys
	conn.DB.Where("id = ?", id).First(&modelKey)
	conn.DB.Where("id = ?", id).Delete(&models.ModelKeys{})
	//删除对应在 config.modelMap中的模型
	modelCateAndModels, ok := config.ModelMap.Load(modelKey.ModelCateID)
	if ok {
		modelCateAndModels.(config.ModelCateAndModels).Models.Delete(modelKey.ID)

		service.ModelKeyService{}.ReIndexWeight(modelCateAndModels.(config.ModelCateAndModels))

	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
}
func (ModelKeyController) UpdateModelKey(c *gin.Context) {
	//根据ID更新
	var modelKey models.ModelKeys
	err := c.ShouldBindJSON(&modelKey)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "参数错误"})
		return
	}
	var originModelKey models.ModelKeys
	conn.DB.Where("id = ?", modelKey.ID).First(&originModelKey)

	result := conn.DB.Model(&models.ModelKeys{}).Where("id = ?", modelKey.ID).Updates(&modelKey)
	if result.Error != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "更新失败"})
		return
	}

	modelCateAndModels, ok := config.ModelMap.Load(originModelKey.ModelCateID)
	if ok {
		modelCateAndModels.(config.ModelCateAndModels).Models.Store(modelKey.ID, modelKey)
		//遍历并且打印所有 modelCateAndModels.(config.ModelCateAndModels).Models
		//判断weight 是否变化
		if originModelKey.Weight != modelKey.Weight || *originModelKey.Enable != *modelKey.Enable {

			service.ModelKeyService{}.ReIndexWeight(modelCateAndModels.(config.ModelCateAndModels))

		}

	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
}

func (ModelKeyController) AddModelKey(c *gin.Context) {
	//新增
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
			return
		}
	}()

	modelKeyList := req.ModelKeyCateList{}
	err := c.ShouldBindJSON(&modelKeyList)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "参数错误"})
		return

	}
	//循环遍历 modelKeyList中的ModelCateID 和KEY 生成modelKeys
	keys := []models.ModelKeys{}
	for _, v := range modelKeyList.Key {
		for _, vv := range modelKeyList.ModelCateID {
			var modelKey models.ModelKeys
			modelKey.ModelCateID = vv
			modelKey.Key = v
			modelKey.ApiAddr = modelKeyList.ApiAddr
			modelKey.Weight = modelKeyList.Weight
			modelKey.Enable = &modelKeyList.Enable
			keys = append(keys, modelKey)

		}

	}
	//批量插入

	result := conn.DB.Create(&keys)

	for _, v := range modelKeyList.ModelCateID {
		modelCateAndModels, ok := config.ModelMap.Load(v)
		for _, key := range keys {

			if ok {
				modelCateAndModels.(config.ModelCateAndModels).Models.Store(key.ID, key)

			}
		}
		if ok {
			service.ModelKeyService{}.ReIndexWeight(modelCateAndModels.(config.ModelCateAndModels))

		}
	}

	if result.Error != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
		return

	}
	//更新config.ModelMap

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
	// modelKeyLis
	// modelKeys := []models.ModelKeys{}

	// var modelKey models.ModelKeys
	// if modelKey.ID != 0 {
	// 	modelKey.ID = 0
	// }
	// result := conn.DB.Create(&modelKey)
	// if result.Error != nil {
	// 	c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
	// 	return
	// }
	// c.JSON(200, gin.H{"code": 200, "msg": "新增成功"})
}
