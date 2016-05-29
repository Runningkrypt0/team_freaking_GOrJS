package main

import (
//	"fmt"
	"os"
)

func main() {

	s := hammer_world{};
	f, err := os.Create("E:/gen.vmf");
	err = err; //KILL ME
	
	hammer_write_world(f,&s);
	f.Close();
	
	//add(&s, &p);
	//fmt.Printf("%d",s.X);
}
