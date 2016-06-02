function touchingTriangles(a,b,c,d,e,f){

	//var rotateQ = new THREE.Quaternion();
	//var tempPlane = new THREE.Plane().setFromCoplanarPoints(a,b,c);
	//rotateQ.setFromUnitVectors(tempPlane.normal.clone() , new THREE.Vector3(0,0,1));
	
	//force z to 0?
	//raise min to -epsilon
	
	var touching = false;

	var me_Point_A = a.clone();//.applyQuaternion(rotateQ);
	var me_Point_B = b.clone();//.applyQuaternion(rotateQ);
	var me_Point_C = c.clone();//.applyQuaternion(rotateQ);
	
	var me_Line_AB = new THREE.Vector3().subVectors(me_Point_A, me_Point_B);
	var me_Line_BC = new THREE.Vector3().subVectors(me_Point_B, me_Point_C);
	var me_Line_CA = new THREE.Vector3().subVectors(me_Point_C, me_Point_A);
	
	var me_Inside_AB = new THREE.Vector3().subVectors(me_Point_A, me_Point_C);
	var me_Inside_BC = new THREE.Vector3().subVectors(me_Point_B, me_Point_A);
	var me_Inside_CA = new THREE.Vector3().subVectors(me_Point_C, me_Point_B);
	
	var me_True_AB = new THREE.Vector3().crossVectors(me_Line_AB, me_Inside_AB);
	var me_True_BC = new THREE.Vector3().crossVectors(me_Line_BC, me_Inside_BC);
	var me_True_CA = new THREE.Vector3().crossVectors(me_Line_CA, me_Inside_CA);
	
	var corner_check = 0;
	
	//Operator Setup
	var him_Point_A = d.clone();//.applyQuaternion(rotateQ);
	var him_Point_B = e.clone();//.applyQuaternion(rotateQ);
	var him_Point_C = f.clone();//.applyQuaternion(rotateQ);

	var him_Line_AB = new THREE.Vector3().subVectors(him_Point_A, him_Point_B);
	var him_Line_BC = new THREE.Vector3().subVectors(him_Point_B, him_Point_C);
	var him_Line_CA = new THREE.Vector3().subVectors(him_Point_C, him_Point_A);

	var him_Inside_AB = new THREE.Vector3().subVectors(him_Point_A, him_Point_C);
	var him_Inside_BC = new THREE.Vector3().subVectors(him_Point_B, him_Point_A);
	var him_Inside_CA = new THREE.Vector3().subVectors(him_Point_C, him_Point_B);

	var him_True_AB = new THREE.Vector3().crossVectors(him_Line_AB, him_Inside_AB);
	var him_True_BC = new THREE.Vector3().crossVectors(him_Line_BC, him_Inside_BC);
	var him_True_CA = new THREE.Vector3().crossVectors(him_Line_CA, him_Inside_CA);
	
	
	
	//Local Search
	var me_Outside_AA = new THREE.Vector3().subVectors(me_Point_A, him_Point_A);
	var me_Outside_BA = new THREE.Vector3().subVectors(me_Point_B, him_Point_A);
	var me_Outside_CA = new THREE.Vector3().subVectors(me_Point_C, him_Point_A);
	
	var me_Outside_AB = new THREE.Vector3().subVectors(me_Point_A, him_Point_B);
	var me_Outside_BB = new THREE.Vector3().subVectors(me_Point_B, him_Point_B);
	var me_Outside_CB = new THREE.Vector3().subVectors(me_Point_C, him_Point_B);
	
	var me_Outside_AC = new THREE.Vector3().subVectors(me_Point_A, him_Point_C);
	var me_Outside_BC = new THREE.Vector3().subVectors(me_Point_B, him_Point_C);
	var me_Outside_CC = new THREE.Vector3().subVectors(me_Point_C, him_Point_C);
	
	
	var me_Compare_AA = new THREE.Vector3().crossVectors(me_Line_AB, me_Outside_AA);
	var me_Compare_BA = new THREE.Vector3().crossVectors(me_Line_BC, me_Outside_BA);
	var me_Compare_CA = new THREE.Vector3().crossVectors(me_Line_CA, me_Outside_CA);
	
	var me_Compare_AB = new THREE.Vector3().crossVectors(me_Line_AB, me_Outside_AB);
	var me_Compare_BB = new THREE.Vector3().crossVectors(me_Line_BC, me_Outside_BB);
	var me_Compare_CB = new THREE.Vector3().crossVectors(me_Line_CA, me_Outside_CB);
	
	var me_Compare_AC = new THREE.Vector3().crossVectors(me_Line_AB, me_Outside_AC);
	var me_Compare_BC = new THREE.Vector3().crossVectors(me_Line_BC, me_Outside_BC);
	var me_Compare_CC = new THREE.Vector3().crossVectors(me_Line_CA, me_Outside_CC);
	
	//Local Tree
	var me_count=0;
	if(me_True_AB.dot(me_Compare_AA)>=-0.0001&&me_True_BC.dot(me_Compare_BA)>=-0.0001&&me_True_CA.dot(me_Compare_CA)>=-0.0001){
		me_count++;
		corner_check = him_Point_A;
	}
	if(me_True_AB.dot(me_Compare_AB)>=-0.0001&&me_True_BC.dot(me_Compare_BB)>=-0.0001&&me_True_CA.dot(me_Compare_CB)>=-0.0001){
		if(me_count>0){
			return true;
		}
		me_count++;
		corner_check = him_Point_B;
	}
	if(me_True_AB.dot(me_Compare_AC)>=-0.0001&&me_True_BC.dot(me_Compare_BC)>=-0.0001&&me_True_CA.dot(me_Compare_CC)>=-0.0001){
		if(me_count>0){
			return true;
		}
		me_count++;
		corner_check = him_Point_C;
	}
	
	
	
	//Operator Search
	var him_Outside_AA = new THREE.Vector3().subVectors(him_Point_A, me_Point_A);
	var him_Outside_BA = new THREE.Vector3().subVectors(him_Point_B, me_Point_A);
	var him_Outside_CA = new THREE.Vector3().subVectors(him_Point_C, me_Point_A);
	
	var him_Outside_AB = new THREE.Vector3().subVectors(him_Point_A, me_Point_B);
	var him_Outside_BB = new THREE.Vector3().subVectors(him_Point_B, me_Point_B);
	var him_Outside_CB = new THREE.Vector3().subVectors(him_Point_C, me_Point_B);
	
	var him_Outside_AC = new THREE.Vector3().subVectors(him_Point_A, me_Point_C);
	var him_Outside_BC = new THREE.Vector3().subVectors(him_Point_B, me_Point_C);
	var him_Outside_CC = new THREE.Vector3().subVectors(him_Point_C, me_Point_C);
	
	
	var him_Compare_AA = new THREE.Vector3().crossVectors(him_Line_AB, him_Outside_AA);
	var him_Compare_BA = new THREE.Vector3().crossVectors(him_Line_BC, him_Outside_BA);
	var him_Compare_CA = new THREE.Vector3().crossVectors(him_Line_CA, him_Outside_CA);
	
	var him_Compare_AB = new THREE.Vector3().crossVectors(him_Line_AB, him_Outside_AB);
	var him_Compare_BB = new THREE.Vector3().crossVectors(him_Line_BC, him_Outside_BB);
	var him_Compare_CB = new THREE.Vector3().crossVectors(him_Line_CA, him_Outside_CB);
	
	var him_Compare_AC = new THREE.Vector3().crossVectors(him_Line_AB, him_Outside_AC);
	var him_Compare_BC = new THREE.Vector3().crossVectors(him_Line_BC, him_Outside_BC);
	var him_Compare_CC = new THREE.Vector3().crossVectors(him_Line_CA, him_Outside_CC);
	
	//Operator Tree
	var him_count=0;
	if(him_True_AB.dot(him_Compare_AA)>=-0.0001&&him_True_BC.dot(him_Compare_BA)>=-0.0001&&him_True_CA.dot(him_Compare_CA)>=-0.0001){
		if(me_count>0){
			if(!corner_check.equals(me_Point_A)){
				return true;
			}
		}
		him_count++;
	}
	if(him_True_AB.dot(him_Compare_AB)>=-0.0001&&him_True_BC.dot(him_Compare_BB)>=-0.0001&&him_True_CA.dot(him_Compare_CB)>=-0.0001){
		if(me_count>0){
			if(!corner_check.equals(me_Point_B)){
				return true;
			}
		}
		if(him_count>0){
			return true;
		}
		him_count++;
	}
	if(him_True_AB.dot(him_Compare_AC)>=-0.0001&&him_True_BC.dot(him_Compare_BC)>=-0.0001&&him_True_CA.dot(him_Compare_CC)>=-0.0001){
		if(me_count>0){
			if(!corner_check.equals(me_Point_C)){
				return true;
			}
		}
		if(him_count>0){
			return true;
		}
	}
	
	return touching;
}

