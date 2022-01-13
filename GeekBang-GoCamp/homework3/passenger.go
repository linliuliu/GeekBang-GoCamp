package Elevator

type passenger struct {
	name string     //名字
	target_floot int //想去的楼层
	current_floot int //当前所在楼层
}


//如果没人按，则新增
func (p *passenger)RequstEle(elevator *elevator) *elevator{
	//如果为空则，添加第一个目标楼层
	if len(elevator.floots) ==0 {
		//elevator.floots[0] =p.target_floot
		elevator.floots=append(elevator.floots,p.target_floot)
	}else{
		//不为空，则需要判断目标楼层是否已存在，如果存在则count 加不满,则返回原值即可，不存在则追加
		count:=0
		for _,v := range elevator.floots {
			if v == p.target_floot{
				continue
			}else{
				count++
			}
		}
		if count != len(elevator.floots){
			return elevator
		}else {
			elevator.floots =append(elevator.floots,p.target_floot)
			return elevator
		}
	}

	return elevator
}

//乘客进电梯
func (p *passenger)InEle(elevator *elevator){
	elevator.OpenDoor()
	elevator.passengers++
}

//乘客出电梯
func (p *passenger)OutEle(elevator *elevator){
	elevator.OpenDoor()
	elevator.passengers--
}


