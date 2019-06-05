package main

import (
	"log"
	"fmt"
)
type TopisNames struct {
	Id int
	Name_topic string
	Section_topic int
	Questions_topic []QuestionsNames
}
type QuestionsNames struct {
	Id int 
	Type_question int
	Text_question string
	Notates_question string
	Type_q string
	Cont int
	Answers_Question []AnswersNames
}
type AnswersNames struct {
	Id int 
	Mark float64
	Text_answer string
	Id_answer int
	Table_n string
	Type_ques int
}

func SelQuestAns() []TopisNames{
	rows0, err0 := db.Query("Select id, name_topic, section from topics")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []TopisNames{}
	for rows0.Next() {
		test0 := TopisNames{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Name_topic, &test0.Section_topic)
	if err0 != nil {
			log.Fatal(err0)
		}
		test0.Questions_topic = FQuestionstopic(test0.Id)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func SelQuestAns1(i string) []TopisNames{
	rows0, err0 := db.Query("Select id, name_topic, section from topics")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []TopisNames{}
	for rows0.Next() {
		test0 := TopisNames{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Name_topic, &test0.Section_topic)
	if err0 != nil {
			log.Fatal(err0)
		}
		test0.Questions_topic = FQuestionstopic1(test0.Id, i)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func FQuestionstopic(id int) []QuestionsNames{
	rows0, err0 := db.Query("Select q.id, q.type_question, q.text_question, q.notates, q.type, count(aq.id) from questions q, answer_question aq  where q.id_topic=$1 and aq.id_question=q.id group by q.id", id)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []QuestionsNames{}
	for rows0.Next() {
		test0 := QuestionsNames{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Type_question,	&test0.Text_question, &test0.Notates_question, &test0.Type_q, &test0.Cont)	 
	if err0 != nil {
			log.Fatal(err0)
		}
		test0.Answers_Question = FAnswersQuestion(test0.Id, test0.Type_question)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func FQuestionstopic1(id int, i string) []QuestionsNames{
	rows0, err0 := db.Query("Select q.id, q.type_question, q.text_question, q.notates, q.type, count(aq.id) from questions q, answer_question aq  where q.id_topic=$1 and aq.id_question=q.id and q.type=$2 group by q.id", id, i)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []QuestionsNames{}
	for rows0.Next() {
		test0 := QuestionsNames{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Type_question,	&test0.Text_question, &test0.Notates_question, &test0.Type_q, &test0.Cont)	 
	if err0 != nil {
			log.Fatal(err0)
		}
		test0.Answers_Question = FAnswersQuestion(test0.Id, test0.Type_question)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func deleteTestById(id int) error{
	res, err := db.Exec(deleteTestTScripts, id)
	if err != nil{
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0{
		return fmt.Errorf("Помилка видалення студента!")
	}
	return nil
}
func FAnswersQuestion(id int, t int) []AnswersNames{
	rows0, err0 := db.Query("select aq.id, aq.mark, a.text_answer, a.id, a.table_name from answer_question aq, answer_options a where aq.id_question=$1 and aq.id_answer=a.id", id)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []AnswersNames{}
	for rows0.Next() {
		test0 := AnswersNames{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Mark,	&test0.Text_answer, &test0.Id_answer, &test0.Table_n)	 
	if err0 != nil {
			log.Fatal(err0)
		}
		test0.Type_ques = t
					defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
