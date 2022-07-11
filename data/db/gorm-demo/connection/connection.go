package main

import (
	"fmt"
	"github.com/minicloudsky/golang-in-action/util/log/zaplog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// APIRequestRecord 接口访问记录表
type APIRequestRecord struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null;comment:'主键'"`                                                              // 主键
	APIPath    string    `gorm:"index:idx_api_path;index:idx_create_by;column:api_path;type:varchar(200);not null;default:'';comment:'接口英文名'"`                                    // 接口英文名
	Method     string    `gorm:"column:method;type:varchar(100);not null;comment:'请求方法'"`                                                                                         // 请求方法
	Params     string    `gorm:"column:params;type:longtext;not null;comment:'请求参数'"`                                                                                             // 请求参数
	APIName    string    `gorm:"column:api_name;type:varchar(200);not null;default:'';comment:'接口中文名'"`                                                                           // 接口中文名
	CostTime   int       `gorm:"column:cost_time;type:int;not null;comment:'请求耗时'"`                                                                                               // 请求耗时
	IP         string    `gorm:"column:ip;type:varchar(20);not null;comment:'来源ip'"`                                                                                              // 来源ip
	Username   string    `gorm:"column:username;type:varchar(100);not null;comment:'请求用户名'"`                                                                                      // 请求用户名
	StatusCode int       `gorm:"column:status_code;type:int;not null;comment:'状态码'"`                                                                                              // 状态码
	StatusText string    `gorm:"column:status_text;type:longtext;not null;comment:'状态文本'"`                                                                                        // 状态文本
	CreateBy   string    `gorm:"index:idx_api_path;index:idx_create_by;column:create_by;type:varchar(20);not null;default:'';comment:'创建人'"`                                      // 创建人
	UpdateBy   string    `gorm:"column:update_by;type:varchar(100);not null;comment:'更新人'"`                                                                                       // 更新人
	CreateTime time.Time `gorm:"index:idx_api_path;index:idx_create_by;index:idx_create_time;column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                              // 更新时间
}

func NewDB() *gorm.DB {
	logger := zaplog.InitZapLogger()
	dsn := "root:root@tcp(127.0.0.1:3306)/ypp_sql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Warn(fmt.Sprintf("fail to connect mysql: %s", dsn))
	}
	err = db.AutoMigrate(&APIRequestRecord{})
	if err != nil {
		return nil
	}
	return db
}
