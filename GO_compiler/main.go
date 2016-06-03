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
				new_Floor.Positions = append(new_Floor.Positions,Vector3{float32(value_X),float32(new_Floor.Elevation),float32(value_Z)})
				fmt.Printf("-Point: %f %f %f\n",new_Floor.Positions[len(new_Floor.Positions)-1].X,new_Floor.Positions[len(new_Floor.Positions)-1].Y,new_Floor.Positions[len(new_Floor.Positions)-1].Z)
			
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
			new_Door.Position.X = float32(X);
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			Y,_ := strconv.Atoi(temp)
			new_Door.Position.Y = float32(Y);
			data, pos = read_Chunk(file,pos);
			temp = string(data[:len(data)-1])
			Z,_ := strconv.Atoi(temp)
			new_Door.Position.Z = float32(Z);
			
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
}

type room_Floor struct{
	Positions []Vector3;
	Height int;
	Elevation int;
	Width int;
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
	
	make_Stuff(f_parsed);
	
	f_parsed.Close();


	my_world := hammer_world{};
	f, err := os.Create("E:/gen.vmf");
	err = err; //KILL ME
	dim := Vector3{128.0,256.0,128.0};
	pos := Vector3{0,0,0};
	
	vec_a := Vector3{-256,-256,-1000};
	vec_b := Vector3{256,-256,-800};
	
	vec_list := make([]Vector3, 0)
	vec_list=append(vec_list,Vector3{-256,-256,0});
	vec_list=append(vec_list,Vector3{-256,256,0});
	vec_list=append(vec_list,Vector3{256,256,0});
	vec_list=append(vec_list,Vector3{256,-256,0});
	
	my_solid_a := hammer_make_box(&dim,&pos);
	hammer_fix_solid(&my_solid_a);
	my_solid_b := hammer_make_wall(&vec_a,&vec_b,16.0);
	hammer_fix_solid(&my_solid_b);
	my_solid_c := hammer_make_floor(vec_list,-1000,-1032);
	hammer_fix_solid(&my_solid_c);
	
	my_world.Solids = append(my_world.Solids, my_solid_a);
	my_world.Solids = append(my_world.Solids, my_solid_b);
	my_world.Solids = append(my_world.Solids, my_solid_c);
	
	hammer_write_world(f,&my_world);
	f.Close();
	
	//add(&s, &p);
	//fmt.Printf("%d",s.X);
}
