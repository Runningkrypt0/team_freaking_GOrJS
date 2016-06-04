package main

import(
	"math"
)

type Stack_Element struct{
	Value interface{}
	Support *Stack_Element
}

type Stack struct{
	Top *Stack_Element
	Bottom *Stack_Element
	Length int
}

func (me *Stack)Pop() (Top interface{}){
	if(me.Length>0){
		me.Length--
		Top, me.Top = me.Top.Value, me.Top.Support
		if(me.Top==nil){
			me.Length=0;
		}
		return
	}
	return nil
}

func (me *Stack)Push(Top interface{}){
	me.Top = &Stack_Element{Top,me.Top}
	if(me.Length<1){
		me.Bottom = me.Top
	}
	me.Length++
}

func (me *Stack)Append(Top interface{}){
	if(me.Length>0){
		me.Bottom.Support = &Stack_Element{Top, nil}
		me.Bottom = me.Bottom.Support
	}else{
		me.Top = &Stack_Element{Top, nil}
		me.Bottom = me.Top
	}
	me.Length++
}


//testing

func inset(path []Vector3, width float64)([]Vector3){
	//create inset area
	
	splits := make([]Vector3, len(path))
	inset_path := make([]Vector3, len(path))
	
	splits[0] = path[0].Clone()
	splits[0].Sub(&path[len(path)-1])
	splits[0].Normalize()
	
	for i := 1; i<len(path); i++{
		splits[i] = path[i].Clone()
		splits[i].Sub(&path[i-1])
		splits[i].Normalize()
	}
	
	//now to create the decension angles

	for i:=0; i<len(path)-1; i++{
		//going toward the center is not the bisector
		
		raw_attractor := splits[i].Clone()
		raw_attractor.Sub(&splits[i+1])
		raw_attractor.Normalize()
		
		back_angle := math.Acos(raw_attractor.Dot(&splits[i]))//angle between
		raw_attractor.Divide(math.Sin(back_angle))
		raw_attractor.Scale(width)
		
		test_angle := Vector3{}
		test_angle.X = -splits[i].Y
		test_angle.Y = splits[i].X
		
		inset_path[i] = path[i].Clone()
		if(test_angle.Dot(&splits[i+1])<0){//convex
			inset_path[i].Add(&raw_attractor)
		}else{ //concave
			inset_path[i].Sub(&raw_attractor)
		}
	}
	//-1 to 0
	
	raw_attractor := splits[len(splits)-1].Clone()
	raw_attractor.Sub(&splits[0])
	raw_attractor.Normalize()
	
	back_angle := math.Acos(raw_attractor.Dot(&splits[len(splits)-1]))//angle between
	raw_attractor.Divide(math.Sin(back_angle))
	raw_attractor.Scale(width)
	
	test_angle := Vector3{}
	test_angle.X = -splits[len(splits)-1].Y
	test_angle.Y = splits[len(splits)-1].X
	
	inset_path[len(splits)-1] = path[len(splits)-1].Clone()
	if(test_angle.Dot(&splits[0])<0){//convex
		inset_path[len(splits)-1].Add(&raw_attractor)
	}else{ //concave
		inset_path[len(splits)-1].Sub(&raw_attractor)
	}
	
	return inset_path;
}