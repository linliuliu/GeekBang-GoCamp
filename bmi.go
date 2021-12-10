package main

import "fmt"

//计算多个人的平均体脂
//实现完整的体脂计算器
//连续输入三人的姓名、性别、身高、体重、年龄信息
//计算每个人的BMI、体脂率
//输出：
//每个人的姓名、BMI、体脂率、建议
//总人数、平均体脂率

//判断两条线是否平行
//提示：
//两点决定一条直线
//两条线是否平行取决于两条线的斜率是否一样

type personinfo struct {
	name   string
	sex    string
	tall   float64
	weigth float64
	age    int
}

func BoolInfo() personinfo {
	var p personinfo;
	fmt.Println("请依次输入姓名，性别（男/女），身高（m),体重（kg),年龄，空格区分。")
	fmt.Scanln(&p.name, &p.sex, &p.tall, &p.weigth, &p.age)
	if p.sex != "男" || p.sex != "女" { //这里好像有点问题。需要输入两次
		fmt.Println("重来吧你，只能性别男或女")
		fmt.Scanln(&p.name, &p.sex, &p.tall, &p.weigth, &p.age)
	}
	if p.tall > 3.5 || p.tall <= 0 {
		fmt.Println("重来吧你，身高有问题呀")
		fmt.Scanln(&p.name, &p.sex, &p.tall, &p.weigth, &p.age)
	}
	if p.weigth <= 0 || p.weigth >= 600 {
		fmt.Println("重来吧你，体重有问题呀")
		fmt.Scanln(&p.name, &p.sex, &p.tall, &p.weigth, &p.age)
	}
	if p.age <= 18 || p.age > 160 {
		fmt.Println("重来吧你，小于十八岁不允许参加")
		fmt.Scanln(&p.name, &p.sex, &p.tall, &p.weigth, &p.age)
	}

	fmt.Println("确认下，姓名：", p.name, "，性别：", p.sex, "，身高：", p.tall, "m，体重：", p.weigth, "kg,年龄：", p.age)
	return p
}

func (p *personinfo)BodyFatRatio()  (float64,float64){
	//
      var bmi float64
      bmi = p.weigth /(p.tall*p.tall)
	//体脂率：1.2×BMI+0.23×年龄-5.4-10.8×性别（男为1，女为0）
      var bfr float64
      if p.sex == "男"{
		  bfr = ((1.2*bmi) +(0.23*float64(p.age))- 5.4-10.8)/100
		  //fmt.Println("体脂率为：",bfr)
	  }else{
		  bfr = ((1.2*bmi) +(0.23*float64(p.age))- 5.4)/100
		  //fmt.Println("体脂率为：",bfr)
	  }
      return bmi,bfr
}

