package models

import(
	"gorm.io/gorm"
)

type Patient struct{
	gorm.Model
	Name string; 
	Age int;
	Disease string;
	Address string;
	AssignedTo string;
}