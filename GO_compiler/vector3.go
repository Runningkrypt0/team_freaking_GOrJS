package main

import (
	"math"
)

type Vector3 struct{
	X float64
	Y float64
	Z float64
};

func (a *Vector3)Divide(b float64){
	a.X = a.X/b
	a.Y = a.Y/b
	a.Z = a.Z/b
}

func (a *Vector3)Scale(b float64){
	a.X = a.X*b
	a.Y = a.Y*b
	a.Z = a.Z*b
}

func (a *Vector3)Length() float64{
	return math.Sqrt(a.X*a.X+a.Y*a.Y+a.Z*a.Z)
}
func (a *Vector3)Normalize(){
	a.Divide(a.Length())
}

func (a *Vector3)Add(b *Vector3){
	a.X = a.X + b.X
	a.Y = a.Y + b.Y
	a.Z = a.Z + b.Z
}

func (a *Vector3)Sub(b *Vector3){
	a.X = a.X - b.X
	a.Y = a.Y - b.Y
	a.Z = a.Z - b.Z
}

func (a *Vector3)Equals(b *Vector3) bool{
	if(math.Abs(a.X-b.X)<.0001){
		if(math.Abs(a.Y-b.Y)<.0001){
			if(math.Abs(a.Z-b.Z)<.0001){
				return true
			}
		}
	}
	return false
}

func (a *Vector3)Copy(b *Vector3){
	a.X = b.X
	a.Y = b.Y
	a.Z = b.Z
}

func (a *Vector3)Clone() Vector3{
	return Vector3{a.X,a.Y,a.Z}
}


func (a *Vector3)Cross(b *Vector3) Vector3{
	c := Vector3{};
	c.X = a.Y*b.Z - a.Z*b.Y;
	c.Y = a.Z*b.X - a.X*b.Z;
	c.Z = a.X*b.Y - a.Y*b.X;
	return c;
}

func (a *Vector3)Dot(b *Vector3) float64{
	return a.X*b.X+a.Y*b.Y+a.Z*b.Z;
}