function bind(a,b){
	if(a>=0){
		return a%b;
	}else{
		return bind(b+a,b);
	}
}

function relativeAngle(set,index,q){
	var front = set[bind(index,set.length)].clone().sub(set[bind(index+1,set.length)]).applyQuaternion(q).normalize();
	var back = set[bind(index,set.length)].clone().sub(set[bind(index-1,set.length)]).applyQuaternion(q).normalize();
	
	var angle = Math.atan2(front.y,front.x) - Math.atan2(back.y,back.x);
	
	if(angle>Math.PI){
		angle-=Math.PI*2;
	}else if(angle<-Math.PI){
		angle+=Math.PI*2;
	}

	return angle;

}

function GenerateColor(){
	return Math.random() * 0xffffff;
}

function SnapToGrid(v){
	v.x = Math.round(v.x/roster.GRID_SIZE)*roster.GRID_SIZE;
	v.y = Math.round(v.y/roster.GRID_SIZE)*roster.GRID_SIZE;
	v.z = Math.round(v.z/roster.GRID_SIZE)*roster.GRID_SIZE;
	return v;
}

var room_Part = function(){
	
	var geometry = new THREE.Geometry();
	geometry.vertices.push(
		new THREE.Vector3(-192,0,-192),
		new THREE.Vector3(-192,0,192),
		new THREE.Vector3(192,0,192),
		new THREE.Vector3(192,0,-192),
		new THREE.Vector3(-192,0,-192)
	);
	
	this.type = 0;
	this.border = new THREE.Line(geometry);
	this.object = 0;
	this.edges = [];
	this.width = 8;
	this.elevation = 0;
	this.height = 192;
	this.center = new THREE.Vector3();
	this.radius = 0;
	
	this.border.dad = this;
	
	this.update = function(){
		this.border.geometry.vertices[0].copy(this.border.geometry.vertices[this.border.geometry.vertices.length-1]);
		this.border.geometry.verticesNeedUpdate = true;
		this.border.geometry.dynamic = true;
		this.border.geometry.elementsNeedUpdate = true;
		this.border.geometry.computeLineDistances();
		this.center.x = 0;this.center.y = 0;this.center.z = 0;
		for(var i=0;i<this.border.geometry.vertices.length-1;i++){
			this.center.add(this.border.geometry.vertices[i]);
		}
		this.center.divideScalar(this.border.geometry.vertices.length-1);
		this.radius = 0;
		for(var i=0;i<this.border.geometry.vertices.length-1;i++){
			if(this.border.geometry.vertices[i].clone().sub(this.center).length()>this.radius){
				this.radius = this.border.geometry.vertices[i].clone().sub(this.center).length();
			}
		}
	}
	
	this.compare_edges = function(part){
		
		//compare radii
		//console.log("comparing...");
		if(part.center.clone().sub(this.center).length()>this.radius+part.radius){
			//console.log("trivial");
			return [];
		}
		shared_edges = [];
		
		for(var i=0;i<part.border.geometry.vertices.length-1;i++){
			var test_edge_A = part.border.geometry.vertices[i].clone().sub(part.border.geometry.vertices[i+1]);
			test_edge_A.y = 0;
			for(var j=0;j<this.border.geometry.vertices.length-1;j++){
				var test_edge_B = this.border.geometry.vertices[j].clone().sub(part.border.geometry.vertices[i+1]);
				var test_edge_C = this.border.geometry.vertices[j+1].clone().sub(part.border.geometry.vertices[i+1]);
				test_edge_B.y = 0;
				test_edge_C.y = 0;

				var valid = true;
				if(test_edge_A.x!=0){
					test_edge_B.x = test_edge_B.x/test_edge_A.x;
				}else if(test_edge_B.x!=0){
					valid = false;
				}
				if(test_edge_A.z!=0){
					test_edge_B.z = test_edge_B.z/test_edge_A.z;
				}else if(test_edge_B.z!=0){
					valid = false;
				}
				if(test_edge_B.x!=0&&test_edge_B.z!=0&&Math.abs(test_edge_B.x-test_edge_B.z)>.001){
					valid = false;
				}
				if(test_edge_B.x==0){
					test_edge_B.x=test_edge_B.z;
				}
				
				if(test_edge_A.x!=0){
					test_edge_C.x = test_edge_C.x/test_edge_A.x;
				}else if(test_edge_C.x!=0){
					valid = false;
				}
				if(test_edge_A.z!=0){
					test_edge_C.z = test_edge_C.z/test_edge_A.z;
				}else if(test_edge_C.z!=0){
					valid = false;
				}
				if(test_edge_C.x!=0&&test_edge_C.z!=0&&Math.abs(test_edge_C.x-test_edge_C.z)>.001){
					valid = false;
				}
				if(test_edge_C.x==0){
					test_edge_C.x=test_edge_C.z;
				}
				
				if(valid){
					//console.log(i,j)
					//c is in line
					if((test_edge_B.x<=1&&test_edge_B.x>=0)||(test_edge_C.x<=1&&test_edge_C.x>=0)||(test_edge_B.x<=0&&test_edge_C.x>=1)||(test_edge_C.x<=0&&test_edge_B.x>=1)){
						
						//a point is in region
						//they touch
						//find borders
						var border_low = Math.max(Math.min(test_edge_B.x,test_edge_C.x),0);
						var border_high = Math.min(Math.max(test_edge_B.x,test_edge_C.x),1);
						if(border_low==border_high){
							//point touch, ignoreCase
							continue;
						}
						console.log("-valid")
						shared_edges.push(test_edge_A.clone().multiplyScalar(border_low).add(part.border.geometry.vertices[i+1]));
						shared_edges.push(test_edge_A.clone().multiplyScalar(border_high).add(part.border.geometry.vertices[i+1]));
					}
				}
			}
		}
		
		return shared_edges;
	}
	
	this.add_edge = function(){
		collisionObjects = [this.border];
		roster.CONTROL_MODE = 1;
		
	}
	this.remove_edge = function(){
		collisionObjects = [];
		this.adjust();
		roster.CONTROL_MODE = 2;
		
	}
	this.add = function(position){
		//determine part from position
		var n, temp, test_vector_A, test_vector_B, test_angle_A, test_angle_B;
		for( n=0;n<this.border.geometry.vertices.length-1;n++){
			test_vector_A = this.border.geometry.vertices[n+1].clone().sub(this.border.geometry.vertices[n]);
			test_vector_B = position.clone().sub(this.border.geometry.vertices[n]);
			test_angle_A = test_vector_A.clone().normalize();
			test_angle_B = test_vector_B.clone().normalize();
			if(test_angle_A.dot(test_angle_B)>.99&&test_vector_B.length()<test_vector_A.length()&&test_vector_B.length()>0){//perfect
				temp = position.clone();
				break;
			}
		}
		if(temp===undefined){
			return;
		}
		roster.CONTROL_MODE = 0;
		var geometry = this.border.geometry.clone();
		geometry.vertices.splice(n+1,0,temp);
		
		
		
		this.border = new THREE.Line(geometry);
		this.border.dad = this;
		roster.displayEdit();
		this.update();
		this.adjust();
	}
	this.remove = function(chosen){
		//determine part from position
		var n;
		if(chosen==this.border||chosen===undefined){
			for( n=0;n<objects.length;n++){
				if(objects[n]===this){
					objects.splice(n,1)
					roster.CONTROL_MODE = 0;
					roster.displayEdit();
					return;
				}
			}
		}
		
		var geometry = this.border.geometry.clone();
		for( n=0;n<this.edges.length;n++){
			if(this.edges[n]===chosen){
				geometry.vertices.splice(n+1,1);
				roster.CONTROL_MODE = 0;
				break;
			}
		}
		
		
		console.log(this.border);
		this.border = new THREE.Line(geometry);
		this.border.dad = this;
		roster.displayEdit();
		console.log(this.border);
		this.update();
		this.adjust();
	}
	this.snap = function(){
		for(var i=0;i<this.border.geometry.vertices.length;i++){
			this.border.geometry.vertices[i].copy(SnapToGrid(this.border.geometry.vertices[i]));
			if(this.edges[i]!==undefined){
				this.edges[i].position.copy(SnapToGrid(this.edges[i].position.copy));
			}
		}
		this.update();
		this.adjust();
	}
	
	this.move = function(x,y,z){
		for(var i=0;i<this.border.geometry.vertices.length;i++){
			this.border.geometry.vertices[i].add(new THREE.Vector3(x,y,z));
			if(this.edges[i]!==undefined){
				this.edges[i].position.add(new THREE.Vector3(x,y,z));
			}
		}
		this.update();
	}
	this.set_height = function(height){
		for(var i=0;i<this.border.geometry.vertices.length;i++){
			this.border.geometry.vertices[i].y=height;
		}
		this.update();
	}
	
	this.generate_Inset = function(){
		//create inset area
		
		if(this.width<=0){
			this.object = this.border.clone();
		}
		
		var splits = [];
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
	
	this.generate_object = function(){
		//load inset area
		this.generate_Inset();
		//extrude up
		var path_points = [this.object.geometry.vertices[0].clone().add(new THREE.Vector3(0,this.height,0)),this.object.geometry.vertices[0]];
		var randomSpline = new THREE.CatmullRomCurve3( path_points );
		var extrudeSettings = {
			steps			: 2,
			bevelEnabled	: false,
			extrudePath		: randomSpline
		};
				
		var pts = [];
		for(var i=1;i<this.object.geometry.vertices.length;i++){
			var temp = this.object.geometry.vertices[i].clone();
			temp.sub(this.object.geometry.vertices[0])
			pts.push(new THREE.Vector2(-temp.z,temp.x));
		}
		var shape = new THREE.Shape( pts );
		
		var geometry = new THREE.ExtrudeGeometry( shape, extrudeSettings );
		
		this.object = new THREE.Mesh(geometry);

	}
	
	
	this.adjust = function(){
		this.update();
		this.edges = [];
		//collisionObjects = [];
		for(var i=1;i<this.border.geometry.vertices.length;i++){
			var box = new THREE.Mesh(new THREE.CubeGeometry(16,16,16));
			box.position.x = this.border.geometry.vertices[i].x;
			box.position.y = this.border.geometry.vertices[i].y;
			box.position.z = this.border.geometry.vertices[i].z;
			box.non_focusable = true;
			box.dad = box;
			box.target = this.border.geometry.vertices[i];
			box.owner = this;
			box.remove = function(me){
				this.owner.remove(this);
				
			}
			box.move = function(x,y,z){
				this.target.add(new THREE.Vector3(x,y,z));
				this.position.add(new THREE.Vector3(x,y,z));
				box.owner.update();
			}
			box.snap = function(){
				this.position.copy(SnapToGrid(this.position));
				this.target.copy(SnapToGrid(this.target));
				box.owner.update();
			}
			this.edges.push(box);
			collisionObjects.push(box);
			reality.add(box);
		}
	}
	
}

var door_Part = function(){
	this.type = 1;
	this.enabled = true;
	this.elevation = 0;
	this.widget = new THREE.Mesh(new THREE.BoxGeometry( 16, 16, 16 ));
	this.widget.dad = this;
	this.object = 0;
	this.room_A = 0;
	this.room_B = 0;
	this.width = 128;
	this.height = 128;
	this.base = 0;
	this.rotation = 0;
	this.remove = function(){
		for( n=0;n<doors.length;n++){
			if(doors[n]===this){
				doors.splice(n,1)
				roster.CONTROL_MODE = 0;
				roster.displayEdit();
				return;
			}
		}
	}
	this.generate_object = function(){
		this.object = new THREE.Mesh(new THREE.BoxGeometry( this.width, this.height, 2*Math.max(this.room_A.width,this.room_B.width) ));
		this.object.geometry.rotateY(-this.rotation);
		this.object.geometry.translate(this.widget.position.x,this.widget.position.y,this.widget.position.z);
		this.object.geometry.translate(0,this.height/2+this.base+Math.max(this.room_A.elevation,this.room_B.elevation),0);
		
		//rotate to face normal
	}
	this.adjust = function(){}
	this.snap = function(){
		//need to grab on to a door frame if possible, other wise disable self
		var n, temp, test_vector_A, test_vector_B, test_angle_A, test_angle_B;
		for(n=0;n<merges.length;n++){
			test_vector_A = merges[n].geometry.vertices[1].clone().sub(merges[n].geometry.vertices[0]);
			test_vector_B = this.widget.position.clone().sub(merges[n].geometry.vertices[0]);
			test_angle_A = test_vector_A.clone().normalize();
			test_angle_B = test_vector_B.clone().normalize();
			if(test_angle_A.dot(test_angle_B)>.99&&test_vector_B.length()<test_vector_A.length()&&test_vector_B.length()>0){//perfect
				temp = test_angle_A.clone().multiplyScalar(test_vector_B.length()).add(merges[n].geometry.vertices[0]);
				break;
			}
		}
		if(temp===undefined){
			return;
		}
		//snap to frame somehow
		this.widget.position.copy(temp);
		this.enabled = true;
		this.room_A = merges[n].part_A;
		this.room_B = merges[n].part_B;
		this.rotation = Math.atan2(test_vector_A.z,test_vector_A.x);
		if(this.rotation>Math.PI){
			this.rotation-=Math.PI*2;
		}else if(this.rotation<-Math.PI){
			this.rotation+=Math.PI*2;
		}
	}
	this.move = function(x,y,z){
		this.enabled = false;
		this.widget.position.add(new THREE.Vector3(x,y,z));
	}
}

var Selector = function(){
	this.object = 0;
	this.type = 0;
	
	this.width = 0;
	this.height = 0;
	this.depth = 0;
	
	this.Folder = gui_edit.addFolder('Selected');
	this.update = function(anything){
		
		if(anything===this.object||anything===undefined){
			return;
		}
		
		this.object = anything;
		
		if(anything.non_focusable){
			return;
		}
		roster.displayEdit();
		this.object.adjust();
		for(i=0;i<objects.length;i++){
			if (objects[i]===this.object){
				roster.index=i;
			}
		}
		gui_edit.remove(this.Folder);
		this.Folder = gui_edit.addFolder('Selected');
		if(this.object.type==0){
			this.Folder.add(this.object.border, 'visible',0,1);
			this.Folder.add(this.object, 'add_edge').name("Add Corner");
			this.Folder.add(this.object, 'remove_edge').name("Remove Corner");
			this.Folder.add(this.object, 'remove').name("Remove This");
			this.Folder.add(this.object, "elevation").step(16).name("Elevation").listen().onFinishChange(function(value){this.object.set_height(value)});
			this.Folder.add(this.object, "height").step(64).name("Room Height");
		
			this.Folder.add(this.object, "width").step(4);
		}
		if(this.object.type==1){
			this.Folder.add(this.object, "height").step(64).name("Door Height");
			this.Folder.add(this.object, "base").step(16).name("Door Base");
			this.Folder.add(this.object, "width").step(64).name("Door Width");
		}
		//this.Folder.add(this.object, 'sign', 0, 1);
		
		
		//this.Folder.add(this.object, 'width',0,512).step(32).onChange(function(value){selected.object.update();});
		//this.Folder.add(this.object, 'height',0,512).step(32).onChange(function(value){selected.object.update();});
		//this.Folder.add(this.object, 'depth',0,512).step(32).onChange(function(value){selected.object.update();});
		//this.Folder.add(this.object, 'x').listen().name('Pos-x');
		//this.Folder.add(this.object, 'y').name('Pos-y').onChange(function(value){selected.object.update();});
		//this.Folder.add(this.object, 'z').listen().name('Pos-z');
		//this.Folder.add(this.object.rotation, 'x').name('Rot-x').onChange(function(value){selected.object.update();});
		//this.Folder.add(this.object.rotation, 'y').name('Rot-y').onChange(function(value){selected.object.update();});
		//this.Folder.add(this.object.rotation, 'z').name('Rot-z').onChange(function(value){selected.object.update();});
		this.Folder.open();
	}
	this.refrence = function(id){
		this.update(objects[id]);
	}
}

var Room = function(){
	this.geometry = new THREE.CubeGeometry(180, 180, 180);
	this.thickness = 20;
	this.compSimple = function(){
		var polygon = this.geometry.clone();
		polygon.computeFaceNormals(); // highly recommended...
		return polygon;
	}
	this.compSurfaces = function(){
		var polygon = this.compSimple();
		
		this.shell = polygon.clone();
		var Surfaces = [];
		var count = this.shell.faces.length;
		console.clear();
		if(count>_LIMIT_){count=_LIMIT_;}
		for(i=0;i<count;i++){
			
			console.log("extracting face: "+i);
			//check if this touches a surface
			
			if(i==_STOP_){
				console.log("--Interupt--");
			}
			
			var found=-1;
			for(j=0;j<Surfaces.length;j++){
				//MAJOR
				//if a point shares a BORDER, not 2 vertices, it should be merged
				//compare line segments, not points
				
				var surfacePlane = new THREE.Plane();
				surfacePlane.setFromCoplanarPoints(Surfaces[j].vertices[Surfaces[j].faces[0].c],Surfaces[j].vertices[Surfaces[j].faces[0].b],Surfaces[j].vertices[Surfaces[j].faces[0].a]);
				var facePlane = new THREE.Plane();
				facePlane.setFromCoplanarPoints(this.shell.vertices[this.shell.faces[i].a].clone(),this.shell.vertices[this.shell.faces[i].b].clone(),this.shell.vertices[this.shell.faces[i].c].clone());
				if (Math.abs(surfacePlane.constant-facePlane.constant)<.00001&&Math.abs(surfacePlane.normal.x-facePlane.normal.x)<.00001&&Math.abs(surfacePlane.normal.y-facePlane.normal.y)<.00001&&Math.abs(surfacePlane.normal.z-facePlane.normal.z)<.00001){
					//the surfaces are on the same plane
					//are they touching?
					// no. fucking. clue.
					var touching = false;
					
					//normalize plane before feeding
					
					for(k=0;k<Surfaces[j].faces.length;k++){
						touching = touchingTriangles(this.shell.vertices[this.shell.faces[i].a],this.shell.vertices[this.shell.faces[i].b],this.shell.vertices[this.shell.faces[i].c],Surfaces[j].vertices[Surfaces[j].faces[k].a],Surfaces[j].vertices[Surfaces[j].faces[k].b],Surfaces[j].vertices[Surfaces[j].faces[k].c]);
						if(touching){break;}
					}
					
					//merger
					if(touching){
						if(found<0){
							console.log("  joined to surface "+j);
							//this is the first surface its jointed to
							found=j;
							Surfaces[j].vertices.push(
								this.shell.vertices[this.shell.faces[i].a].clone(),
								this.shell.vertices[this.shell.faces[i].b].clone(),
								this.shell.vertices[this.shell.faces[i].c].clone()
							);
							Surfaces[j].faces.push( new THREE.Face3( Surfaces[j].vertices.length-1, Surfaces[j].vertices.length-2, Surfaces[j].vertices.length-3 ) );
							Surfaces[found].mergeVertices();
						}else{
							console.log("  merged surface " + j + " into " + found);
							//merge jointed surfaces
							Surfaces[found].merge(Surfaces[j]);
							Surfaces[found].mergeVertices();
							Surfaces.splice(j,1);
						}
					}
					else{
						console.log("  Dead End...");
					}
				}
			
				var incur=0;
			}
			if(found<0){
				console.log("  created new surface");
				var nextSurface = new THREE.Geometry();
				nextSurface.vertices.push(
					this.shell.vertices[this.shell.faces[i].a].clone(),
					this.shell.vertices[this.shell.faces[i].b].clone(),
					this.shell.vertices[this.shell.faces[i].c].clone()
				);
				nextSurface.faces.push( new THREE.Face3( nextSurface.vertices.length-1, nextSurface.vertices.length-2, nextSurface.vertices.length-3 ) );
				Surfaces.push(nextSurface);
			
			}else{
				found++;
			}
		}
		
		for(i=0;i<Surfaces.length;i++){
			Surfaces[i].computeFaceNormals();
		}
		
		for(i=0; i< this.shell.vertices.length; i++){
			var unit = new THREE.Vector3( 0, 0, 0);
			for(k=0; k<Surfaces.length; k++){
				for(j=0; j<Surfaces[k].faces.length;j++){
					if(Surfaces[k].vertices[Surfaces[k].faces[j].a].equals(this.shell.vertices[i])||Surfaces[k].vertices[Surfaces[k].faces[j].b].equals(this.shell.vertices[i])||Surfaces[k].vertices[Surfaces[k].faces[j].c].equals(this.shell.vertices[i])){
						var normal = Surfaces[k].faces[j].normal.clone();
						unit.add(normal);
						break;
					}
				}
			}
			unit.normalize();
			unit.multiplyScalar(this.thickness);
			unit.negate();
			this.shell.vertices[i].addVectors(polygon.vertices[i], unit);
		}
		console.log('TOTAL SURFACE COUNT: '+Surfaces.length);
		return Surfaces;
	}
	this.compWalls = function(){
		var escape=true;
		
		
		//compute surfaces																		
		var Surfaces = this.compSurfaces();
		
		
		var newline = function(s,e){
			this.count=1;
			this.a = s;
			this.b = e;
			this.is = function(s,e){
				return (this.a.equals(s)&&this.b.equals(e))||(this.b.equals(s)&&this.a.equals(e));
			}
			this.has = function(s){
				return this.a.equals(s)||this.b.equals(s);
			}
			this.next = function(s){
				if(this.a.equals(s)){
					return this.b;
				}
				if(this.b.equals(s)){
					return this.a;
				}
				return s;
			}
		}
		
		//break down surfaces into a set of sets of perimeter vertices.																	--TODO
		
		var finalgroups = [];
		for(i=0;i<Surfaces.length;i++){
			console.log("parsing surface : "+i);
			var perimeterMark = 0;
			//find which sides have 1 refrence by faces
			var edges = [];
			var ownerN = [];
			var ownerS = [];
			
			
			for(j=0;j<Surfaces[i].faces.length;j++){
				var founda=false;
				var foundb=false;
				var foundc=false;
				
				for(k=0;k<edges.length;k++){
					//vertices vs index?
					if(edges[k].is(Surfaces[i].vertices[Surfaces[i].faces[j].a],Surfaces[i].vertices[Surfaces[i].faces[j].b])){
						edges[k].count++;
						founda=true;
					}
					if(edges[k].is(Surfaces[i].vertices[Surfaces[i].faces[j].b],Surfaces[i].vertices[Surfaces[i].faces[j].c])){
						edges[k].count++;
						foundb=true;
					}
					if(edges[k].is(Surfaces[i].vertices[Surfaces[i].faces[j].c],Surfaces[i].vertices[Surfaces[i].faces[j].a])){
						edges[k].count++;
						foundc=true;
					}
				}
				if(!founda){
					edges.push(new newline(Surfaces[i].vertices[Surfaces[i].faces[j].a],Surfaces[i].vertices[Surfaces[i].faces[j].b]));
					ownerN.push(j);
					ownerS.push(0);
				}
				if(!foundb){
					edges.push(new newline(Surfaces[i].vertices[Surfaces[i].faces[j].b],Surfaces[i].vertices[Surfaces[i].faces[j].c]));
					ownerN.push(j);
					ownerS.push(1);
				}
				if(!foundc){
					edges.push(new newline(Surfaces[i].vertices[Surfaces[i].faces[j].c],Surfaces[i].vertices[Surfaces[i].faces[j].a]));
					ownerN.push(j);
					ownerS.push(2);
				}
			
			}
			//now using count, filter out interior edges
			
			var iter = 0;
			while(iter<edges.length){
				if(edges[iter].count!=1){
					edges.splice(iter,1);
					ownerN.splice(iter,1);
					ownerS.splice(iter,1);
				}else{
					iter++;
				}
			}
			
			var rotateQ = new THREE.Quaternion();
			rotateQ.setFromUnitVectors(Surfaces[i].faces[0].normal.clone() , new THREE.Vector3(0,0,1));
			
			//also: using collisions, filter out uneven interiors
			var template = 0;
			while(template<edges.length){
				var response = 0;
				var found=false;
				
				var template_A = Surfaces[i].vertices[Surfaces[i].faces[ownerN[template]].a].clone();
				var template_B = Surfaces[i].vertices[Surfaces[i].faces[ownerN[template]].b].clone();
				var template_C = Surfaces[i].vertices[Surfaces[i].faces[ownerN[template]].c].clone();
				
				var template_X = (template_A.x + template_B.x + template_C.x) / 3;
				var template_Y = (template_A.y + template_B.y + template_C.y) / 3;
				var template_Z = (template_A.z + template_B.z + template_C.z) / 3;
				
				if(ownerS[template]==0){
					template_C = new THREE.Vector3(template_X, template_Y, template_Z);
				}else if(ownerS[template]==1){
					template_A = new THREE.Vector3(template_X, template_Y, template_Z);
				}else{
					template_B = new THREE.Vector3(template_X, template_Y, template_Z);
				}
				
				
				while(response<edges.length){
					if(ownerN[template]==ownerN[response]){
						response++;
						continue;
					}
					var response_A = Surfaces[i].vertices[Surfaces[i].faces[ownerN[response]].a].clone();
					var response_B = Surfaces[i].vertices[Surfaces[i].faces[ownerN[response]].b].clone();
					var response_C = Surfaces[i].vertices[Surfaces[i].faces[ownerN[response]].c].clone();
					
					var response_X = (response_A.x + response_B.x + response_C.x) / 3;
					var response_Y = (response_A.y + response_B.y + response_C.y) / 3;
					var response_Z = (response_A.z + response_B.z + response_C.z) / 3;
					
					if(ownerS[response]==0){
						response_C = new THREE.Vector3(response_X, response_Y, response_Z);
					}else if(ownerS[response]==1){
						response_A = new THREE.Vector3(response_X, response_Y, response_Z);
					}else{
						response_B = new THREE.Vector3(response_X, response_Y, response_Z);
					}
					
					if(touchingTriangles(template_A,template_B,template_C,response_A,response_B,response_C)){
						if(response<template){
							template--;
						}
						edges.splice(response,1);
						ownerN.splice(response,1);
						ownerS.splice(response,1);
						found=true;
					}else{
						response++;
					}
				}
				if(found){
					edges.splice(template,1);
					ownerN.splice(template,1);
					ownerS.splice(template,1);
				}else{
					template++;
				}
			}
			
			
			
			
			
			
			//now trace edges
			console.log("  tracing edges");
			paths = [];
			polarities = [];
			while(edges.length>0){
			
				var path = [];
				var position = edges[0].b.clone();
				var start = edges[0].a.clone();
			
				path.push(start.clone());
				path.push(position.clone());
			
				edges.splice(0,1);
			
				while(!position.equals(start)){
					for(j=0;j<edges.length;j++){
						if(edges[j].has(position)){
							path.push(position.clone());
							position = edges[j].next(position).clone();
							edges.splice(j,1);
							break;
						}
					}
				}
				
				//now filter out straight edges.
				iter = 0;
				while(iter<path.length){
					var front = new THREE.Line3(path[bind(iter-1,path.length)],path[iter]).delta();
					var back = new THREE.Line3(path[iter],path[bind(iter+1,path.length)]).delta();
					var product = front.cross(back);
					if(Math.abs(product.x)<.01&&Math.abs(product.y)<.01&&Math.abs(product.z)<.01){
						path.splice(iter,1);
					}
					else{
						iter++;
					}
				}
				if(path.length>0){
					paths.push(path);
				}
			}
			
			console.log("  tracing perimeter");
			
			//determine permiter and hole sets
			var maxP = 0;
			var maxV = -999999999
			for(j=0;j<paths.length;j++){
				for(k=0;k<paths[j].length;k++){
					var temp = paths[j][k].clone().applyQuaternion(rotateQ);
					if(temp.X>maxV){
						maxP = j;
						maxV - temp.X;
					}
				}
			}
			
			var perimeter = [];
			for(j=0;j<paths[maxP].length;j++){
				perimeter.push(paths[maxP][j].clone());
			}
			
			paths.splice(maxP,1);
			
			
			console.log("  merging perimeters: " + paths.length);
			
			//now (through splittin) add holes to perimeter
			//find closest point, merge that path, splice from paths
			while(paths.length>0){
				var minDistance = 999999999;
				var holeP = 0;
				var holeI = 0;
				var permI = 0;
				for(j=0;j<paths.length;j++){
					for(k=0;k<paths[j].length;k++){
						for(l=0;l<perimeter.length;l++){
							if (perimeter[l].distanceTo(paths[j][k])<minDistance){
								holeP = j;
								holeI = k;
								permI = l;
								minDistance = perimeter[l].distanceTo(paths[j][k]);
							}
						}
					}	
				}
				//distance of holeI + 1 to permI +/- 1  determines order of adding
				
				
				if(perimeter[bind(permI-1,perimeter.length)].distanceTo(paths[holeP][bind(holeI+1,paths[holeP].length)])<perimeter[bind(permI+1,perimeter.length)].distanceTo(paths[holeP][bind(holeI+1,paths[holeP].length)])){
					perimeter.splice(permI+1,0,perimeter[permI]);
					for(j=paths[holeP].length;j>=0;j--){
						perimeter.splice(permI+1,0,paths[holeP][bind(j+holeI,paths[holeP].length)]);
					}
				}else{
					perimeter.splice(permI+1,0,perimeter[permI]);
					for(j=0;j<=paths[holeP].length;j++){
						perimeter.splice(permI+1,0,paths[holeP][bind(j+holeI,paths[holeP].length)]);
					}
				
				}
				
				
				
				//perimeter.splice(permI,0,perimeter[permI])
				paths.splice(holeP,1);
			}
			
			var totalangle = 0;
				
			//determine concavity polarity for surface and marks
			for(j=0;j<perimeter.length;j++){
				var angle = relativeAngle(perimeter, j, rotateQ);
				totalangle+=angle;
				
				var temppass = 0;
			}
			
			//if the total angle in not ~equal~ to +/- tau SOMETHING IS VERY WRONG
			//also: maybe rotate the vectors FIRST then set z to 0 and normalize, reduce error and prevent issues with non-flat surfaces (which shouldn't exist...)
			
			var polarity=1;
			if(totalangle<0){
				polarity=-1;
			}
			
			console.log("  determine concavity: "+ polarity);
			
			//form groups!
			
			//start at next concave (or 0 if none)
			//n=2
			//create face of 0..n-1..n
			//repeat until you hit end or concave at n
			//eliminate all points BETWEEN 0 and n
			//check if -1..0..n..n+1 is straight and appropriately delete
			//repeat
			
			console.log("  forming groups: "+perimeter.length);
			
			var indexs = [];
			var sourceSurface = new THREE.Geometry();
			
			for(j=0;j<perimeter.length;j++){
				sourceSurface.vertices.push(perimeter[j].clone());
				indexs[j]=j;
			}
			var gAA = 0;
			escape=true;
			while(perimeter.length>2&&escape){
				gAA++;
				if(gAA>100){
					console.log("##### Possible broken case #####");
					console.log("Details:");
					for(ii=0;ii<perimeter.length;ii++){
						console.log(perimeter[ii].x + " " + perimeter[ii].y + " " +perimeter[ii].z);
					}
					escape=0;
				}
				var tempSurface = sourceSurface.clone();
				//new surface
				var step = 0;
				var size = 1;
				for(j=0;j<perimeter.length;j++){
					if(relativeAngle(perimeter, j, rotateQ)*polarity<0){
						var breach = [];
						breach.push(perimeter[bind(j+1,perimeter.length)].clone());
						breach.push(perimeter[j].clone());
						var width=0;
						for(k=1;k<perimeter.length;k++){
							width=k;
							if(relativeAngle(perimeter, j+k, rotateQ)*polarity<0){
								break;
							}
							breach.splice(2,1,perimeter[bind(j+k+2,perimeter.length)].clone());
							if(relativeAngle(breach, 1, rotateQ)*polarity>0){
								break;
							}
						}
						if(width>size){
							step = j;
							size = width;
						}
						
					}
				}
				
				var breach = [];
				breach.push(perimeter[bind(step+1,perimeter.length)].clone());
				breach.push(perimeter[step].clone());
				
				while(true){
					breach.splice(2,1,perimeter[bind(step+2,perimeter.length)].clone());
					if(relativeAngle(breach, 1, rotateQ)*polarity>0){
						break;
					}
					if(relativeAngle(perimeter, step+1, rotateQ)*polarity<0){
						break;
					}
					if(perimeter.length<3){
						break;
					}
					tempSurface.faces.push( new THREE.Face3(indexs[step],indexs[bind(step+1,indexs.length)],indexs[bind(step+2,indexs.length)]) );
					perimeter.splice(bind(step+1,perimeter.length),1);
					indexs.splice(bind(step+1,indexs.length),1);
					if(step>=perimeter.length){
						step=perimeter.length-1;
					}
						
					
				}
				if(tempSurface.faces.length>0){
					finalgroups.push(tempSurface);
				}
			}
			
			
			var tempstop=0;
		}
		
		
		
		//change to use split surfaces																	--TODO
		//for(i=0;i<Surfaces.length;i++){
			//var NextWall = new THREE.Geometry();
			//NextWall.vertices.push(
			//	SnapToGrid(polygon.vertices[polygon.faces[i].a].clone()),
			//	SnapToGrid(polygon.vertices[polygon.faces[i].b].clone()),
			//	SnapToGrid(polygon.vertices[polygon.faces[i].c].clone()),
			//	
			//	SnapToGrid(this.shell.vertices[polygon.faces[i].a].clone()),
			//	SnapToGrid(this.shell.vertices[polygon.faces[i].b].clone()),
			//	SnapToGrid(this.shell.vertices[polygon.faces[i].c].clone())
			//);
			//
			//NextWall.faces.push( 
			//	new THREE.Face3(1,0,2)
				
				//new THREE.Face3(0,1,3),
				//new THREE.Face3(1,2,4),
				//new THREE.Face3(2,0,5),
				
				//new THREE.Face3(1,4,3),
				//new THREE.Face3(2,5,4),
				//new THREE.Face3(0,3,5)
			//);
		
		console.log('TOTAL GROUP COUNT: '+finalgroups.length);
		return finalgroups;
	}
	this.drawSurfaces = function(){
		var Surfaces = this.compSurfaces();
		Walls = new Array();
		for(i=0;i<Surfaces.length;i++){
			Surfaces[i].computeFaceNormals();
			Walls.push(new THREE.Mesh( Surfaces[i], new THREE.MeshBasicMaterial( { color: GenerateColor(), opacity: 0.4 } ) ));
			if(roster._WIREFRAME_){
				Walls[Walls.length-1].material.wireframe=true;
			}
		}
		return Walls;
	}
	this.drawWalls = function(){
		var Surfaces = this.compWalls();
		Walls = new Array();
		for(i=0;i<Surfaces.length;i++){
			Surfaces[i].computeFaceNormals();
			Walls.push(new THREE.Mesh( Surfaces[i], new THREE.MeshBasicMaterial( { color: GenerateColor(), opacity: 0.4 } ) ));
			if(roster._WIREFRAME_){
				Walls[Walls.length-1].material.wireframe=true;
			}
		}
		return Walls;
	}
};