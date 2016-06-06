package main

import(
	"flag"
	"os"
	"fmt"
	"log"
	"strconv"
)

var fileName string;

func init(){
	flag.StringVar(&fileName, "file","data.txt","File to pull from")
}

func read_Chunk(file *os.File, pos int) ([]byte, int){
	data := make([]byte,1)
	var c_pos int = pos;
	for string(data)!="," {
		_, err := file.ReadAt(data,int64(c_pos))
		c_pos=c_pos+1;
		if(err!=nil){
			break;
		}
	}
	data = make([]byte,c_pos-pos)
	_, err := file.ReadAt(data,int64(pos))
	if(err!=nil){
		return data,-1
	}
	return data,c_pos
}

func make_Stuff(rooms []room_Floor, doors []room_Door){

	my_world := hammer_world{};
	f, _ := os.Create("E:/gen_B.vmf");
	
	for r_id:=0;r_id<len(rooms);r_id++{
		holder := make([]Vector3, 0)
			
		dump := rooms[r_id].Positions.Top
		for dump!=nil {
			holder = append(holder,dump.Value.(Vector3))
			dump = dump.Support
		}
		
		inset_holder := inset(holder,rooms[r_id].Width)
		for i:=0; i<len(inset_holder); i++ {
			rooms[r_id].Insets.Append(inset_holder[i])
		}
		
		//first insert door alcoves into holder
		door_marks := make([]int,0)
		door_refs := make([]int,0)
		door_holder := make([]Vector3,0)
		
		for i:=0; i<len(holder); i++ {
			door_holder = append(door_holder,holder[i])
		
			//now we need to test each portion for doors
			
			
			for j:=0; j<len(doors); j++ {
			
				test_vector_A := Vector3{}
				if(i<len(holder)-1){
					test_vector_A.Copy(&holder[i+1])
				}else{
					test_vector_A.Copy(&holder[0])
				}
				test_vector_A.Sub(&holder[i])
				test_vector_A.Z = 0
				
				
				//if(doors[j].Position.Z<rooms[r_id].Elevation||doors[j].Position.Z>rooms[r_id].Elevation+rooms[r_id].Height){continue}
				test_vector_B := doors[j].Position.Clone()
				test_vector_B.Sub(&holder[i])
				test_vector_B.Z = 0 //need to check if door is in room height range
				ratio := (test_vector_B.Length()/test_vector_A.Length())
				//fmt.Printf("\n")
				//fmt.Printf("-Length Ratio: %f\n",test_vector_B.Length()/test_vector_A.Length())
				if(ratio<1&&ratio>0){
					
					test_vector_A.Normalize()
					test_vector_B.Normalize()
					//fmt.Printf("-Dir A: %f %f %f\n",test_vector_A.X,test_vector_A.Y,test_vector_A.Z)
					//fmt.Printf("-Dir B: %f %f %f\n",test_vector_B.X,test_vector_B.Y,test_vector_B.Z)
					if(test_vector_A.Dot(&test_vector_B)>.99){
						fmt.Printf("-A: %f %f %f\n",holder[i].X,holder[i].Y,holder[i].Z)
						fmt.Printf("-D: %f %f %f\n",doors[j].Position.X,doors[j].Position.Y,doors[j].Position.Z)
						//fmt.Printf("-B: %f %f %f\n",holder[i+1].X,holder[i+1].Y,holder[i+1].Z)
					
						test_vector_A.Scale(doors[j].Width/2+rooms[r_id].Width)
						
						upper_vec_In := doors[j].Position.Clone()
						upper_vec_In.Z = rooms[r_id].Elevation
						fmt.Printf("-Dir B: %f %f %f\n",upper_vec_In.X,upper_vec_In.Y,upper_vec_In.Z)
						lower_vec_In := doors[j].Position.Clone()
						lower_vec_In.Z = rooms[r_id].Elevation
						upper_vec_In.Add(&test_vector_A)
						lower_vec_In.Sub(&test_vector_A)
						
						upper_vec_Out := upper_vec_In.Clone()
						lower_vec_Out := lower_vec_In.Clone()
						test_vector_A.X,test_vector_A.Y = -test_vector_A.Y,test_vector_A.X
						upper_vec_Out.Sub(&test_vector_A)
						lower_vec_Out.Sub(&test_vector_A)
						
						fmt.Printf("-low in: %f %f %f\n",lower_vec_In.X,lower_vec_In.Y,lower_vec_In.Z)
						door_marks = append(door_marks, len(door_holder))
						door_refs = append(door_refs, j)
						door_holder = append(door_holder,lower_vec_In)
						door_holder = append(door_holder,lower_vec_Out)
						door_holder = append(door_holder,upper_vec_Out)
						door_holder = append(door_holder,upper_vec_In)
						
						fmt.Printf("relevant \n")
						
					}
				}
			}
			
			
			
		}
		fmt.Printf("--Norm--\n")
		for i:=0;i<len(holder);i++{
			fmt.Printf(hammer_print_vector(&holder[i])+"\n")
		}
		fmt.Printf("--DOOR--\n")
		for i:=0;i<len(door_holder);i++{
			fmt.Printf(hammer_print_vector(&door_holder[i])+"\n")
		}
		inset_doors := inset(door_holder, rooms[r_id].Width)
		
		var skip int
		for i:=0; i<len(door_holder); i++ {
			skip = -1
			for j:=0;j<len(door_marks);j++{
				if(door_marks[j]==i){
					skip = j
					break
				}
			}
			
			if(skip>=0){
				//make door floor chunk
				door_vecs := make([]Vector3,4)
				door_vecs[0] = door_holder[i]
				door_vecs[1] = inset_doors[i]
				door_vecs[2] = inset_doors[i+3]
				door_vecs[3] = door_holder[i+3]
				wall_solid := hammer_solid{}
				if(doors[door_refs[skip]].Position.Z>rooms[r_id].Elevation){
					wall_solid = hammer_make_floor(door_vecs,doors[door_refs[skip]].Position.Z,rooms[r_id].Elevation)
				}else{
					wall_solid = hammer_make_floor(door_vecs,rooms[r_id].Elevation,rooms[r_id].Elevation-32)
				}
				
				hammer_fix_solid(&wall_solid)
				my_world.Solids = append(my_world.Solids, wall_solid)
			
				i = i+2
				continue
			}
			
			//h(i, i+1), i(i+1,i) also account for i+1 = 0
			wall_vecs := make([]Vector3,4)
			wall_vecs[0] = door_holder[i]
			if(i<len(inset_doors)-1){
				wall_vecs[1] = door_holder[i+1]
				wall_vecs[2] = inset_doors[i+1]
			}else{
				wall_vecs[1] = door_holder[0]
				wall_vecs[2] = inset_doors[0]
			}
			wall_vecs[3] = inset_doors[i]
			
			wall_solid := hammer_make_floor(wall_vecs,rooms[r_id].Elevation,rooms[r_id].Elevation+rooms[r_id].Height)
		
			hammer_fix_solid(&wall_solid)
			my_world.Solids = append(my_world.Solids, wall_solid)
		}
		

		
		new_Solids := decompose_Room(rooms[r_id])
		for n:=0; n<len(new_Solids);n++ {
			hammer_fix_solid(&new_Solids[n])
			my_world.Solids = append(my_world.Solids, new_Solids[n]);
		}
	
	}
	
	hammer_write_world(f,&my_world);
	f.Close()
	
}
func read_Stuff(file *os.File) (ROOMS []room_Floor, DOORS []room_Door){
	var pos int = 0;
	var data []byte;
	var temp string;
	
	ROOMS = make([]room_Floor,0)
	DOORS = make([]room_Door,0)
	
	for pos>=0 {
		data, pos = read_Chunk(file,pos);
		temp = string(data[:len(data)-1])
		if(temp=="O"){
			fmt.Printf("New Floor:\n")
			new_Floor := room_Floor{}
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Floor.Height,_ = strconv.ParseFloat(temp,64)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Floor.Width,_ = strconv.ParseFloat(temp,64)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Floor.Elevation,_ = strconv.ParseFloat(temp,64)
			
			for pos>=0 {
				var tpos int;
				data_X, tpos := read_Chunk(file,pos);
				temp = string(data_X[:len(data_X)-1])
				value_X, err := strconv.ParseFloat(temp,64)
				if(err!=nil){
					break;
				}
				data, pos = read_Chunk(file,tpos);
				temp = string(data[:len(data)-1])
				value_Z, _ := strconv.ParseFloat(temp,64)
				new_Floor.Positions.Push(Vector3{value_X,value_Z,new_Floor.Elevation})
				//fmt.Printf("-Point: %f %f %f\n",new_Floor.Positions[len(new_Floor.Positions)-1].X,new_Floor.Positions[len(new_Floor.Positions)-1].Y,new_Floor.Positions[len(new_Floor.Positions)-1].Z)
			}
			ROOMS = append(ROOMS,new_Floor)
			
			fmt.Printf("-H: %d ",new_Floor.Height)
			fmt.Printf("-W: %d ",new_Floor.Width)
			fmt.Printf("-E: %d\n",new_Floor.Elevation)
			
		}else if(temp=="D"){
			fmt.Printf("New Door:\n")
			new_Door := room_Door{}
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Position.X, _ = strconv.ParseFloat(temp,64)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Position.Z, _ = strconv.ParseFloat(temp,64)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Position.Y, _ = strconv.ParseFloat(temp,64)
			
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Height,_ = strconv.ParseFloat(temp,64)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Width,_ = strconv.ParseFloat(temp,64)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Base,_ = strconv.ParseFloat(temp,64)
			
			fmt.Printf("-H: %d ",new_Door.Height)
			fmt.Printf("-W: %d ",new_Door.Width)
			fmt.Printf("-B: %d\n",new_Door.Base)
			
			DOORS = append(DOORS,new_Door)
		}
	}
	
	
	fmt.Printf("\nDONE\n")
	
	
	return
}