func (p *personinfo)BodyFatRange(){
	bmi, bfr := p.BodyFatRatio()
	if p.sex =="男"{
		if p.age <=39{
			if bfr <=0.1{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏瘦，多吃点" )
			}else if bfr >0.1 && bfr <=0.16{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"标准，继续保持" )
			}else if bfr >0.17 && bfr <=0.21{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏旁，适当动动" )
			}else if bfr >0.21 && bfr <=0.26{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"肥胖，需要定期运动" )
			}else {
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"严重肥胖，天天运动，控制饮食甚至需要看医生" )
			}
		}else if p.age>39 && p.age <=59{
			if bfr <=0.11{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏瘦，多吃点" )
			}else if bfr >0.11 && bfr <=0.17{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"标准，继续保持" )
			}else if bfr >0.18 && bfr <=0.22{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏旁，适当动动" )
			}else if bfr >0.22 && bfr <=0.27{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"肥胖，需要定期运动" )
			}else {
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"严重肥胖，天天运动，控制饮食甚至需要看医生" )
			}
		}else{
			if bfr <=0.13{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏瘦，多吃点" )
			}else if bfr >0.13 && bfr <=0.19{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"标准，继续保持" )
			}else if bfr >0.19 && bfr <=0.24{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏旁，适当动动" )
			}else if bfr >0.24 && bfr <=0.29{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"肥胖，需要定期运动" )
			}else {
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"严重肥胖，天天运动，控制饮食甚至需要看医生" )
			}
		}
	}else{
		if p.age <=39{
			if bfr <=0.2{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏瘦，多吃点" )
			}else if bfr >0.2 && bfr <=0.27{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"标准，继续保持" )
			}else if bfr >0.27 && bfr <=0.34{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏旁，适当动动" )
			}else if bfr >0.34 && bfr <=0.39{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"肥胖，需要定期运动" )
			}else {
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"严重肥胖，天天运动，控制饮食甚至需要看医生" )
			}
		}else if p.age>39 && p.age <=59{
			if bfr <=0.21{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏瘦，多吃点" )
			}else if bfr >0.21 && bfr <=0.28{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"标准，继续保持" )
			}else if bfr >0.28 && bfr <=0.35{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏旁，适当动动" )
			}else if bfr >0.35 && bfr <=0.4{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"肥胖，需要定期运动" )
			}else {
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"严重肥胖，天天运动，控制饮食甚至需要看医生" )
			}
		}else{
			if bfr <=0.22{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏瘦，多吃点" )
			}else if bfr >0.22 && bfr <=0.29{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"标准，继续保持" )
			}else if bfr >0.29 && bfr <=0.36{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"偏旁，适当动动" )
			}else if bfr >0.36 && bfr <=0.41{
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"肥胖，需要定期运动" )
			}else {
				fmt.Println(p.name,"的BMI是",bmi,"，体脂率是",bfr,"严重肥胖，天天运动，控制饮食甚至需要看医生" )
			}
		}
	}
}


func main() {

	//fmt.Println("请依次输入姓名，性别（男/女），身高（m),体重（kg),年龄，空格区分。")
	//fmt.Scanln (&personinfo.name,&personinfo.sex,&personinfo.tall,&personinfo.weigth,&personinfo.age)
	//if personinfo.sex != "男" || personinfo.sex != "女"{
	//    fmt.Println("重来吧你，只能男或女")
	//	fmt.Scanln (&personinfo.name,&personinfo.sex,&personinfo.tall,&personinfo.weigth,&personinfo.age)
	//}
	//if personinfo.tall >3.5 || personinfo.tall <=0  {
	//	fmt.Println("重来吧你，身高有问题呀")
	//	fmt.Scanln (&personinfo.name,&personinfo.sex,&personinfo.tall,&personinfo.weigth,&personinfo.age)
	//}
	//if personinfo.weigth <=0 || personinfo.weigth >= 600{
	//	fmt.Println("重来吧你，体重有问题呀")
	//	fmt.Scanln (&personinfo.name,&personinfo.sex,&personinfo.tall,&personinfo.weigth,&personinfo.age)
	//}
	//if personinfo.age <= 18 || personinfo.age >160{
	//	fmt.Println("重来吧你，小于十八岁不允许参加")
	//	fmt.Scanln (&personinfo.name,&personinfo.sex,&personinfo.tall,&personinfo.weigth,&personinfo.age)
	//}
	//fmt.Println("确认下，姓名：",personinfo.name,"，性别：",personinfo.sex,"，身高：",personinfo.tall,"m，体重：",personinfo.weigth,"kg,年龄：",personinfo.age)

	//BoolInfo()
	var person [3]personinfo
	person[0] = BoolInfo()
	person[1] = BoolInfo()
	person[2] = BoolInfo()

	person[0].BodyFatRange()
	person[1].BodyFatRange()
	person[2].BodyFatRange()

	_,bfr1 :=person[0].BodyFatRatio()
	_,bfr2 :=person[1].BodyFatRatio()
	_,bfr3 :=person[2].BodyFatRatio()

	fmt.Println("总人数",len(person),"平均体脂率：",(bfr1+bfr2+bfr3)/float64(len(person)))
}
