package calculator

import "testing"

func TestCalcFatRate(t *testing.T) {
	inputBmi, inputAge,inputSex := 1.0, 28,"男"
	expectedOutput := 1.0
	t.Logf("开始计算，输入：bmi: %f, 年龄: %d, 性别：%s,期望结果：%f", inputBmi, inputAge,inputSex, expectedOutput)
	actualOutput, err := CalcFatRate(inputBmi, inputAge,inputSex)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

func TestCalcFatRate1(t *testing.T) {
	inputBmi, inputAge,inputSex := 0.0, 28,"男"
	expectedOutput := 0.0
	t.Logf("开始计算，输入：bmi: %f, 年龄: %d, 性别：%s,期望结果：%f", inputBmi, inputAge,inputSex, expectedOutput)
	actualOutput, err := CalcFatRate(inputBmi, inputAge,inputSex)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

func TestCalcFatRate2(t *testing.T) {
	inputBmi, inputAge,inputSex := 2.0, 188,"男"
	expectedOutput := 0.0
	t.Logf("开始计算，输入：bmi: %f, 年龄: %d, 性别：%s,期望结果：%f", inputBmi, inputAge,inputSex, expectedOutput)
	actualOutput, err := CalcFatRate(inputBmi, inputAge,inputSex)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

func TestCalcFatRate3(t *testing.T) {
	inputBmi, inputAge,inputSex := 0.0, 28,"修勾"
	expectedOutput := 0.0
	t.Logf("开始计算，输入：bmi: %f, 年龄: %d, 性别：%s,期望结果：%f", inputBmi, inputAge,inputSex, expectedOutput)
	actualOutput, err := CalcFatRate(inputBmi, inputAge,inputSex)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}