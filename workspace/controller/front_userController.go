package controller

import (
	"bufio"
	"bytes"
	"config"
	"conn"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"models"
	"net/http"
	"req"
	"service"
	"strings"
	"time"
	"tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FrontUserController struct{}

func (fu FrontUserController) InitData(c *gin.Context) {
	//获取初始化信息
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(404, gin.H{"notice": "", "models": ""})
			return
		}
	}()
	notice := config.GetConfigMap("notice").(*config.Notice).Notice
	//遍历 config.ModelMap 转换成 切片 里面放置 ID 积分 和NAME
	var modelList []map[string]interface{}
	config.ModelMap.Range(func(key, value interface{}) bool {
		if *value.(config.ModelCateAndModels).ModelCate.Visible == 1 {
			modelList = append(modelList, map[string]interface{}{
				"id":          key,
				"name":        value.(config.ModelCateAndModels).ModelCate.Name,
				"upload":      value.(config.ModelCateAndModels).ModelCate.Upload,
				"upload_type": value.(config.ModelCateAndModels).ModelCate.UploadType,
				"bill":        value.(config.ModelCateAndModels).ModelCate.Bill,
				"coin":        value.(config.ModelCateAndModels).ModelCate.Coin,
				"sort":        value.(config.ModelCateAndModels).ModelCate.Sort,
			})
		}

		return true
	})
	//根据Sort排序
	for i := 0; i < len(modelList); i++ {
		for j := i + 1; j < len(modelList); j++ {
			if *modelList[i]["sort"].(*int) > *modelList[j]["sort"].(*int) ||
				(*modelList[i]["sort"].(*int) == *modelList[j]["sort"].(*int) && modelList[i]["id"].(int) > modelList[j]["id"].(int)) {
				modelList[i], modelList[j] = modelList[j], modelList[i]
			}
		}
	}
	userState, ok := c.Get("userState")
	if !ok || userState == "0" {
		userState = "0"
	} else {
		//获取用户ID
		userId, _ := c.Get("userId")
		//获取用户信息
		user := models.UserBillInfo{}
		conn.DB.Where("userId = ?", userId).First(&user)
		//如果用户有积分或者未到期则state为2
		if user.Coins > 0 {
			userState = "2"
		}
	}

	c.JSON(200, gin.H{"status": "Success", "data": gin.H{"notice": notice, "models": modelList, "userState": userState, "emailDuration": config.GetConfigMap("email_reg").(*config.EmailReg).Duration}})
}

//获取用户信息

func (fu FrontUserController) SendEmailCode(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"status": "Fail", "message": "发送失败"})
			return
		}
	}()
	//获取邮箱

	var req req.SendEmailCode

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "参数错误"})
		return
	}

	//如果REDIS中存在数据则返回验证过快
	//判断models.User 邮箱是否存在
	user := models.User{}
	conn.DB.Where("email = ?", req.Email).First(&user)
	if user.Id != 0 {
		c.JSON(200, gin.H{"status": "Fail", "message": "邮箱已存在"})
		return
	}

	_, err = conn.RedisPool.Get(context.Background(), "email-"+req.Email).Result()
	if err == nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "发送邮件过于频繁,请稍后再试"})
		return

	}
	code := new(string)
	duration := new(int)
	err = tools.SendMail(req.Email, code, duration)
	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "发送失败"})
		return
	}
	//将 code 和 email存储到 redis中，名字为 email-email,值为 code
	//设置过期时间为 duration

	conn.RedisPool.Set(context.Background(), "email-"+req.Email, *code, 300*time.Second)
	c.JSON(200, gin.H{"status": "Success", "message": "发送成功"})

}

