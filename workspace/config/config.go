package config

import (
	"conn"
	"encoding/json"
	"fmt"
	"math/rand"
	"models"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

type Config struct {
	Database Database
	Redis    Redis
	Jwt      Jwt
}

type Database struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DBname   string `ini:"dbname"`
}

type Redis struct {
	Host           string `ini:"host"`
	Port           string `ini:"port"`
	Balance        int    `ini:"balance"`
	Balancetimeout int    `ini:"balance_timeout"`
}

type Jwt struct {
	Secret string `ini:"secret"`
}

type Email struct {
	EmailServer string `json:"emailServer"`
	EmailPort   string `json:"emailPort"`
	EmailUser   string `json:"emailUser"`
	EmailPass   string `json:"emailPass"`
	EmailTLS    string `json:"emailTLS"`
}
type Notice struct {
	Notice string `json:"notice"`
}
type EmailReg struct {
	EmailValidation string `json:"emailValidation"`
	EmailFormat     string `json:"emailFormat"`
	EmailSubject    string `json:"emailSubject"`
	EmailFrom       string `json:"emailFrom"`
	Duration        string `json:"duration"`
}

type wxLogin struct {
	Login     string `json:"login"`
	Auto      string `json:"auto"`
	Name      string `json:"name"`
	AppID     string `json:"appID"`
	Token     string `json:"token"`
	AppSecret string `json:"appSecret"`
	Ref1      string `json:"ref1"`
	Ref2      string `json:"ref2"`
	Ref3      string `json:"ref3"`
	Ref4      string `json:"ref4"`
}
type ModelCateAndModels struct {
	ModelCate models.ModelCate
	Models    *sync.Map
}

type R2Bucket struct {
	Enable    string `json:"enable"`
	AccountID string `json:"accountID"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	EndPoint  string `json:"endPoint"`
}

type AliBucket struct {
	Enable    string `json:"enable"`
	EndPoint  string `json:"endPoint"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
}

type TencentBucket struct {
	Enable    string `json:"enable"`
	EndPoint  string `json:"endPoint"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

var KeyWeightIndices sync.Map

var ToolMap sync.Map

// 生成一个通用的模型MAP 存放在内存中

var ModelMap sync.Map

type GPTsSafeSlice struct {
	sync.RWMutex
	items []int
}

func (ss *GPTsSafeSlice) Append(item int) {
	ss.Lock()
	defer ss.Unlock()

	ss.items = append(ss.items, item)
}

func (ss *GPTsSafeSlice) GetRandom() (int, bool) {
	ss.RLock()
	defer ss.RUnlock()
	//随机从切片中取出一个ID
	if len(ss.items) == 0 {
		return 0, false
	}
	//随机选取一个ID
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ss.items))
	return ss.items[index], true
}

func (ss *GPTsSafeSlice) Remove(item int) {
	ss.Lock()
	defer ss.Unlock()
	//遍历删除指定值
	for i, v := range ss.items {
		if v == item {
			ss.items = append(ss.items[:i], ss.items[i+1:]...)
			break
		}

	}
}

var GTPsSlice GPTsSafeSlice

// 定义一个切片 存放key切片

var ConfigList = &Config{}

// 定义一个map 存放 结构体 空接口
var ConfigLock = make(map[string]*sync.RWMutex)

var ConfigMap = map[string]interface{}{
	"email":          &Email{},
	"email_reg":      &EmailReg{},
	"wx_login":       &wxLogin{},
	"notice":         &Notice{},
	"R2_bucket":      &R2Bucket{},
	"Ali_bucket":     &AliBucket{},
	"Tencent_bucket": &TencentBucket{},
}

//定义读取config和写入configmap的函数 自动使用锁

func GetConfigMap(listname string) interface{} {
	ConfigLock[listname].RLock()

	defer ConfigLock[listname].RUnlock()
	return ConfigMap[listname]
}

// 写入configmap
func SetConfigMap(listname string, jsonString string) error {
	ConfigLock[listname].Lock()
	defer ConfigLock[listname].Unlock()
	err := json.Unmarshal([]byte(jsonString), ConfigMap[listname])
	return err
}

func init() {
	// Load config from file
	// 读取config.ini配置文件
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("读取配置文件失败: %v ,请确认文件是否在文件夹下", err)
		return
	}
	err = cfg.Section("database").MapTo(&ConfigList.Database)
	if err != nil {
		fmt.Printf("failed to read database config: %v", err)
		ConfigList = nil

		return
	}
	err = cfg.Section("redis").MapTo(&ConfigList.Redis)
	if err != nil {
		fmt.Printf("failed to read redis config: %v", err)
		ConfigList = nil

		return
	}
	err = cfg.Section("jwt").MapTo(&ConfigList.Jwt)

	if err != nil {
		fmt.Printf("failed to read jwt config: %v", err)
		ConfigList = nil

		return
	}
	if ConfigList.Jwt.Secret == "" || ConfigList.Jwt.Secret == "random string" {
		fmt.Println("请在config.ini中配置jwt的secret 否则用户数据可能被仿造")
	}
	err = conn.InitDB(ConfigList.Database.User, ConfigList.Database.Password, ConfigList.Database.Host, ConfigList.Database.Port, ConfigList.Database.DBname)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	err = conn.InitRedis(ConfigList.Redis.Host, ConfigList.Redis.Port)
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}

	//定义一个string 切片 存放需要初始化的配置名
	configName := []string{"email", "email_reg", "wx_login", "notice", "R2_bucket", "Ali_bucket", "Tencent_bucket"}
	//循环切片
	for _, v := range configName {
		code := models.Codelkup{}
		conn.DB.Debug().Select("Value").Where("listname = ?", v).First(&code)
		//反序列化 将code的value转换到 Email结构体中
		err = json.Unmarshal([]byte(code.Value), ConfigMap[v])
		if err != nil {
			fmt.Println(v + "配置初始化失败")

		} else {
			fmt.Println(ConfigMap[v])
		}
		//读取数据库中key对应的value
	}
	//将数据库的模型载入进来
	modelCateList := []models.ModelCate{}
	conn.DB.Find(&modelCateList)
	//按照ID 将模型存入到map中

	for _, v := range modelCateList {

		modelKeys := &sync.Map{}

		var modelKeysSlice []models.ModelKeys

		conn.DB.Where("model_cate_id = ? and enable = 1", v.ID).Find(&modelKeysSlice)

		for _, model := range modelKeysSlice {
			modelKeys.Store(model.ID, model)
		}
		ModelMap.Store(v.ID, ModelCateAndModels{
			ModelCate: v,
			Models:    modelKeys,
		})
		//如果他是gmiz模型 则 存储名字而非ID
		if v.Model == "gpt-4-gizmo" {
			GTPsSlice.Append(v.ID)
		}
		if *v.Tool == 1 {
			ToolMap.Store(v.ID, ModelCateAndModels{
				ModelCate: v,
				Models:    modelKeys,
			})
		}
		weightIndex := make([]int, 0)

		modelKeys.Range(func(key, value interface{}) bool {
			model := value.(models.ModelKeys)

			for i := 0; i < model.Weight; i++ {
				weightIndex = append(weightIndex, model.ID)
			}
			return true
		})

		// 将权重索引存储在sync.Map中
		KeyWeightIndices.Store(v.ID, weightIndex)

	}

	for listname := range ConfigMap {
		ConfigLock[listname] = &sync.RWMutex{}
	}

}
