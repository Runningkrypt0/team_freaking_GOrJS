package main

import (
//	"fmt"
	"os"
)

func main() {

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
