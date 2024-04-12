package service

import (
	"config"
	"fmt"
	"math/rand"
	"models"
	"time"
)

type ModelKeyService struct{}

func (ModelKeyService ModelKeyService) ReIndexWeight(modelCateAndModels config.ModelCateAndModels) {
	weightIndex := make([]int, 0)
	modelCateAndModels.Models.Range(func(key, value interface{}) bool {
		model := value.(models.ModelKeys)
		if *model.Enable == 0 {
			return true
		}
		for i := 0; i < model.Weight; i++ {
			weightIndex = append(weightIndex, model.ID)
		}
		return true
	})
	fmt.Println("权重更新了")
	fmt.Println(weightIndex)
	config.KeyWeightIndices.Store(modelCateAndModels.ModelCate.ID, weightIndex)
}

// 根据权重随机获取一个modelKey
func (modelKeyService ModelKeyService) GetModelKeyByWeight(modelCateAndModels config.ModelCateAndModels) (string, string) {
	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
	weightIndex, ok := config.KeyWeightIndices.Load(modelCateAndModels.ModelCate.ID)
	if !ok {
		return modelKeyService.GetModelKey(modelCateAndModels)
	}
	weightIndexSlice := weightIndex.([]int)
	if len(weightIndexSlice) == 0 {
		return modelKeyService.GetModelKey(modelCateAndModels)
	}
	// 随机获取一个索引
	index := randGen.Intn(len(weightIndexSlice))
	modelKeyID := weightIndexSlice[index]
	modelKey, ok := modelCateAndModels.Models.Load(modelKeyID)
	if !ok {
		return modelKeyService.GetModelKey(modelCateAndModels)
	}
	return modelKey.(models.ModelKeys).Key, modelKey.(models.ModelKeys).ApiAddr

}

// 随机获取一个modelKey
func (modelKeyService ModelKeyService) GetModelKey(modelCateAndModels config.ModelCateAndModels) (string, string) {
	var modelKey models.ModelKeys
	modelCateAndModels.Models.Range(func(_, value interface{}) bool {
		//取出一个就行
		modelKey = value.(models.ModelKeys)
		return false
	})
	return modelKey.Key, modelKey.ApiAddr
}
