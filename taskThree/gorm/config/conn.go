package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB
var dsn = "root:root@1234@tcp(127.0.0.1:3306)/sqlx?charset=utf8mb4&parseTime=True&loc=Local" //loc=Asia%2FShanghai
func init() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:        "mysql",
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info), //打印sql语句
		PrepareStmt: true,                                //启用全局预编译，提高后续查询效率，但不支持嵌套事务
		//SkipDefaultTransaction: true, //禁用事务
		//DryRun:      true,// 不执行sql，只把sql打印出来
		//DisableNestedTransaction: true, //禁用嵌套事务
		//FullSaveAssociations: true, //支持修改关联数据，即在有外键约束的情况下，修改被约束表时，可以将约束表的字段进行修改
		//AllowGlobalUpdate: true, //支持全局更新，即update table set xx=yy，将所有记录的某字段都改成某值，不推荐，可以使用update table set xx=yy where 1=1伪全局
	})
	if err != nil {
		log.Println(err)
		return
	}
	setPool(db)
	fmt.Println("数据库初始化成功")
}
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(5)
	// 最大连接数
	sqlDB.SetMaxOpenConns(10)
	// 最长存活时间
	sqlDB.SetConnMaxLifetime(time.Second * 20)
}

func GetDB() *gorm.DB {
	return db
}
