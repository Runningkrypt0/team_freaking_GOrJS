package main

import (
	"os"
)

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

func hammer_write_face(file *os.File, b *hammer_face){

}

type hammer_solid struct{
	Faces []hammer_face;
}

func hammer_write_solid(file *os.File, b *hammer_solid){

}

type hammer_entity struct{
	Keyes []string;
	Values []string;
}

func hammer_write_entity(file *os.File, b *hammer_entity){

}

type hammer_world struct{
	Solids []hammer_solid;
	Entities []hammer_entity;
}

func hammer_write_world(file *os.File, b *hammer_world){

	//write header
	file.WriteString(hammer_world_header);
	//write solids
	//write world close
	file.WriteString(`
	}
	`);
	file.Sync();
	//write entities
	//write footer
	file.WriteString(hammer_world_footer);
	file.Sync();
}