func (fu FrontUserController) Register(c *gin.Context) {
	//注册
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"status": "Fail", "message": "注册失败"})
			return
		}
	}()
	var req req.RegUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "参数错误"})
		return
	}
	//验证邮箱验证码
	code, err := conn.RedisPool.Get(context.Background(), "email-"+req.Email).Result()
	if err != nil || code != req.Code {
		c.JSON(200, gin.H{"status": "Fail", "message": "验证码错误"})
		return
	}

	//验证通过 删除redis中的验证码
	conn.RedisPool.Del(context.Background(), "email-"+req.Email)
	//注册用户
	userid := new(int)
	err = conn.DB.Transaction(func(tx *gorm.DB) error {
		user := models.User{Username: req.Username, Email: req.Email, Password: req.Password}
		err := tx.Create(&user).Error
		if err != nil {
			return err
		} //插入到billinfo
		userBill := models.UserBillInfo{Userid: user.Id}
		*userid = user.Id
		err = tx.Create(&userBill).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "注册失败,用户名已存在"})
		return

	}
	//生成token
	access_token := tools.GenerateAccessToken(*userid, 3, 0)
	c.JSON(200, gin.H{"status": "Success", "data": gin.H{"access_token": access_token}})

}

func (fu FrontUserController) Login(c *gin.Context) {
	//登录
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"status": "Fail", "message": "登录失败"})
			return
		}
	}()
	var req req.LoginUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "参数错误"})
		return
	}
	//查询数据库
	user := models.User{}
	result := conn.DB.Where("username = ? and password = ?", req.Username, req.Password).First(&user)
	if result.Error != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "用户名或密码错误"})
		return
	}
	//生成token
	userBill := models.UserBillInfo{}
	conn.DB.Where("Userid = ?", user.Id).First(&userBill)
	//如果用户有积分或者未到期则state为2 后续要加上判断是否过期
	userState := "1"
	if userBill.Coins > 0 {
		userState = "2"
	}

	access_token := tools.GenerateAccessToken(user.Id, 3, 0)
	c.JSON(200, gin.H{"status": "Success", "data": gin.H{"access_token": access_token, "userState": userState}})
}

func (fu FrontUserController) BillInfo(c *gin.Context) {
	//获取用户Bill信息
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"status": "Fail", "message": "获取失败"})
			return
		}
	}()
	userId, _ := c.Get("userId")
	userBill := models.UserBillInfo{}
	conn.DB.Where("Userid = ?", userId).First(&userBill)
	c.JSON(200, gin.H{"status": "Success", "data": gin.H{"coins": userBill.Coins, "used_coins": userBill.UsedCoins, "coins_expired": userBill.CoinsExpired, "expired": userBill.Expired}})
}

