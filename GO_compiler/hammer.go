package main

import (
	"bytes"
	"os"
	"strconv"
	"fmt"
	"math"
)

var face_int int = 1000;

func hammer_print_vector(a *Vector3) string{
	var buffer bytes.Buffer;
	buffer.WriteString(fmt.Sprintf("(%f %f %f)",a.X,a.Y,a.Z));
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

func (me *hammer_face)Normal()(Vector3){
	A := me.A.Clone()
	B := me.B.Clone()
	
	A.Sub(&me.C)
	B.Sub(&me.C)
	
	Normal := A.Cross(&B)
	
	return Normal
}

func (me *hammer_face)Gen_UV()([]int){
	UV := make([]int,6)
	Normal := me.Normal()
	Xc := math.Abs(Normal.X/Normal.Length())
	Yc := math.Abs(Normal.Y/Normal.Length())
	Zc := math.Abs(Normal.Z/Normal.Length())
	if(Xc>Yc){
		if(Zc>Xc){
			//Z
			UV = []int{1,0,0,0,-1,0}
		}else{
			//X
			UV = []int{0,1,0,0,0,-1}
		}
	}else{
		if(Zc>Yc){
			//Z
			UV = []int{1,0,0,0,-1,0}
		}else{
			//X
			UV = []int{1,0,0,0,0,-1}
		}
	}
	
	
	return UV
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
	
	UVs := face.Gen_UV();
	
	file.WriteString(`"
			"material" "TOOLS/TOOLSNODRAW"
			"uaxis" "[`+strconv.Itoa(UVs[0])+` `+strconv.Itoa(UVs[1])+` `+strconv.Itoa(UVs[2])+` 0] 0.25"
			"vaxis" "[`+strconv.Itoa(UVs[3])+` `+strconv.Itoa(UVs[4])+` `+strconv.Itoa(UVs[5])+` 0] 0.25"
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
		centre.Add(&element.A);
		centre.Add(&element.B);
		centre.Add(&element.C);
		count = count + 3;
	}
	centre.Divide(float64(count));
	
	for i,element := range solid.Faces {
		temp_A = element.A.Clone()
		temp_B = element.C.Clone()
		temp_A.Sub(&element.B)
		temp_B.Sub(&element.B)
		temp_A = temp_A.Cross(&temp_B)
		temp_B = centre.Clone()
		temp_B.Sub(&element.B)
		if(temp_A.Dot(&temp_B)>0){
			//reverse
			solid.Faces[i].C.Copy(&element.B)
			solid.Faces[i].B.Copy(&element.C)
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