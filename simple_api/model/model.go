package model

type Weapon struct {
	WeaponType string `json:"type"`
	WeaponName string `json:"name"`
}

type WeaponsRequest struct {
	Weapons []Weapon `json:"Weapon"`
}
