package main

type Vector3 struct{
	X int;
	Y int;
	Z int;
};

func add(a *Vector3, b *Vector3){
	a.X = a.X + b.X;
	a.Y = a.Y + b.Y;
	a.Z = a.Z + b.Z;
	return 
}


