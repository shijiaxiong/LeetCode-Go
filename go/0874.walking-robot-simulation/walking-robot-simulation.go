package main

func robotSim(commands []int, obstacles [][]int) int {
	obstacles_m:=make(map[[2]int]bool)
	for _,v:=range obstacles{
		if len(v)==2{
			obstacles_m[[2]int{v[0],v[1]}]=true
		}
	}
	//direct* = [4]int{north ,east, south, west}
	directForY := [4]int{1,0,-1,0}
	//y轴坐标四个方向的移动如何处理
	directForX := [4]int{0,1,0,-1}
	//x轴坐标四个方向的移动如何处理
	x:=0 //start x
	y:=0 //start y
	direct:=0 // 方向索引
	var distance int
	for len(commands)!=0{
		command:=commands[0]
		commands=commands[1:]
		if command==-1{
			distance=max(distance,x*x+y*y)
			direct=(direct+1)%4
			continue
		}
		if command==-2{
			distance=max(distance,x*x+y*y)
			direct=(direct+3)%4
			continue
		}

		for i := 0; i < command; i++ {
			tmpx := x + directForX[direct]*1
			tmpy := y + directForY[direct]*1
			if _, ok := obstacles_m[[2]int{tmpx, tmpy}]; ok {
				break
			}
			x = tmpx
			y = tmpy
		}
	}
	return max(distance,x*x+y*y)
}
func max(a,b int )int {
	if a>b{
		return a
	}
	return b
}
