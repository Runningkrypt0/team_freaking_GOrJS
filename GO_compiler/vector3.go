package main

import (
	"math"
)

type Vector3 struct{
	X float32;
	Y float32;
	Z float32;
};

func vector_div_sca(a *Vector3, b float32){
	a.X = a.X/b;
	a.Y = a.Y/b;
	a.Z = a.Z/b;
}

func vector_mul_sca(a *Vector3, b float32){
	a.X = a.X*b;
	a.Y = a.Y*b;
	a.Z = a.Z*b;
}

func vector_normalize(a *Vector3){
	var distance float32 = a.X*a.X+a.Y*a.Y+a.Z*a.Z;
	distance = float32(math.Sqrt(float64(distance)));
	vector_div_sca(a,distance);
	
}

func vector_ortho_X(a *Vector3) Vector3{
	return Vector3{a.X,-a.Z,a.Y};
}
func vector_ortho_Y(a *Vector3) Vector3{
	return Vector3{-a.Z,a.Y,a.X};
}
func vector_ortho_Z(a *Vector3) Vector3{
	return Vector3{-a.Y,a.X,a.Z};
}


func vector_add(a *Vector3, b *Vector3){
	a.X = a.X + b.X;
	a.Y = a.Y + b.Y;
	a.Z = a.Z + b.Z;
}

func vector_sub(a *Vector3, b *Vector3){
	a.X = a.X - b.X;
	a.Y = a.Y - b.Y;
	a.Z = a.Z - b.Z;
}

func vector_clone(a *Vector3, b *Vector3){
	a.X = b.X;
	a.Y = b.Y;
	a.Z = b.Z;
}

func vector_cross(a *Vector3, b *Vector3) Vector3{
	c := Vector3{};
	c.X = a.Y*b.Z - a.Z*b.Y;
	c.Y = a.Z*b.X - a.X*b.Z;
	c.Z = a.X*b.Y - a.Y*b.X;
	return c;
}

func vector_dot(a *Vector3, b *Vector3) float32{
	return a.X*b.X+a.Y*b.Y+a.Z*b.Z;
}

