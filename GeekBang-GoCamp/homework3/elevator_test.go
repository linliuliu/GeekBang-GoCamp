package Elevator

import "testing"

func Test1(t *testing.T){
	//楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
	ele:=new(elevator)
	ele.floots=[]int{}
	ele.movestatus= "stop"
	ele.currentfloot=3
	ele.passengers=0
	expectedOut := "stop"
	actualout:=ele.GetCurrentStatus()
	if expectedOut != actualout {
		t.Errorf("expecting %s, but got %s", expectedOut, actualout)
	}
}


func Test2(t *testing.T){
	//楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼
	ele:=new(elevator)
	ele.floots=[]int{}
	ele.movestatus= "stop"
	ele.currentfloot=1
	ele.passengers=0
	p :=make([]passenger,2)
	p1 :=new(passenger)
	p1.target_floot=3
	p1.current_floot=3
	p1.name ="A"
	p=append(p,*p1)
	p1.RequstEle(ele)
	//ele.Movepolicy(p)
	expectedOut:=3

	actualout:=ele.GetCurrentFloot()
	if expectedOut != actualout {
		t.Errorf("expecting %d, but got %d", expectedOut, actualout)
	}
}



func Test3(t *testing.T){
	//楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
	ele:=new(elevator)
	ele.floots=[]int{}
	ele.movestatus= "stop"
	ele.currentfloot=3
	ele.passengers=0
	p :=make([]passenger,2)
	p1 :=new(passenger)
	p1.target_floot=3
	p1.current_floot=4
	p1.name ="A"
	p2 :=new(passenger)
	p2.target_floot=3
	p2.current_floot=2
	p2.name ="A"
	p=append(p,*p1)
	p=append(p,*p2)
	p1.RequstEle(ele)
	p2.RequstEle(ele)
	ele.Movepolicy(p)
	expectedOut:=2

	actualout:=ele.GetCurrentFloot()
	if expectedOut != actualout {
		t.Errorf("expecting %d, but got %d", expectedOut, actualout)
	}
}