// 充值用户积分 明天再优化 需要不限制到期类型的可以无限充值
func (fu FrontUserController) Recharge(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"status": "Fail", "message": "充值失败"})
			return
		}
	}()
	var reCharge []req.ReChargeUp
	var req req.ReCharge

	// Remove the invalid line
	// req.ReChargeUp

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "参数错误"})
		return
	}
	//根据卡密获取卡密信息
	card := models.Card{}
	conn.DB.Where("card_no= ?", req.CardNo).First(&card)
	//如果查找不到或者已经被使用则返回
	result := conn.DB.Debug().Raw("select duration,coin from card_type where id in (select card_type_id from card where card_no = ? and `use` = 0)", req.CardNo).Scan(&reCharge)
	//获取当前用户的到期时间和coins
	if result.Error != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "卡密不存在或者已经被使用"})
		return

	}
	if len(reCharge) == 0 {
		c.JSON(200, gin.H{"status": "Fail", "message": "卡密不存在或者已经被使用"})
		return
	}

	//用户的bill信息
	userId, _ := c.Get("userId")
	user_bill := models.UserBillInfo{}
	conn.DB.Where("userId = ?", userId).First(&user_bill)
	updates := map[string]interface{}{}

	//将 user_bill.CoinsExpired 转换成unix时间戳
	// expiredStamp := time.Time(*user_bill.CoinsExpired).Unix()
	canContinue := 0
	if user_bill.CoinsExpired == nil {
		if reCharge[0].Duration == -1 {
			//可以充值
			canContinue = 2

		} else {
			//如果积分未消耗完则不能充值
			if user_bill.Coins > 0 {
				c.JSON(200, gin.H{"status": "Fail", "message": "积分尚未消耗完,无法充值不同类型卡密"})

			} else {
				canContinue = 1
				//可以充值
			}
		}
	} else {
		expiredStamp := time.Time(*user_bill.CoinsExpired).Unix()
		//取当前时间戳
		timeStamp := time.Now().Unix()
		if user_bill.Coins > 0 && expiredStamp > timeStamp {
			c.JSON(200, gin.H{"status": "Fail", "message": "积分尚未消耗完,无法充值不同类型卡密"})

		} else {
			canContinue = 1

			//可以充值
		}
	}

	if canContinue == 0 {
		return
	} else {
		err := conn.DB.Transaction(func(tx *gorm.DB) error {
			result := tx.Model(&card).Where("card_no = ? and `use`=0", req.CardNo).Updates(map[string]interface{}{
				"use":     1,
				"user_id": userId, // 替换为你需要更新的另一个字段名和新值
			})
			if result.Error != nil || result.RowsAffected == 0 {
				//卡密不存在或者已经被使用
				return errors.New("1")
			}

			//后续可能会增加余额缓存更新模式 现在采用的是持久化更新的方式
			if canContinue == 2 {
				updates["coins"] = gorm.Expr("coins + ?", reCharge[0].Coin)
			} else {
				updates["coins"] = reCharge[0].Coin
			}
			if reCharge[0].Duration == -1 {
				updates["coins_expired"] = nil
			} else {
				expirationDate := time.Now().AddDate(0, 0, reCharge[0].Duration)
				updates["coins_expired"] = expirationDate.Format("2006-01-02 15:04:05")
			}
			result = tx.Model(&models.UserBillInfo{}).Where("userId = ?", userId).Updates(updates)

			if result.Error != nil {
				return result.Error
			}
			return nil

		})
		//取ERR信息
		if err != nil {
			c.JSON(200, gin.H{"status": "Fail", "message": "充值失败"})
			return
		}
		//查询更新后的到期时间与coins

		c.JSON(200, gin.H{"status": "Success", "message": "充值成功", "data": gin.H{"coins_expired": updates["coins_expired"], "coins": reCharge[0].Coin, "update": canContinue}})

		return
	}

}

