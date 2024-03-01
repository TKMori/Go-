package main

import "fmt"

// 為醫師建立一個醫療表格, 以便填寫患者的姓名、年齡、以及是否對花生過敏
/*
名字為字串

姓氏也是字串

年齡是整數 int

過敏與否以布林值 bool 表示
*/
type Patient struct {
	Name       string
	FamilyName string
	Age        int
	IsAllergy  bool
}

func (m *Patient) SetName(name string) {
	m.Name = name
}

func (m *Patient) SetFamilyName(familyName string) {
	m.FamilyName = familyName
}

func (m *Patient) SetAge(age int) {
	m.Age = age
}

func (m *Patient) SetIsAllergy(isAllergy bool) {
	m.IsAllergy = isAllergy
}

func main() {
	patient := &Patient{}
	patient.SetName("Bob")
	patient.SetFamilyName("Smith")
	patient.SetAge(34)
	patient.SetIsAllergy(false)

	fmt.Println(patient.Name)
	fmt.Println(patient.FamilyName)
	fmt.Println(patient.Age)
	fmt.Println(patient.IsAllergy)
}
