package controller

import (
	"config"
	"conn"
	"fmt"
	"models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type ModelCateController struct{}

func (ModelCateController) GetModelCateList(c *gin.Context) {
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
	filter := c.Query("filter")
	//如果filter 是 1 则代表获取所有的模型名以及ID
	if filter == "1" {
		var modelCateList []models.ModelCate
		var modelCateIDName []models.ModelCateIDName
		conn.DB.Debug().Select("id,name").Find(&modelCateList).Scan(&modelCateIDName)
		c.JSON(200, gin.H{"code": 200, "data": modelCateIDName})
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

	var modeCateList []models.ModelCate

	//分页查询ModelCate
	result := conn.DB.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("sort asc").Find(&modeCateList)
	//查询总数
	var total int64
	conn.DB.Model(&models.ModelCate{}).Count(&total)
	if result.Error != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "查询失败"})
		return

	} else {
		c.JSON(200, gin.H{"code": 200, "data": gin.H{"list": modeCateList, "total": total, "pageNum": pageNum, "pageSize:": pageSize}})
	}
}

func (ModelCateController) DeleteModelCate(c *gin.Context) {
	//删除用户
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"code": 500, "msg": "id不能为空"})
		return
	}

	//获取对应的modelcate 是否是工具模型
	var modelCate models.ModelCate
	conn.DB.Where("id = ?", id).First(&modelCate)
	//删除
	conn.DB.Where("model_cate_id = ?", id).Delete(&models.ModelKeys{})
	//删除对应模型KEY
	conn.DB.Delete(&models.ModelCate{}, id)
	//删除对应在 config.modelMap中的模型

	config.ModelMap.Delete(id)
	config.ToolMap.Delete(id)
	config.KeyWeightIndices.Delete(id)
	if modelCate.Model == "gpt-4-gizmo" {
		idInt, ok := strconv.ParseInt(id, 10, 64)
		if ok != nil {
			config.GTPsSlice.Remove(int(idInt))
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})
}

// 新增模型类
func (ModelCateController) AddModelCate(c *gin.Context) {
	//新增
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err})
			return
		}

	}()
	var modelCate models.ModelCate
	err := c.ShouldBindJSON(&modelCate)
	fmt.Println(modelCate)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "参数错误"})
		return
	}
	//增加在ModelMap中的模型
	conn.DB.Create(&modelCate)
	config.ModelMap.Store(modelCate.ID, config.ModelCateAndModels{
		ModelCate: modelCate,
		Models:    &sync.Map{},
	})
	if modelCate.Model == "gpt-4-gizmo" {
		config.GTPsSlice.Append(modelCate.ID)

	}
	// config.ModelMap.Store(modelCate.ID, modelCate)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})

}

// 修改模型类
func (ModelCateController) UpdateModelCate(c *gin.Context) {
	//更新
	var modelCate models.ModelCate
	//获取ID

	err := c.ShouldBindJSON(&modelCate)

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 500, "msg": "参数错误"})
		return
	}
	//通过ID 更新字段 而不是新增
	//更新

	result := conn.DB.Model(&models.ModelCate{}).Where("id = ?", modelCate.ID).Updates(modelCate)
	if result.Error != nil {
		c.JSON(200, gin.H{"code": 404, "msg": "更新失败"})
	}
	//更新在ModelMap中的模型
	// 先从config.ModelMap中取出来
	// modelCateAndModels, ok := config.ModelMap.Load(modelCate.ID)
	// if ok {
	// 	modelCateAndModels.(config.ModelCateAndModels).ModelCate = modelCate
	// 	config.ModelMap.Store(modelCate.ID, ModelCateAndModels)
	// }

	modelCateAndModels, ok := config.ModelMap.Load(modelCate.ID)
	if ok {
		config.ModelMap.Store(modelCate.ID, config.ModelCateAndModels{
			ModelCate: modelCate,
			Models:    modelCateAndModels.(config.ModelCateAndModels).Models,
		})

	}
	//如果原来不是gizmo模型 现在是 则添加到GTPsSlice中 如果原来是现在不是 则删除
	if modelCate.Model == "gpt-4-gizmo" {
		config.GTPsSlice.Append(modelCate.ID)
	} else {
		config.GTPsSlice.Remove(modelCate.ID)
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"status": 0}})

}
