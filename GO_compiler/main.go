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

func make_Stuff(file *os.File){
	var pos int = 0;
	var data []byte;
	var temp string;
	
	my_world := hammer_world{};
	f, _ := os.Create("E:/gen_B.vmf");
	
	for pos>=0 {
		data, pos = read_Chunk(file,pos);
		temp = string(data[:len(data)-1])
		if(temp=="O"){
			fmt.Printf("New Floor:\n")
			new_Floor := room_Floor{}
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Floor.Height,_ = strconv.Atoi(temp)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Floor.Width,_ = strconv.Atoi(temp)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Floor.Elevation,_ = strconv.Atoi(temp)
			
			for pos>=0 {
				var tpos int;
				data_X, tpos := read_Chunk(file,pos);
				temp = string(data_X[:len(data_X)-1])
				value_X, err := strconv.Atoi(temp)
				if(err!=nil){
					break;
				}
				data, pos = read_Chunk(file,tpos);
				temp = string(data[:len(data)-1])
				value_Z, _ := strconv.Atoi(temp)
				new_Floor.Positions.Push(Vector3{float64(value_X),float64(value_Z),float64(new_Floor.Elevation)})
				//fmt.Printf("-Point: %f %f %f\n",new_Floor.Positions[len(new_Floor.Positions)-1].X,new_Floor.Positions[len(new_Floor.Positions)-1].Y,new_Floor.Positions[len(new_Floor.Positions)-1].Z)
			}
			
			holder := make([]Vector3, new_Floor.Positions.Length)
			
			fmt.Printf("%d  #################\n",new_Floor.Positions.Length)
			
			for i:=0; i<len(holder); i++ {
				holder[i] = new_Floor.Positions.Pop().(Vector3)
			}
			
			fmt.Printf("%d  #################\n",new_Floor.Positions.Length)
			holder = inset(holder, float64(new_Floor.Width))
			
			for i:=0; i<len(holder); i++ {
				new_Floor.Positions.Append(holder[i])
			}
			
			fmt.Printf("%d  #################\n",new_Floor.Positions.Length)
			
			new_Solids := decompose_Room(new_Floor)
			
			for n:=0; n<len(new_Solids);n++ {
				hammer_fix_solid(&new_Solids[n])
				my_world.Solids = append(my_world.Solids, new_Solids[n]);
			}
			
			fmt.Printf("-Height: %d\n",new_Floor.Height)
			fmt.Printf("-Width: %d\n",new_Floor.Width)
			fmt.Printf("-Elevation: %d\n",new_Floor.Elevation)
			
		}else if(temp=="D"){
			fmt.Printf("New Door:\n")
			new_Door := room_Door{}
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			X,_ := strconv.Atoi(temp)
			new_Door.Position.X = float64(X);
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			Y,_ := strconv.Atoi(temp)
			new_Door.Position.Y = float64(Y);
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			Z,_ := strconv.Atoi(temp)
			new_Door.Position.Z = float64(Z);
			
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Height,_ = strconv.Atoi(temp)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Width,_ = strconv.Atoi(temp)
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			new_Door.Width,_ = strconv.Atoi(temp)
			
			fmt.Printf("-Height: %d\n",new_Door.Height)
			fmt.Printf("-Width: %d\n",new_Door.Width)
			fmt.Printf("-Base: %d\n",new_Door.Base)
			fmt.Printf("-Point: %f %f %f\n",new_Door.Position.X,new_Door.Position.Y,new_Door.Position.Z)
		}
	}
	
	hammer_write_world(f,&my_world);
	fmt.Printf("\nDONE\n")
	f.Close();
}

type room_Floor struct{
	Positions Stack;
	Height int;
	Elevation int;
	Width int;
}

func validate(tri []Vector3, leftovers room_Floor) bool{
	test_A := tri[0].Clone()
	test_B := tri[2].Clone()
	test_A.Sub(&tri[1])
	test_B.Sub(&tri[1])
	test_C := Vector3{}
	
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
		
		dot00 := test_A.Dot(&test_A)
		dot01 := test_A.Dot(&test_B)
		dot02 := test_A.Dot(&test_C)
		dot11 := test_B.Dot(&test_B)
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

func decompose_Room(target room_Floor) (Triangles []hammer_solid){
	nT := make([]Vector3,3)
	nT[0] = target.Positions.Pop().(Vector3)
	nT[1] = target.Positions.Pop().(Vector3)
	nT[2] = target.Positions.Pop().(Vector3)
	Triangles = make([]hammer_solid,0)
	safety := 300
	
	for true {
		safety--
		if(safety<1){break}
		
		if(validate(nT,target)){
			//out ABC
			Triangles = append(Triangles,hammer_make_floor(nT,float64(target.Elevation),float64(target.Elevation)-32))
			nT[1] = nT[2]
		}else{
			target.Positions.Append(nT[0])
			nT[0] = nT[1]
			nT[1] = nT[2]
		}
		if(target.Positions.Length<1){
			break
		}
		nT[2] = target.Positions.Pop().(Vector3)
	}
	return
}

type room_Door struct{
	Position Vector3;
	Height int;
	Base int;
	Width int;
}

func main() {

	flag.Parse();
	fmt.Printf("Opening: %s \n",fileName);
	f_parsed, err := os.Open(fileName);
	if(err!=nil){
		log.Fatal(err)
	}
	
	test_stack := Stack{}
	test_stack.Push("are")
	test_stack.Push("how")
	test_stack.Push("hello")
	test_stack.Append("you?")
	fmt.Printf(test_stack.Bottom.Value.(string)+"\n")
	fmt.Printf(test_stack.Pop().(string)+"\n")
	fmt.Printf(test_stack.Pop().(string)+"\n")
	fmt.Printf(test_stack.Pop().(string)+"\n")
	fmt.Printf(test_stack.Pop().(string)+"\n")
	
	make_Stuff(f_parsed);
	
	f_parsed.Close();

	//test
	
	

	my_world := hammer_world{};
	f, _ := os.Create("E:/gen.vmf");
	//err = err; //KILL ME
	//dim := Vector3{128.0,256.0,128.0};
	//pos := Vector3{0,0,0};
	
	//vec_a := Vector3{-256,-256,-1000};
	//vec_b := Vector3{256,-256,-800};
	
	vec_list := make([]Vector3, 0)
	vec_list=append(vec_list,Vector3{-256,-256,0});
	vec_list=append(vec_list,Vector3{-256,256,0});
	vec_list=append(vec_list,Vector3{256,256,0});
	//vec_list=append(vec_list,Vector3{256,-256,0});
	
	//my_solid_a := hammer_make_box(&dim,&pos);
	//hammer_fix_solid(&my_solid_a);
	//my_solid_b := hammer_make_wall(&vec_a,&vec_b,16.0);
	//hammer_fix_solid(&my_solid_b);
	my_solid_c := hammer_make_floor(vec_list,-1000,-1032);
	hammer_fix_solid(&my_solid_c);
	
	//my_world.Solids = append(my_world.Solids, my_solid_a);
	//my_world.Solids = append(my_world.Solids, my_solid_b);
	my_world.Solids = append(my_world.Solids, my_solid_c);
	
	hammer_write_world(f,&my_world);
	f.Close();
	
	//add(&s, &p);
	//fmt.Printf("%d",s.X);
}
