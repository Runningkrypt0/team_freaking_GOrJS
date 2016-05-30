package main

import(
	"fmt"
	"math"
)

func hammer_make_box(dim *Vector3, pos *Vector3) hammer_solid{ //rotation?

	//8 vectors
	v_nnn := Vector3{pos.X-dim.X/2,pos.Y-dim.Y/2,pos.Z-dim.Z/2};
	v_nnp := Vector3{pos.X-dim.X/2,pos.Y-dim.Y/2,pos.Z+dim.Z/2};
	v_npn := Vector3{pos.X-dim.X/2,pos.Y+dim.Y/2,pos.Z-dim.Z/2};
	v_npp := Vector3{pos.X-dim.X/2,pos.Y+dim.Y/2,pos.Z+dim.Z/2};
	v_pnn := Vector3{pos.X+dim.X/2,pos.Y-dim.Y/2,pos.Z-dim.Z/2};
	v_pnp := Vector3{pos.X+dim.X/2,pos.Y-dim.Y/2,pos.Z+dim.Z/2};
	v_ppn := Vector3{pos.X+dim.X/2,pos.Y+dim.Y/2,pos.Z-dim.Z/2};
	v_ppp := Vector3{pos.X+dim.X/2,pos.Y+dim.Y/2,pos.Z+dim.Z/2};

	//1 solid
	solid := hammer_solid{};
	//6 faces
	solid.Faces = append(solid.Faces,hammer_face{v_nnn,v_npn,v_pnn});
	solid.Faces = append(solid.Faces,hammer_face{v_nnp,v_npp,v_pnp});
	solid.Faces = append(solid.Faces,hammer_face{v_nnn,v_nnp,v_pnn});
	solid.Faces = append(solid.Faces,hammer_face{v_npp,v_npn,v_ppn});
	solid.Faces = append(solid.Faces,hammer_face{v_npp,v_npn,v_nnp});
	solid.Faces = append(solid.Faces,hammer_face{v_ppp,v_ppn,v_pnp});
	
	return solid;
}

func hammer_make_wall(start *Vector3, end *Vector3, thickness float32) hammer_solid{ //rotation?
	
	dim := Vector3{};
	vector_clone(&dim,end);
	vector_sub(&dim,start);
	dim = vector_ortho_Z(&dim);
	vector_normalize(&dim);
	vector_mul_sca(&dim,thickness)
	fmt.Printf(hammer_print_vector(&dim));
	
	//8 vectors
	v_nnn := Vector3{start.X-dim.X/2,start.Y-dim.Y/2,start.Z};
	v_nnp := Vector3{start.X+dim.X/2,start.Y+dim.Y/2,start.Z};
	v_npn := Vector3{start.X-dim.X/2,start.Y-dim.Y/2,end.Z};
	v_npp := Vector3{start.X+dim.X/2,start.Y+dim.Y/2,end.Z};
	v_pnn := Vector3{end.X-dim.X/2,end.Y-dim.Y/2,start.Z};
	v_pnp := Vector3{end.X+dim.X/2,end.Y+dim.Y/2,start.Z};
	v_ppn := Vector3{end.X-dim.X/2,end.Y-dim.Y/2,end.Z};
	v_ppp := Vector3{end.X+dim.X/2,end.Y+dim.Y/2,end.Z};

	//1 solid
	solid := hammer_solid{};
	//6 faces
	solid.Faces = append(solid.Faces,hammer_face{v_nnn,v_npn,v_pnn});
	solid.Faces = append(solid.Faces,hammer_face{v_nnp,v_npp,v_pnp});
	solid.Faces = append(solid.Faces,hammer_face{v_nnn,v_nnp,v_pnn});
	solid.Faces = append(solid.Faces,hammer_face{v_npp,v_npn,v_ppn});
	solid.Faces = append(solid.Faces,hammer_face{v_npp,v_npn,v_nnp});
	solid.Faces = append(solid.Faces,hammer_face{v_ppp,v_ppn,v_pnp});
	
	return solid;
}

func hammer_make_floor(Corners []Vector3, bottom float32, top float32) hammer_solid{
	solid := hammer_solid{};
	temp := hammer_face{};
	for i,_ := range Corners {
		a := Corners[i]
		b := Corners[int(math.Mod(float64(i+1),float64(len(Corners))))];
		temp = hammer_face{Vector3{a.X,a.Y,top},Vector3{a.X,a.Y,bottom},Vector3{b.X,b.Y,bottom}};
		solid.Faces = append(solid.Faces, temp);
	}
	
	temp = hammer_face{Vector3{Corners[0].X,Corners[0].Y,top},Vector3{Corners[1].X,Corners[1].Y,top},Vector3{Corners[2].X,Corners[2].Y,top}};
	solid.Faces = append(solid.Faces, temp);
	
	temp = hammer_face{Vector3{Corners[0].X,Corners[0].Y,bottom},Vector3{Corners[1].X,Corners[1].Y,bottom},Vector3{Corners[2].X,Corners[2].Y,bottom}};
	solid.Faces = append(solid.Faces, temp);
	
	return solid;
}

func hammer_move_solid(solid *hammer_solid, vector *Vector3){
	for i,_ := range solid.Faces {
		vector_add(&solid.Faces[i].A,vector);
		vector_add(&solid.Faces[i].B,vector);
		vector_add(&solid.Faces[i].C,vector);
	}
}
//func hammer_rotate_solid()