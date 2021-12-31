package calculator

import (
	"fmt"
)

func calcFatRate(weight float64, tall float64, age int, sex string) (fatRate float64, err error) {
	// 计算体脂率
	bmi, err := CalcBMI(tall, weight)
	if err != nil {
		return 0, err
	}

	fatRate ,err= CalcFatRate(bmi, age, sex)
	fmt.Println("体脂率是：", fatRate)
	if 	err != nil {
		return 0, err
	}

	return fatRate ,err
}