func (fu FrontUserController) Chat(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(403, gin.H{"status": "Fail", "message": "异常错误,请稍后再试"})
			return
		}
	}()

	var chatReq req.ChatReq
	err := c.ShouldBindJSON(&chatReq)
	if err != nil {
		c.JSON(403, gin.H{"status": "Fail", "message": "异常错误,请稍后再试"})
		return
	}

	userId, _ := c.Get("userId")
	//如果chatReq.Model 是string
	var modelId int
	var modelName string
	if _, ok := chatReq.Model.(string); ok {
		if strings.HasPrefix(chatReq.Model.(string), "gpt-4-gizmo") {
			model, _ := config.GTPsSlice.GetRandom()
			modelId = model
			modelName = chatReq.Model.(string)
		} else {
			c.JSON(403, gin.H{"status": "Fail", "message": "当前模型不存在请切换或刷新模型"})
			return
		}
	} else {
		modelId = int(chatReq.Model.(float64))
	}
	//如果chatreq.model 是int
	modelCateAndModels, ok := config.ModelMap.Load(modelId)

	if _, ok := chatReq.Model.(float64); ok {
		modelName = modelCateAndModels.(config.ModelCateAndModels).ModelCate.Model

	}

	if !ok {
		c.JSON(403, gin.H{"status": "Fail", "message": "当前模型不存在请切换或刷新模型"})
		return

	}

	//获取一个API KEY
	apiKey, ApiAddr := service.ModelKeyService{}.GetModelKeyByWeight(modelCateAndModels.(config.ModelCateAndModels))
	//如果parentMsgID为空 则直接组装messages
	if apiKey == "" {
		c.JSON(403, gin.H{"status": "Fail", "message": "当前模型无可用key"})
		return

	}
	if ApiAddr == "" {
		//调用系统默认地址 暂时不做处理
		ApiAddr = ""
	}

	//判断当前模型是否收费
	var ModelCate = modelCateAndModels.(config.ModelCateAndModels).ModelCate
	if *ModelCate.Bill == 1 {
		userBill := models.UserBillInfo{}
		err := conn.DB.Where("userId = ?", userId).First(&userBill).Error
		if err != nil {
			c.JSON(403, gin.H{"status": "Fail", "message": "获取用户积分失败,请稍后再试"})
			return
		}
		if userBill.Coins-*ModelCate.Coin < 0 {
			c.JSON(403, gin.H{"status": "Fail", "message": "积分不足,请充值"})
			return
		} else {
			//扣除积分
			result := conn.DB.Model(&models.UserBillInfo{}).Where("userId = ? and coins = ? ", userId, userBill.Coins).Updates(map[string]interface{}{
				"coins":      gorm.Expr("coins - ?", *ModelCate.Coin),
				"used_coins": gorm.Expr("used_coins + ?", *ModelCate.Coin),
			})
			if result.Error != nil || result.RowsAffected == 0 {
				c.JSON(403, gin.H{"status": "Fail", "message": "扣除积分失败,请稍后再试"})
				return
			}
		}
	}

	//如果oriMsgID为空 则生成一个随机UUID作为MSGID
	var msgID string
	var messages []map[string]string
	//将userID转换成string
	userIdStr := fmt.Sprintf("%v", userId)
	if chatReq.OriMsgID == "" {
		msgID, _ = tools.GenerateKey()
	} else {
		msgID = chatReq.OriMsgID
	}

	if chatReq.ParentMsgID == "" {
		messages = chatReq.Messages
	} else {
		//在redis中寻找 userid+msg+parentMsgID
		//获取当前调用模型对应的model 以及 history 以及 maxToken

		modelCateAndModels := modelCateAndModels.(config.ModelCateAndModels)
		parentMsgID := chatReq.ParentMsgID
		history := modelCateAndModels.ModelCate.History
		maxToken := modelCateAndModels.ModelCate.MaxToken
		if history == nil {
			history = new(int)
		}
		if maxToken == nil {
			maxToken = new(int)
		}
		//计算chatReq 中 messages 中的content的长度
		contentLen := 0
		var totalToken int
		for _, v := range chatReq.Messages {
			contentLen += len(v["content"])
		}
		for i := 0; ; i++ {
			if *history != 0 && i >= *history {
				break
			}
			value, err := conn.RedisPool.Get(context.Background(), userIdStr+"msg"+parentMsgID).Result()
			if err != nil {
				break
			}
			chatChain := req.ChatChain{}
			err = json.Unmarshal([]byte(value), &chatChain)
			if err != nil {
				break
			}
			messages = append(messages, chatChain.Content...)

			if *maxToken != 0 {
				totalToken += chatChain.Token

				if totalToken > *maxToken {
					// remove the last content and token
					messages = messages[:len(messages)-1]
					break
				}

			}
			if chatChain.ParentMsgID == "" {
				break
			} else {
				parentMsgID = chatChain.ParentMsgID
			}
			// your code here
			//根据parentID 去循环获取redis中的数据 比如获取到一个信息的格式将是{"content":}

		}

		//如果messages为空则直接返回
		if len(messages) == 0 {
			messages = chatReq.Messages
		} else {
			//逆序messages并且加上chatReq.Messages
			for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
				messages[i], messages[j] = messages[j], messages[i]
			}
			messages = append(messages, chatReq.Messages...)
		}

	}
	//如果parentMsgID不为空 则根据parentMsgID查询到对应的MSGID

	bearerToken := "Bearer " + apiKey
	url := ApiAddr + "/v1/chat/completions"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model":    modelName,
		"messages": messages,
		"stream":   true,
	})

	requests, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	requests.Header.Set("Authorization", bearerToken)
	requests.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(requests)
	if err != nil {

		c.JSON(403, gin.H{"status": "Fail", "message": "1异常错误,请稍后再试"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		c.JSON(500, gin.H{"status": "Fail", "message": "2异常错误,请稍后再试"})
		return
	}
	scanner := bufio.NewScanner(resp.Body)
	var resContent string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			jsonData := line[6:]
			var objmap req.ChatMessage
			if err := json.Unmarshal([]byte(jsonData), &objmap); err == nil {
				objmap.MsgID = msgID

				newJson, _ := json.Marshal(objmap)
				line = "data: " + string(newJson)
				resContent += objmap.Choices[0].Delta.Content
				//反序列choice的内容到 choices中

			}
		}
		fmt.Fprintln(c.Writer, line)
	}
	//将内容转换成 "content":[{role: "user", content: chatreq.message[0]},{role: "assistant", content: resContent}],"token":resContent,parentMsgID的长度 的形式存放在 redis中
	//计算resContent的长度

	var chatChain req.ChatChain
	chatChain.Content = []map[string]string{{"role": "user", "content": chatReq.Messages[0]["content"]}, {"role": "assistant", "content": resContent}}
	chatChain.Token = len(resContent)
	chatChain.ParentMsgID = chatReq.ParentMsgID
	//将chatChain存放到redis中
	chatChainJson, _ := json.Marshal(chatChain)
	conn.RedisPool.Set(context.Background(), userIdStr+"msg"+msgID, chatChainJson, 0)
	if err := scanner.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Fail", "message": "3异常错误,请稍后再试"})
	}
}

