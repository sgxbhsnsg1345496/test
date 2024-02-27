package main

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"mxshop/Goods/model"
	"os"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, err := io.WriteString(Md5, code)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(Md5.Sum(nil))

}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	dsn := "root:root@tcp(192.168.0.101:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.Category{}, &model.Brands{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})

	// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	//options := &password.Options{10, 50, 50, sha512.New}
	//salt, encodedPwd := password.Encode("xxss", options)
	//newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//passwordInfo := strings.Split(newpassword, "$")
	//check := password.Verify("xxs", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check) // true
}
