package main

import (
//	"fmt"
	"os"
)

func main() {

	my_world := hammer_world{};
	f, err := os.Create("E:/gen.vmf");
	err = err; //KILL ME
	
	my_solid := hammer_solid{};
	
	my_vector_A := Vector3{0,0,0};
	my_vector_B := Vector3{-128,-128,-128};
	my_vector_C := Vector3{128,-128,-128};
	my_vector_D := Vector3{128,-128,128};
	my_vector_E := Vector3{-128,-128,128};
	
	my_face_A := hammer_face{};
	my_face_A.A = my_vector_A;
	my_face_A.B = my_vector_B;
	my_face_A.C = my_vector_C;
	my_face_B := hammer_face{};
	my_face_B.A = my_vector_A;
	my_face_B.B = my_vector_C;
	my_face_B.C = my_vector_D;
	my_face_C := hammer_face{};
	my_face_C.A = my_vector_A;
	my_face_C.B = my_vector_D;
	my_face_C.C = my_vector_E;
	my_face_D := hammer_face{};
	my_face_D.A = my_vector_A;
	my_face_D.B = my_vector_E;
	my_face_D.C = my_vector_B;
	my_face_E := hammer_face{};
	my_face_E.A = my_vector_B;
	my_face_E.B = my_vector_D;
	my_face_E.C = my_vector_E;
	
	my_solid.Faces = append(my_solid.Faces, my_face_A);
	my_solid.Faces = append(my_solid.Faces, my_face_B);
	my_solid.Faces = append(my_solid.Faces, my_face_C);
	my_solid.Faces = append(my_solid.Faces, my_face_D);
	my_solid.Faces = append(my_solid.Faces, my_face_E);
	
	hammer_fix_solid(&my_solid);
	
	my_world.Solids = append(my_world.Solids, my_solid);
	
	hammer_write_world(f,&my_world);
	f.Close();
	
	//add(&s, &p);
	//fmt.Printf("%d",s.X);
}
