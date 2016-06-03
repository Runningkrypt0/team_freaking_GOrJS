package main

import (
	"bytes"
	"os"
	"strconv"
	"fmt"
)

var face_int int = 1000;

func hammer_print_vector(a *Vector3) string{
	var buffer bytes.Buffer;
	buffer.WriteString("(");
	buffer.WriteString(strconv.Itoa(int(a.X)));
	buffer.WriteString(" ");
	buffer.WriteString(strconv.Itoa(int(a.Y)));
	buffer.WriteString(" ");
	buffer.WriteString(strconv.Itoa(int(a.Z)));
	buffer.WriteString(")");
	return buffer.String();
}

type hammer_face struct{
	A Vector3;
	B Vector3;
	C Vector3;
	
	//todo:
	//add u,v axises and normal calculation
	//add materials
	//add lightmap scales
	//add smoothing groups (?)
}


func hammer_write_face(file *os.File, face *hammer_face, id int){
	//write something about the key value pair
	file.WriteString(`
		side
		{
			"id" "`);
	file.WriteString(strconv.Itoa(face_int));	
	face_int=face_int+1;
	file.WriteString(`"
			"plane" "`);
	file.WriteString(hammer_print_vector(&face.A));
	file.WriteString(" ");
	file.WriteString(hammer_print_vector(&face.B));
	file.WriteString(" ");
	file.WriteString(hammer_print_vector(&face.C));
	
	uaxis := Vector3{}
	vaxis := Vector3{}
	vector_clone(&uaxis,&face.A)
	vector_clone(&uaxis,&face.C)
	vector_sub(&uaxis,&face.B);
	vector_normalize(&uaxis)
	vector_sub(&vaxis,&face.B);
	vector_normalize(&vaxis)
	
	file.WriteString(`"
			"material" "TOOLS/TOOLSNODRAW"
			"uaxis" "[`);
	//file.WriteString(strconv.FormatFloat(float64(uaxis.X), 'f', 2, 32)+" "+strconv.FormatFloat(float64(uaxis.Y), 'f', 2, 32)+" "+strconv.FormatFloat(float64(uaxis.Z), 'f', 2, 32))
	file.WriteString(`1 0 0 0]0.25"
			"vaxis" "[`);
	
	//file.WriteString(strconv.FormatFloat(float64(vaxis.X), 'f', 2, 32)+" "+strconv.FormatFloat(float64(vaxis.Y), 'f', 2, 32)+" "+strconv.FormatFloat(float64(vaxis.Z), 'f', 2, 32))	
	file.WriteString(`0 1 0 0]0.25"
			"rotation" "0"
			"lightmapscale" "16"
			"smoothing_groups" "0"
		}
`);
}

type hammer_solid struct{
	Faces []hammer_face;
}

func hammer_fix_solid(solid *hammer_solid){
	//compute a point inside the shape
	centre := Vector3{};
	var count int = 0;
	temp_A := Vector3{};
	temp_B := Vector3{};
	
	for _,element := range solid.Faces {
		vector_add(&centre,&element.A);
		vector_add(&centre,&element.B);
		vector_add(&centre,&element.C);
		count = count + 3;
	}
	vector_div_sca(&centre, float32(count));
	
	for i,element := range solid.Faces {
		vector_clone(&temp_A,&element.A);
		vector_clone(&temp_B,&element.C);
		vector_sub(&temp_A,&element.B);
		vector_sub(&temp_B,&element.B);
		temp_A = vector_cross(&temp_A,&temp_B);
		vector_clone(&temp_B,&centre);
		vector_sub(&temp_B,&element.B);
		if(vector_dot(&temp_A,&temp_B)>0){
			//reverse
			vector_clone(&solid.Faces[i].C,&element.B);
			vector_clone(&solid.Faces[i].B,&element.C);
			fmt.Printf("switch");
		}
	}
	
	//now check each face and invert it if backwards
	
	
}

func hammer_write_solid(file *os.File, solid *hammer_solid, id int){
	//write header
	
	file.WriteString(hammer_solid_header);
	file.WriteString(strconv.Itoa(id));
	file.WriteString(`"
`);
	
	//write solids
	for _,element := range solid.Faces {
		hammer_write_face(file,&element,id);
	}
	
	//write footer
	file.WriteString(hammer_solid_footer);
	file.Sync();
}

type hammer_entity struct{
	Keyes []string;
	Values []string;
}

func hammer_write_entity(file *os.File, entity *hammer_entity, id int){
	//write header
	file.WriteString(hammer_entity_header);
	file.WriteString(strconv.Itoa(id));
	file.WriteString(`"
`);
	
	//write solids
	for i,element := range entity.Keyes {
		//write something about the key value pair
		file.WriteString(`"`);
		file.WriteString(element);
		file.WriteString(`" "`);
		file.WriteString(entity.Values[i]);
		file.WriteString(`"
`);
	}
	
	//write footer
	file.WriteString(hammer_entity_footer);
	file.Sync();
}

type hammer_world struct{
	Solids []hammer_solid;
	Entities []hammer_entity;
}

func hammer_write_world(file *os.File, world *hammer_world){
	var id = 3;
	//write header
	file.WriteString(hammer_world_header);
	file.Sync();
	
	//write solids
	for _,element := range world.Solids {
		hammer_write_solid(file, &element, id);
		id=id+1;
	}
	
	//write world close
	file.WriteString(`
}`);
	file.Sync();
	
	//write entities
	for _,element := range world.Entities {
		hammer_write_entity(file, &element, id);
		id=id+1;
	}
	
	//write footer
	file.WriteString(hammer_world_footer);
	file.Sync();
}