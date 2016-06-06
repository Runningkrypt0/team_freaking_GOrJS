package main

import(
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

