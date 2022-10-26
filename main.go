package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(Name string) Profile {
	var profile Profile
	profile.Name = Name

	if profile.Name == "Sasuke" {
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	} else if profile.Name == "Goku" {
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	} else if profile.Name == "Naruto" {
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	}
	return profile
}

func PowerUp(p Profile, i int) Profile {
	var powerUp Profile

	powerUp.Name = p.Name
	powerUp.Health = p.Health * i
	powerUp.Power = p.Power * i
	powerUp.Exp = p.Exp * i
	return powerUp
}

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name :", powerUp.Name)
	fmt.Println("Health :", powerUp.Health)
	fmt.Println("Power :", powerUp.Power)
	fmt.Println("Exp :", powerUp.Exp)
}