func (fu FrontUserController) PreSign(c *gin.Context) {
	//判断R2是否启动
	var req req.PreSign
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "参数错误"})
		return
	}
	key, err := tools.GenerateKey()
	key = req.Filename + key
	if err != nil {
		c.JSON(200, gin.H{"status": "Fail", "message": "生成签名失败"})
		return
	}
	r2Bucket := config.GetConfigMap("R2_bucket").(*config.R2Bucket)

	if r2Bucket.Enable == "1" {

		url, ok := service.R2Service{}.PreSign(key+req.Filename, r2Bucket)
		if !ok {
			c.JSON(200, gin.H{"status": "Fail", "message": "生成签名失败"})
			return
		} else {
			c.JSON(200, gin.H{"status": "Success", "data": gin.H{"up": url, "url": r2Bucket.EndPoint + "/" + key}})
			return
		}
	}
	aliBucket := config.GetConfigMap("Ali_bucket").(*config.AliBucket)
	if aliBucket.Enable == "1" {
		url, ok := service.AliBucketService{}.PreSign(key+req.Filename, aliBucket)
		if !ok {
			c.JSON(200, gin.H{"status": "Fail", "message": "生成签名失败"})
			return
		} else {
			objUrl := "https://" + aliBucket.Bucket + "." + strings.TrimPrefix(aliBucket.EndPoint, "https://") + "/" + key
			c.JSON(200, gin.H{"status": "Success", "data": gin.H{"up": url, "url": objUrl}})
			return
		}
	}

	tencentBucket := config.GetConfigMap("Tencent_bucket").(*config.TencentBucket)
	if tencentBucket.Enable == "1" {
		url, ok := service.TencentBucketService{}.PreSign(key+req.Filename, tencentBucket)
		if !ok {
			c.JSON(200, gin.H{"status": "Fail", "message": "生成签名失败"})
			return
		} else {
			c.JSON(200, gin.H{"status": "Success", "data": gin.H{"up": url, "url": tencentBucket.EndPoint + "/" + key}})
			return
		}
	}
}
