package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// 全局变量
var (
	db  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

// 配置结构体
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Charset      string `mapstructure:"charset"`
	Loc          string `mapstructure:"loc"`
	Timeout      string `mapstructure:"timeout"`
	ReadTimeout  string `mapstructure:"readTimeout"`
	WriteTimeout string `mapstructure:"writeTimeout"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxLifetime  string `mapstructure:"maxLifetime"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	DialTimeout  string `mapstructure:"dialTimeout"`
	ReadTimeout  string `mapstructure:"readTimeout"`
	WriteTimeout string `mapstructure:"writeTimeout"`
}

// 全局配置变量
var Cfg *Config

// InitConfig 使用 Viper 加载配置文件
func InitConfig() error {
	viper.SetConfigFile("config/config.yaml")
	viper.AutomaticEnv() // 支持环境变量覆盖

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置失败: %w", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}
	log.Println("config:", *Cfg)
	return nil
}

// initDB 初始化 MySQL
func initDB() error {
	cfg := Cfg.Mysql

	// 构建 DSN
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&timeout=%s&readTimeout=%s&writeTimeout=%s"
	dsn = fmt.Sprintf(
		dsn,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
		cfg.Loc,
		cfg.Timeout,
		cfg.ReadTimeout,
		cfg.WriteTimeout,
	)

	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	setPool(db)
	fmt.Println("✅ 数据库连接成功")
	return nil
}

// setPool 设置连接池
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("获取数据库连接池失败:", err)
		return
	}

	sqlDB.SetMaxIdleConns(Cfg.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Cfg.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(toDuration(Cfg.Mysql.MaxLifetime))
}

// initRedis 初始化 Redis
func initRedis() error {
	cfg := Cfg.Redis

	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  toDuration(cfg.DialTimeout),
		ReadTimeout:  toDuration(cfg.ReadTimeout),
		WriteTimeout: toDuration(cfg.WriteTimeout),
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Redis 连接失败: %w", err)
	}

	fmt.Println("✅ Redis 连接成功")
	return nil
}

// toDuration 辅助函数：字符串转 time.Duration
func toDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Printf("无效的时间格式 %s，使用默认值 5s", s)
		return 5 * time.Second
	}
	return d
}

// 全局初始化函数（在 main 中调用）
func InitDatabase() error {
	if err := InitConfig(); err != nil {
		return err
	}

	if err := initDB(); err != nil {
		return err
	}

	if err := initRedis(); err != nil {
		return err
	}

	return nil
}

// 导出函数（供其他包使用）
func GetDB() *gorm.DB {
	return db
}

func GetRDB() *redis.Client {
	return rdb
}

func GetContext() context.Context {
	return ctx
}

func GetServerPort() string {
	return ":" + Cfg.Server.Port
}
