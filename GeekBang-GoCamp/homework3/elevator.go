package Elevator

import (
	"fmt"
	"time"
)

type elevator struct {
	floots []int //存放按键楼层
	passengers int //电梯内人数
	currentfloot int //当前楼层
	//movestatus bool //状态
	movestatus string//状态
}
//获取当前电梯楼层
func (ele *elevator)GetCurrentFloot() int{
	return ele.currentfloot
}
//获取当前电梯内人数
func (ele *elevator)GetCurrentPassengers() int{
	return ele.passengers
}
//获取当前电梯状态
func (ele *elevator)GetCurrentStatus() string  {
	return ele.movestatus
}
//电梯移动
func (ele *elevator)move(){
	//ele.movestatus != "stop"
	time.Sleep(1000)
}
//电梯向上移动，楼层+1
func (ele *elevator)Up(){
	ele.move()
	ele.movestatus ="up"
	ele.currentfloot++
}
//电梯向下移动，楼层-1
func (ele *elevator)Down(){
	ele.move()
	ele.movestatus ="down"
	ele.currentfloot--
}
//电梯开门
func (ele *elevator)OpenDoor(){

	if ele.GetCurrentStatus() != "stop"{
		fmt.Errorf("plese waiting")
	}
}
//往下的策略
func(ele *elevator)Downpolicy(p []passenger){
	for{
		ele.Down()
		for i,v:=range p{
			if v.target_floot == ele.GetCurrentFloot(){
				p[i].OutEle(ele)
			}
		}
	}

		for i,v:=range p{
			if v.current_floot<ele.GetCurrentFloot(){
				ele.Down()
				if v.current_floot== ele.GetCurrentFloot(){
					p[i].InEle(ele)
				}
			}

		}


	ele.Down()
			for i,v:=range p{
				if v.current_floot<ele.GetCurrentFloot(){
					ele.Down()
					if v.current_floot== ele.GetCurrentFloot(){
						p[i].InEle(ele)
					}
				}
				//乘客当前楼层和电梯楼层一致，且目标楼层小于当前楼层;进电梯
				if v.current_floot== ele.currentfloot && v.current_floot> v.target_floot{

				}
				if v.target_floot ==ele.currentfloot{
					p[i].OutEle(ele)
				}
			}
			//来到了一楼，跳出循环
			if ele.currentfloot ==1 {
				//break
			}

}


func(ele *elevator)Uppolicy(p []passenger){
	//
		for i:=0;i<len(ele.floots);i++ {
			if
			ele.Up()
			//遍历乘客
			for i, v := range p {
				//乘客当前楼层和电梯楼层一致，且目标楼层大于当前楼层;进电梯
				if v.current_floot == ele.currentfloot && v.current_floot < v.target_floot {
					p[i].InEle(ele)
				}
				if v.target_floot == ele.currentfloot {
					p[i].OutEle(ele)
				}
			}
			//来到了5楼，跳出循环
			if ele.currentfloot == 5 {
				break
			}
		}
}



func(ele *elevator)Movepolicy(p []passenger){
	//如果没有按楼层则
	if len(ele.floots) == 0{
		ele.movestatus = "stop"
		return
	}
	if len(ele.floots) ==1 {
		if p[0].current_floot !=ele.currentfloot{
			if p[0].current_floot> ele.currentfloot{
				for {
					ele.Up()
					if p[0].current_floot == ele.GetCurrentFloot(){
						ele.OpenDoor()
						break
					}
				}
			}else {
				if p[0].current_floot<ele.currentfloot{
					for {
						ele.Down()
						if p[0].current_floot == ele.GetCurrentFloot(){
							ele.OpenDoor()
							break
						}
					}
				}
			}
		}
	}
	//如果有人按楼层时需要判断 是上还是下
	//第一个乘客当前楼层小于电梯楼层或者第一个乘客与电梯楼层相同，目标楼层小于现在楼层
	if p[0].current_floot < ele.currentfloot ||
		(p[0].current_floot == ele.currentfloot&& p[0].current_floot>p[0].target_floot) {
		ele.Downpolicy(p)
	}
	//第一个乘客当前楼层大于电梯楼层或者第一个乘客与电梯楼层相同，目标楼层大于现在楼层
	if p[0].current_floot > ele.currentfloot ||
		(p[0].current_floot == ele.currentfloot&& p[0].current_floot<p[0].target_floot) {
		ele.Uppolicy(p)
	}


}