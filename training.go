package main

import (
	"log"
)

type Testing struct {
	Id int
	Text_q string
	Picture_q string
	Type_q int
	Topic_id int
	Not string
	Questions []Question
}
type Question struct {
	Id int
	Id_a int
	Id_q int
	Correct bool
	Answers []Answer
}
type Answer struct {
	Id int
	Text_a string
	Table_type int
	Table_name string
	Corr bool
}

func getTextById(id int) []Testing{
rows, err := db.Query("SELECT q.id, q.text_question, q.picture, q.type_question, q.id_topic, q.notates from questions q where id_topic=$1 and type='train'", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tests := []Testing{}
	for rows.Next() {
		test := Testing{}
		err := rows.Scan(&test.Id, &test.Text_q, &test.Picture_q, &test.Type_q, &test.Topic_id, &test.Not)
		if err != nil {
			log.Fatal(err)
		}
		test.Questions = NewQuestion(test.Id)
		defer rows.Close()
		tests = append(tests, test)
	}
	return tests
}

func NewQuestion(ID int) []Question {
	rows, err := db.Query(getAllAQ, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Questions := []Question{}
	for rows.Next() {
		question := Question{}
		err := rows.Scan(&question.Id, &question.Id_q, &question.Id_a, &question.Correct)
		if err != nil {
			log.Fatal(err)
		}
		question.Answers = NewAnswer(question.Id_a)
		defer rows.Close()
		Questions = append(Questions, question)

	}
	return Questions
}
func NewAnswer(ID int) []Answer {
	rows, err := db.Query(getAllAnsw, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Answers := []Answer{}
	for rows.Next() {
		answer := Answer{}
		err := rows.Scan(&answer.Id, &answer.Text_a, &answer.Table_type, &answer.Table_name, &answer.Corr)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		Answers = append(Answers, answer)

	}
	return Answers
}
