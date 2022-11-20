package svc

import (
	"cloudclass_go/internal/config"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  *redis.Client
}

type Userinf struct {
	Uuid     int64  `gorm:"primaryKey"`
	Class    int    `gorm:"index:idx_class"`
	Cid      int    `gorm:"index:idx_cid"`
	Name     string `gorm:"size:10"`
	Username string `gorm:"size:20"`
	Passwd   string `gorm:"size:20"`
	Manage   string `gorm:"size:10"`
}
type Subject struct {
	Uuid    int64  `gorm:"primaryKey"`
	Class   int    `gorm:"index:idx_class"`
	Weekly  int    `gorm:"index:idx_weekly"`
	Sbjname string `gorm:"size:10"`
	Th      int    `gorm:"index:idx_th"`
	Src     string `gorm:"size:100"`
}
type SeatGroup struct {
	Uuid   int64  `gorm:"primaryKey"`
	Class  int    `gorm:"index:idx_class"`
	Group  int    `gorm:"index:idx_group"`
	Gtype  string `gorm:"size:2"`
	Seat   string `gorm:"size:10"`
	Gseat  int
	People string `gorm:"size:10"`
	Status string `gorm:"size:10"`
}
type Message struct {
	Uuid  int64 `gorm:"primaryKey"`
	Class int
	Type  string `gorm:"size:10"`
	Msg   string `gorm:"size:200"`
	time.Time
}
type Log struct {
	Uuid int64  `gorm:"primaryKey"`
	Type string `gorm:"size:10;index:idx_type"`
	Logs string `gorm:"size:250"`
	time.Time
}

func NewServiceContext(c config.Config) *ServiceContext {

	dsn := "cloudclass:r6PT86Yyj4jFcZ8i.@tcp(119.23.69.180:3306)/cloudclass?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	Mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	sqlDB, err2 := Mysql.DB()
	if err != nil {
		println(err.Error())
		fmt.Printf(err.Error())
	}
	if err2 != nil {
		println(err2)
		fmt.Printf(err2.Error())
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(95)
	sqlDB.SetConnMaxLifetime(30)
	//Mysql.AutoMigrate(&Userinf{})
	//Mysql.AutoMigrate(&Subject{})
	//Mysql.AutoMigrate(&SeatGroup{})
	//Mysql.AutoMigrate(&Message{})
	//Mysql.AutoMigrate(&Log{})
	Redis := redis.NewClient(&redis.Options{
		Addr:     "119.23.69.180:6379",
		Password: "Qxr7g9hh386.",
		DB:       0,
	})
	_, rerr := Redis.Ping().Result()
	if rerr != nil {
		print(rerr.Error())
	}

	return &ServiceContext{
		Config: c,
		Mysql:  Mysql,
		Redis:  Redis,
	}
}
