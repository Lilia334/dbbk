package main

import (
	"log"
)

type Testing2 struct {
	Id int
	Text_q string
	Picture_q string
	Type_q int
	Topic_id int
	Not string
	Questions []Question2
}
type Question2 struct {
	Id int
	Id_a int
	Id_q int
	Correct bool
	Table_typo int
	Answers []Answer2
}
type Answer2 struct {
	Id int
	Text_a string
	Table_type int
	Table_name string
	Corr bool
}

func getTextById2(id int) []Testing2{
rows, err := db.Query("SELECT q.id, q.text_question, q.picture, q.type_question, q.id_topic, q.notates from questions q where id_topic=$1 and type='test'", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tests := []Testing2{}
	for rows.Next() {
		test := Testing2{}
		err := rows.Scan(&test.Id, &test.Text_q, &test.Picture_q, &test.Type_q, &test.Topic_id, &test.Not)
		if err != nil {
			log.Fatal(err)
		}
		test.Questions = NewQuestion2(test.Id)
		defer rows.Close()
		tests = append(tests, test)
	}
	return tests
}

func NewQuestion2(ID int) []Question2 {
	rows, err := db.Query(getAllAQ8, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Questions := []Question2{}
	for rows.Next() {
		question := Question2{}
		err := rows.Scan(&question.Id, &question.Id_q, &question.Id_a, &question.Correct, &question.Table_typo)
		if err != nil {
			log.Fatal(err)
		}
		question.Answers = NewAnswer2(question.Id_a)
		defer rows.Close()
		Questions = append(Questions, question)

	}
	return Questions
}
func NewAnswer2(ID int) []Answer2 {
	rows, err := db.Query(getAllAnsw, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Answers := []Answer2{}
	for rows.Next() {
		answer := Answer2{}
		err := rows.Scan(&answer.Id, &answer.Text_a, &answer.Table_type, &answer.Table_name, &answer.Corr)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		Answers = append(Answers, answer)

	}
	return Answers
}
