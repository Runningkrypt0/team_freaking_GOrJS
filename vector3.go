package main

import (
)

type Vector3 struct{
	X int;
	Y int;
	Z int;
};

func vector_div_sca(a *Vector3, b int){
	a.X = a.X/b;
	a.Y = a.Y/b;
	a.Z = a.Z/b;
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

func vector_dot(a *Vector3, b *Vector3) int{
	return a.X*b.X+a.Y*b.Y+a.Z*b.Z;
}

