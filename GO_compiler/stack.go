package main

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


/*testing

func inset(path []Vector3){
	//create inset area
	
	var splits = []Vector3;
	
	splits.push(this.border.geometry.vertices[0].clone().sub(this.border.geometry.vertices[this.border.geometry.vertices.length-2]).normalize());
	for(var i=1;i<this.border.geometry.vertices.length;i++){
		splits.push(this.border.geometry.vertices[i].clone().sub(this.border.geometry.vertices[i-1]).normalize());
	}
	
	//now to create the decension angles
	var raw_attractor;
	var offset;

	var geometry = this.border.geometry.clone();
	for(var i=0;i<this.border.geometry.vertices.length-1;i++){
		//going toward the center is not the bisector
		raw_attractor = splits[i].clone().sub(splits[i+1]);
		raw_attractor.normalize();
		
		var back_angle = Math.acos(raw_attractor.dot(splits[i]));//angle between
		raw_attractor.divideScalar(Math.sin(back_angle));
		raw_attractor.multiplyScalar(this.width);
		var test_angle = new THREE.Vector3();
		test_angle.x = -splits[i].z;
		test_angle.z = splits[i].x;
		if(test_angle.dot(splits[i+1])<0){//convex
			geometry.vertices[i].sub(raw_attractor);
		}else{ //concave
			geometry.vertices[i].add(raw_attractor);
		}
	}
	geometry.vertices[geometry.vertices.length-1].copy(geometry.vertices[0]);
	
	this.object = new THREE.Line(geometry);
}

*/