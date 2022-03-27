module github.com/linliuliu/GeekBang-GoCamp

go 1.16

//github.com/armstrongli/go-bmi => D:\GoProject\src\github.com\armstrongli\go-bmi
replace learn.go => D:\GoProject\src\github.com\armstrongli\learn.go

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.2
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
	learn.go v0.0.0-00010101000000-000000000000
)

replace (
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
	github.com/spf13/cobra => github.com/armstrongli/cobra v0.0.0-20211214182251-178edbb247f3 // 原来是 master (branch name)，tidy之后转为精确定位
)
