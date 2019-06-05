package main

import (
	"log"
)

type Top1 struct {
	Id int
	Name string
	Tests []Stud_test1
}
type 	Stud_test1 struct {
	Id int
	Topic_name string
	Data_test string
	Mark1 float64
	Mark_m float64
}
type 	Stud1 struct {
	Fam string
	Name string
	Group string
	Log string
	Pas []byte
}
type responsData1 struct {
	Error2  string
	St2 []Top1
	Studts []Stud1
}
func NewUsErr (ID int) []responsData1 {
	Toq := []responsData1{}
	toq := responsData1{}
	toq.Error2 = ""
	toq.St2 = NewTop(ID)
	toq.Studts = NewAns(ID)
	Toq = append(Toq, toq)
	return Toq
}
func NewUsErr1 (ID int, er string) []responsData1 {
	Toq := []responsData1{}
	toq := responsData1{}
	toq.Error2 = er
	toq.St2 = NewTop(ID)
	toq.Studts = NewAns(ID)
	Toq = append(Toq, toq)
	return Toq
}
func NewTop(ID int) []Top1 {
	rows, err := db.Query(getAllTopics)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Tops := []Top1{}
	for rows.Next() {
		top := Top1{}
		err := rows.Scan(&top.Id, &top.Name)
		if err != nil {
			log.Fatal(err)
		}
		top.Tests = NewQuest(ID, top.Id)
		defer rows.Close()
		Tops = append(Tops, top)

	}
	return Tops
}
func NewQuest(ID int, Id_s int) []Stud_test1 {
	rows, err := db.Query(getAllTests, ID, Id_s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Quests := []Stud_test1{}
	for rows.Next() {
		quest := Stud_test1{}
		err := rows.Scan(&quest.Id, &quest.Data_test, &quest.Mark1, &quest.Topic_name, &quest.Mark_m)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		Quests = append(Quests, quest)

	}
	return Quests
}
func NewAns(ID int) []Stud1 {
	rows, err := db.Query(getAllStudent, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Quests1 := []Stud1{}
	for rows.Next() {
		ans := Stud1{}
		err := rows.Scan(&ans.Fam, &ans.Name, &ans.Log, &ans.Pas, &ans.Group)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		Quests1 = append(Quests1, ans)

	}
	return Quests1
}