package calculator

import (
	"fmt"
)

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 || tall >= 3.5{
		return 0, fmt.Errorf("身高不能是0或者负数也不会超过3.5米")
	}
	// todo 验证体重的合法性
	if weight <= 0 || weight>= 600{
		return 0, fmt.Errorf("体重不能是0或者负数也不会超过600kg")
	}
	return weight / (tall * tall), nil
}
