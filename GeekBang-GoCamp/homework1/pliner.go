package main

import "fmt"

type point struct {
	x ,y float64
}

type liner struct {
	p1 point
	p2 point
}

func NewLiner() liner{
	var p1 point
	var p2 point
	fmt.Println("请输入直线")
	fmt.Scanln(&p1.x, &p1.y, &p2.x, &p2.y)
	return liner{p1,p2}
}

func (l *liner)BoolLiner() liner{
	for {
		if l.p1.x == l.p2.x && l.p1.y ==l.p2.y{
			fmt.Println("点重合了，请重输")
			*l =NewLiner()
		}else {
			break
		}
	}

	return *l
}


func (l *liner)LinerRake() float64{
	var rake float64
	rake = (l.p2.y -l.p1.y)/(l.p2.x-l.p1.x)
	return rake
}

func main2(){
	fmt.Println("请输入直线1")
	var l1 =NewLiner()
	l1.BoolLiner()
	fmt.Println("请输入直线2")
	var l2 =NewLiner()
	l2.BoolLiner()
	if l1.LinerRake() ==l2.LinerRake(){
		if (l1.p1.y-(l1.p1.x*l1.LinerRake())) ==(l2.p1.y-(l2.p1.x*l2.LinerRake())){
			fmt.Println("两条直线重合")
		}else {
			fmt.Println("两条直线平行")
		}
	}else{
		fmt.Println("两条直线不平行")
	}
}

