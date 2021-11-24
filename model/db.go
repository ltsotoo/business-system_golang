package model

import (
	"business-system_golang/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.SystemConfig.Db.Username,
		config.SystemConfig.Db.Password,
		config.SystemConfig.Db.Host,
		config.SystemConfig.Db.Port,
		config.SystemConfig.Db.Name,
	)

	fmt.Println(dsn)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		fmt.Println("Database connection failed,please check arguments:", err)
	}

	db.AutoMigrate(
		&Contract{},
		&ContractPushMoney{},
		&Customer{},
		&CustomerCompany{},
		&Employee{},
		&Product{},
		&ProductType{},
		&Supplier{},
		&Task{},
		&TaskRemarks{},
		&Payment{},
		&Expense{},
		&DictionaryType{},
		&Dictionary{},
		&PreResearch{},
		&PreResearchTask{},

		&Office{},
		&Area{},
		&Department{},
		&Role{},
		&Permission{},
		&Url{},
	)

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
