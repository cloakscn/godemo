package main

// import (
// 	"crypto/sha512"
// 	"fmt"
// 	"strings"

// 	"github.com/anaskhan96/go-password-encoder"
// )

// import (
// 	"crypto"
// 	"crypto/md5"
// 	"encoding/hex"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"time"

// 	"example.com/test/ms_srv/user_srv/model"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// 	"gorm.io/gorm/schema"
// )

// func genMD5(code string) string {
// 	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
// 	s, s2 := password.Encode(code, options)
// 	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", s, s2)

// }

// func verify(code string, encode string) bool {
// 	s := strings.Split(encode, "$")
// 	return password.Verify(code, s[2], s[3], &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New})
// }

func main() {
	// dsn := "root:Ycu061036@tcp(docker.cloaks.cn:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,
	// 		LogLevel: logger.Info,
	// 		Colorful: true,
	// 	},
	// )

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix: "tb_",
	// 		SingularTable: true,
	// 	},
	// 	Logger: newLogger,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.AutoMigrate(&model.User{})
	// if err != nil {
	// 	panic(err)
	// }

	// passwd := genMD5("123456")
	// fmt.Println(len(passwd))
	// fmt.Println(verify("123456", passwd))
}