type room_Floor struct{
	Positions Stack
	Insets Stack
	Height float64;
	Elevation float64;
	Width float64;
}

func shared_edge(hull_A []Vector3, hull_B []Vector3) (int,int) {

	for i:=0; i<len(hull_A)-1; i++ {
		for j:=0; j<len(hull_B)-1; j++ {
			if(hull_A[i].Equals(&hull_B[j+1])&&hull_A[i+1].Equals(&hull_B[j])){
				return i,j
			}
		}
		if(hull_A[i].Equals(&hull_B[0])&&hull_A[i+1].Equals(&hull_B[len(hull_B)-1])){
			return i,len(hull_B)-1
		}
	}
	for j:=0;j<len(hull_B)-1;j++ {
		if(hull_A[len(hull_A)-1].Equals(&hull_B[j+1])&&hull_A[0].Equals(&hull_B[j])){
			return len(hull_A)-1,j
		}
	}
	if(hull_A[len(hull_A)-1].Equals(&hull_B[0])&&hull_A[0].Equals(&hull_B[len(hull_B)-1])){
		return len(hull_A)-1,len(hull_B)-1
	}
	return -1,-1
}

func validate_Hull(hull []Vector3, test_point *Vector3) (score int){
	
	score = 0
	
	for i:=0;i<len(hull);i++ {
	
		test_vector := Vector3{}
		if(i<len(hull)-1){
			test_vector.Copy(&hull[i+1])
		}else{
			test_vector.Copy(&hull[0])
		}
		test_vector.Sub(&hull[i])
		test_vector.X,test_vector.Y = -test_vector.Y,test_vector.X
		
		r_test_point := test_point.Clone()
		r_test_point.Sub(&hull[i])
		
		if(test_vector.Dot(&r_test_point)<0){
			score++
		}
	
	}

	return
}


