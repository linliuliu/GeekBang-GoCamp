package calculator

import (
	"testing"
)

func TestM(t *testing.T) {
	inputHeight, inputWeight , inputAge,inputSex := 1.7,63.0, 28,"修勾"
	expectedOutput := 0.0
	t.Logf("开始计算，输入：身高: %f,体重: %f, 年龄: %d, 性别：%s,期望结果：%f", inputHeight, inputWeight, inputAge,inputSex, expectedOutput)
	actualOutput ,err:= calcFatRate(inputHeight, inputWeight, inputAge,inputSex)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}