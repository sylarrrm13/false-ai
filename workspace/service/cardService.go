package service

import (
	"conn"
	"errors"
	"fmt"
	"models"
	"os"
	"req"
	"strconv"
	"time"
	"tools"
)

type CardService struct{}

func (CardService) GetCardTypeNameList(cardTypeNameList *[]models.CardTypeNameList) {
	var cardTypeList []models.CardType
	conn.DB.Debug().Select("id,name").Find(&cardTypeList).Scan(cardTypeNameList)
}

func (CardService) GetCardTypeList(cardTypeList *[]models.CardType, total *int64, pageSize int, pageNum int) error {
	result := conn.DB.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("sort asc").Find(cardTypeList)
	//查询总数
	conn.DB.Model(&models.CardType{}).Count(total)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (CardService) GetCardList(cardList *[]models.CardUsage, total *int64, pageSize int, pageNum int) error {
	result := conn.DB.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(cardList)
	//查询总数
	conn.DB.Model(&models.Card{}).Count(total)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (CardService) GenCard(CardGen req.CardGen, addr *string) error {
	var cardType models.CardType
	conn.DB.Debug().Where("id = ?", CardGen.CardTypeID).First(&cardType)
	if cardType.ID == 0 {
		return errors.New("卡密类型不存在")
	}

	var cardList = []models.Card{}
	var file *os.File
	var err error
	if CardGen.Export {
		*addr = strconv.FormatInt(time.Now().UnixNano(), 10)
		file, err = os.OpenFile("../files/card/"+*addr+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	for i := 0; i < CardGen.Number; i++ {
		card := models.Card{}
		card.CardNo, _ = tools.GenerateKey()
		fmt.Println("卡密", card.CardNo)
		card.CardTypeId = CardGen.CardTypeID
		card.Use = new(int)
		*card.Use = 0
		card.UserId = nil
		cardList = append(cardList, card)
		if CardGen.Export {
			if _, err := fmt.Fprintln(file, card.CardNo); err != nil {
				return err
			}
		}

	}
	//批量插入
	//将卡密导出置当前时间戳的txt文件 并且返回文件路径

	fmt.Println("卡密", cardList)
	result := conn.DB.Create(&cardList)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (CardService) DelCard(id string) error {
	//根据ID删除卡密
	result := conn.DB.Where("id = ?", id).Delete(&models.Card{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
