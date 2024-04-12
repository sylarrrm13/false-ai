package service

import (
	"models"
	"utils"

	"gorm.io/gorm"
)

// 更新用户积分 需要 ID +/- 以及 数据库指针
type UserService struct{}

func (UserService) UpdateUserBill(id int, coin int, expired *utils.LocalTime, coins_expired *utils.LocalTime, bill bool, tx *gorm.DB) error {
	//TODO
	//更新用户根据ID 更新用户余额
	//加上更新到期时间
	//更新用户积分
	var updates map[string]interface{}
	if !bill {
		updates = map[string]interface{}{
			"coins":         gorm.Expr("coins + ?", coin),
			"expired":       expired,
			"coins_expired": coins_expired,
		}

	} else {
		//仅更新余额
		updates = map[string]interface{}{
			"coins": gorm.Expr("coins + ?", coin),
		}
	}
	err := tx.Debug().Model(&models.UserBillInfo{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}
