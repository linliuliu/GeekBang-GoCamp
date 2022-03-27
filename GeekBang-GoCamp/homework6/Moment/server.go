
package main

import (

"fmt"
"log"
	"time"

	"gorm.io/driver/mysql"
"gorm.io/gorm"

)

func ConnectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:learngo@tcp(127.0.0.1:3306)/learngo"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func CreateMoment(conn *gorm.DB) error{
	resp := conn.Create(&Moment{
		id:   1,
		User_Nname :   "浦西大怨种",
		User_Age:    28,
		User_Weight: 71,
		User_FatRate:    0.302,
		MomentComment: "今天小区没封，好开心",
		Createtime: time.Now(),
		MomentStatus: 0,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建失败：", err)
		return err
	}
	fmt.Println("创建成功")
	return nil
}
func DeleteMoment(conn *gorm.DB) error{
	resp := conn.Updates(&Moment{
		id:   1,
		User_Nname :   "浦西大怨种",
		User_Age:    28,
		User_Weight: 71,
		User_FatRate:    0.302,
		MomentComment: "今天小区没封，好开心",
		MomentStatus: 1,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新时失败：", err)
		return err
	}
	fmt.Println("更新成功")
	return nil
}
func GetMoment(conn *gorm.DB) ([] *Moment,error){
	output := make([]*Moment, 0)
	resp := conn.Where("momentStatus > ?", 0).Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return nil,resp.Error
	}

	return output,nil
}