func validate(tri []Vector3, leftovers room_Floor) bool{
	test_A := tri[0].Clone()
	test_B := tri[2].Clone()
	test_A.Sub(&tri[1])
	test_B.Sub(&tri[1])
	test_C := Vector3{}
	
	dot00 := test_A.Dot(&test_A)
	dot01 := test_A.Dot(&test_B)
	dot11 := test_B.Dot(&test_B)
	
	test_A.X, test_A.Y = -test_A.Y,test_A.X //normal
	if(test_A.Dot(&test_B)>=0){return false} //triangle has invalid wrapping
	test_A.X, test_A.Y = test_A.Y,-test_A.X //revert
	
	//does triangle contain any other points
	
	dump := leftovers.Positions.Top
	for dump!=nil {
		//is point in triangle
		temp := dump.Value.(Vector3)
		test_C.Copy(&temp)
		test_C.Sub(&tri[1])
		
		dot02 := test_A.Dot(&test_C)
		dot12 := test_B.Dot(&test_C)

		// Compute barycentric coordinates
		invDenom := 1 / (dot00 * dot11 - dot01 * dot01)
		u := (dot11 * dot02 - dot01 * dot12) * invDenom
		v := (dot00 * dot12 - dot01 * dot02) * invDenom
		// Check if point is in triangle
		if(((u >= 0) && (v >= 0) && (u + v < 1))){
			test_C = dump.Value.(Vector3)
			return false
		}
		dump = dump.Support
	}
	return true
	
}

