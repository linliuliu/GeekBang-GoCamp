package calculator

import "fmt"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64,err error) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else if sex =="女"{
		sexWeight = 0
	}
	if bmi <= 0{
		return 0, fmt.Errorf("bmi不能是0或者负数")
	}
	if age <= 18 ||age >= 150{
		return 0, fmt.Errorf("age不对")
	}
	if sex != "男"&& sex != "女"{
		return 0, fmt.Errorf("性别只能为\"男\"或\"女\"")
	}
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return fatRate,nil
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