func decompose_Room(target room_Floor) (Solids []hammer_solid){
	nT := make([]Vector3,3)
	nT[0] = target.Insets.Pop().(Vector3)
	nT[1] = target.Insets.Pop().(Vector3)
	nT[2] = target.Insets.Pop().(Vector3)
	Triangles := make([]([]Vector3),0)
	safety := 300
	
	for true {
		safety--
		if(safety<1){break}
		
		if(validate(nT,target)){
			//out ABC
			Tri := make([]Vector3, 3)
			copy(Tri, nT)
			Triangles = append(Triangles, Tri)
			
			nT[1] = nT[2]
		}else{
			target.Insets.Append(nT[0])
			nT[0] = nT[1]
			nT[1] = nT[2]
		}
		if(target.Insets.Length<1){
			break
		}
		nT[2] = target.Insets.Pop().(Vector3)
	}
	
	//now to do some magic and combine a bunch of triangles...
	//for every combination...
	//if something forms a proper quad, it may only be merged with another quad, which will be done later
	//if (score == 1){combine}
	
	shapes := make([]([]Vector3),0)
	//quads := make([]([]Vector3),0)
	
	var escape bool
	
	
	
	for i:=0;i<len(Triangles);i++{
		escape = false
		//compare it to shapes
		for j:=0;j<len(shapes);j++{
			A,B := shared_edge(Triangles[i],shapes[j])
			if(A>=0){
				//we need to cut up the arrays now
				A_checker := make([]Vector3,1)
				if(A==0){
					A_checker[0] = Triangles[i][2].Clone()
				}else if(A==1){
					A_checker[0] = Triangles[i][0].Clone()
				}else{
					A_checker[0] = Triangles[i][1].Clone()
				}
				if(validate_Hull(shapes[j],&A_checker[0])==1){
					new_Tri := make([]Vector3,0)
					for k:=0;k<len(shapes[j]);k++{
						new_Tri = append(new_Tri,shapes[j][k])
						if(k==B){
							new_Tri = append(new_Tri,A_checker[0])
						}
					}
					copy(shapes[j:], shapes[j+1:])
					shapes = shapes[:len(shapes)-1]

					copy(Triangles[i:], Triangles[i+1:])
					Triangles = Triangles[:len(Triangles)-1]
					
					shapes = append(shapes,new_Tri)
					
					i=-1
					escape = true
				}
			}
			if(escape){break}
		}
		if(escape){continue}
		
		
		//compare it to other triangles
		for j:=0;j<len(Triangles);j++{
			if(i==j){continue}
			A,B := shared_edge(Triangles[i],Triangles[j])
			if(A>=0){
				//we need to cut up the arrays now
				A_checker := make([]Vector3,1)
				if(A==0){
					A_checker[0] = Triangles[i][2].Clone()
				}else if(A==1){
					A_checker[0] = Triangles[i][0].Clone()
				}else{
					A_checker[0] = Triangles[i][1].Clone()
				}
				if(validate_Hull(Triangles[j],&A_checker[0])==1){
					new_Tri := make([]Vector3,0)
					for k:=0;k<len(Triangles[j]);k++{
						new_Tri = append(new_Tri,Triangles[j][k])
						if(k==B){
							new_Tri = append(new_Tri,A_checker[0])
						}
					}
					
					if(i>j){
						copy(Triangles[i:], Triangles[i+1:])
						copy(Triangles[j:], Triangles[j+1:])
					}else{
						copy(Triangles[j:], Triangles[j+1:])
						copy(Triangles[i:], Triangles[i+1:])
					}
					Triangles = Triangles[:len(Triangles)-2]
					
					shapes = append(shapes,new_Tri)
					
					i=-1
					break
				}
			}
		}
	}
	
	//now merge any shapes
	for i:=0;i<len(shapes);i++{
		for j:=0;j<len(shapes);j++{
			A,B := shared_edge(shapes[i],shapes[j])
			if(A>=0){
				//we need to cut up the arrays now
				valid := true
				for k:=0;k<len(shapes[i]);k++ {
					if(k==A){continue}
					if(k==A+1||(k==0&&A==len(shapes[i])-1)){continue}
					
					if(validate_Hull(shapes[j],&shapes[i][k])!=1){
						valid = false
						break
					}
				}
				if(valid){
					new_Tri := make([]Vector3,0)
					for k:=0;k<len(shapes[j]);k++{
						new_Tri = append(new_Tri,shapes[j][k])
						if(k==B){
							for l:=A+2;true;l++{
								if(l>=len(shapes[i])){
									l = l - len(shapes[i])
								}
								if(l==A){break}
								new_Tri = append(new_Tri,shapes[i][l])
							}
							
						}
					}
					
					if(i>j){
						copy(shapes[i:], shapes[i+1:])
						copy(shapes[j:], shapes[j+1:])
					}else{
						copy(shapes[j:], shapes[j+1:])
						copy(shapes[i:], shapes[i+1:])
					}
					shapes = shapes[:len(shapes)-2]

					shapes = append(shapes,new_Tri)
					
					i=-1
					break
				}
			}
		}
	}
	
	for i:=0;i<len(Triangles);i++{
		shapes = append(shapes,Triangles[i])
	}

	
	for i:=0;i<len(shapes);i++{
		Solids = append(Solids,hammer_make_floor(shapes[i],float64(target.Elevation),float64(target.Elevation)-32))
	}
	return
}

type room_Door struct{
	Position Vector3;
	Height float64;
	Base float64;
	Width float64;
}

func main() {

	flag.Parse();
	fmt.Printf("Opening: %s \n",fileName);
	f_parsed, err := os.Open(fileName);
	if(err!=nil){
		log.Fatal(err)
	}

	ROOMS, DOORS := read_Stuff(f_parsed);
	make_Stuff(ROOMS,DOORS)
	
	f_parsed.Close();

	//test
}
