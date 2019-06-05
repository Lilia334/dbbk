package main

import (
	//"crypto/sha256"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"database/sql"
	"time"
	"crypto/sha256"
)
type Details struct {
	Mark float64
	Text_Question  string
	Id_ans int 
	Texa string
}
type Answrs struct {
	Mark_s float64
	Topp string
	Answer []Details
}
type Topic struct {
	ID            int
	Name_topic    string
	Section 	  int
} 

type Student struct {
	ID            int
} 
type Student0 struct {
	ID            int
} 
type Student1 struct {
	Id int
	Fnam string
	Nam string
	Mnam string
	Group string
} 
type respondData struct {
	Err     string
			Student []Student1
			Group1  []Groupsg
		}
type Marks struct {
	Id float64
}
type R struct {
	ID int
}
type Groups struct {
	Id int
	Numb_g string
	Stud []Students0
}
type Groups16 struct {
	Id int
	Numb_g string
	Stud []Students0
	Stud1 []Students01
}
 type ResTT struct {
 	Id int
 	Gro []Groups16
 	Top []Topic
 }
 
 type Students0 struct {
 	Id int
 	Fam string
 	Nam string
 	Otc string
 	Cont int 
 	Result_st []Student_resultrs
 }
 type Students01 struct {
 	Id int
 	Fam string
 	Nam string
 	Otc string
 }

 type Student_resultrs struct {
 	Id int
 	Date_test string
 	Mark float64
 	Name_top string
 	Mark_m float64
 }
 type ANsOpt struct {
 	Mark float64
 	Text_a string
 	Text_q string 
 	Type_q int
 }
var store = sessions.NewCookieStore([]byte("YOUR_SECRET_KEY"))
var router *mux.Router

func initController() {
	router = mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/", index)
	router.HandleFunc("/login", loginStudent)
	router.HandleFunc("/student", logSt)
	router.HandleFunc("/student/logout", logoutStudent)
	router.HandleFunc("/training", Train)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/update", UpdateUser)
	router.HandleFunc("/training{id}", SelectTest)
	router.HandleFunc("/test{id}", SelectTest1)
	router.HandleFunc("/user-account", UsAcc)
	router.HandleFunc("/savetest", SaveTest)
	router.HandleFunc("/details{id}", TestDet)


	router.HandleFunc("/admin", admin)
	router.HandleFunc("/admin/addGrop", AddGroups)
	router.HandleFunc("/admin/groups/{id}/delete", DeleteGroups)

	router.HandleFunc("/admin/students", selStud)
	router.HandleFunc("/admin/students/{id}/delete", DeleteStudent)
	router.HandleFunc("/admin/addStudent", AddStudent)
	router.HandleFunc("/admin/students/{id}/edit", UpdateStudent)

	router.HandleFunc("/tasktest", Testtask)   
	router.HandleFunc("/admin/test/{id}/delete", DeleteTest)
	router.HandleFunc("/admin/test/{id}/edit", UpdateTest)
	router.HandleFunc("/Quest2", QuesTrain) //тренажер
	router.HandleFunc("/Quest1", QuesTest) //тест
	router.HandleFunc("/stattest", Statistic)
	router.HandleFunc("/grafic", Graficr)

	router.HandleFunc("/results", StudRes) 
	router.HandleFunc("/results1", StudRes1) 
	router.HandleFunc("/results2", StudRes2) 
	router.HandleFunc("/results3", StudRes3)
	router.HandleFunc("/results4", StudRes4)


	router.HandleFunc("/login1", loginAdmin)
	router.HandleFunc("/admin/logout", logoutAdmin)

}
func UpdateTest(w http.ResponseWriter, r *http.Request) {
session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
rows0, err0 := db.Query(getGropusAnsOpt, id)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []ANsOpt{}
	for rows0.Next() {
		test0 := ANsOpt{}    	
		err0 := rows0.Scan(&test0.Mark, &test0.Text_a, &test0.Text_q, &test0.Type_q)
	if err0 != nil {
			log.Fatal(err0)
		}
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
		if r.Method == "GET" { 
			tmpl := template.Must(template.ParseFiles("src/templates/admin/editTests.html"))
			tmpl.Execute(w, tests0)
		} else { 
			if r.Method == "POST" {
	p := r.FormValue("mark_a")
	p1, err := strconv.Atoi(p)
	if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
	_, err1 := db.Exec(editExamScripts, p1, id)
		if err1 != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		http.Redirect(w, r, "/tasktest", 302)
			}
		}
	} else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
	}
func Testtask(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		//робота з завданнями
		SS := SelQuestAns()
		tmpl := template.Must(template.ParseFiles("src/templates/admin/exams.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func QuesTest(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		//робота з завданнями
		i := "train"
		SS := SelQuestAns1(i)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/exams.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func QuesTrain(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		//робота з завданнями
		i := "test"
		SS := SelQuestAns1(i)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/exams.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func Statistic(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		//робота з графіками
		SS := OneGrafic1()
		tmpl := template.Must(template.ParseFiles("src/templates/admin/grafics.html"))
			tmpl.Execute(w, SS)

		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func Graficr(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		//робота з графіками
		id_g1 := r.FormValue("sel2")
		id_g, err := strconv.Atoi(id_g1)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
		}
		id_t1 := r.FormValue("sel1")
		id_t, err := strconv.Atoi(id_t1)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
		}
		SS := OneGrafic(id_g, id_t)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/grafics.html"))
			tmpl.Execute(w, SS)

		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
type Grafic struct {
	Mark_max float64
	Topics []Topic
	Groups []Gryp
	Cont_res []Cont_ER1
	Cont_res2 []Cont_ER2
	Cont_res3 []Cont_ER3
}
type Cont_ER1 struct {
	Id1 int
	K string
}
type Cont_ER2 struct {
	Id1 int
	K string
	K2 string
}
type Cont_ER3 struct {
	Id1 int
	K string
}
type Gryp struct {
	Id int
	Numb_g string
}
func OneGrafic1() []Grafic {
tests00 := []Grafic{}
		test00 := Grafic{}
		test00.Mark_max = 0.0
		u := test00.Mark_max/3
		test00.Topics = OneTopI()
		test00.Groups = GroupSelect()
		test00.Cont_res = SelCount(1,1, u)
		test00.Cont_res2 = SelCount2(1,1, u)
		test00.Cont_res3 = SelCount3(1,1, u)
		tests00 = append(tests00, test00)
		return tests00
}
func OneGrafic(id_g int, id_t int) []Grafic {
rows0, err0 := db.Query("select  sum(aq.mark) from answer_question aq, questions q where q.id_topic=$1 and q.id=aq.id_question and q.type='test'",id_t)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Grafic{}
	for rows0.Next() {
		test0 := Grafic{}    	
		err0 := rows0.Scan(&test0.Mark_max)
		if err0 != nil {
			log.Fatal(err0)
		}
u := test0.Mark_max/3
		test0.Topics = OneTopI()
		test0.Groups = GroupSelect()
		test0.Cont_res = SelCount(id_g,id_t, u)
		test0.Cont_res2 = SelCount2(id_g,id_t, u)
		test0.Cont_res3 = SelCount3(id_g,id_t, u)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
		
		return tests0
}
func SelCount(id int, i int, ii float64) []Cont_ER1 {
	po := 0.0
	rows0, err0 := db.Query("select count(sa.id) from student_answers sa where sa.mark_student>=$1 and sa.mark_student<$2 and sa.id_student in (select id from students where id_group=$3) and sa.id_topic=$4",po, ii, id, i)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Cont_ER1{}
	for rows0.Next() {
		test0 := Cont_ER1{}    	
		err0 := rows0.Scan(&test0.Id1)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.K = strconv.FormatFloat(ii, 'f', 2, 64)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func SelCount2(id int, i int, ii float64) []Cont_ER2 {
	po := ii
	ii = ii*2
	rows0, err0 := db.Query("select count(sa.id) from student_answers sa where sa.mark_student>=$1 and sa.mark_student<$2 and sa.id_student in (select id from students where id_group=$3) and sa.id_topic=$4",po, ii, id, i)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Cont_ER2{}
	for rows0.Next() {
		test0 := Cont_ER2{}    	
		err0 := rows0.Scan(&test0.Id1)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.K = strconv.FormatFloat(po, 'f', 2, 64)
		test0.K2 = strconv.FormatFloat(ii, 'f', 2, 64)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func SelCount3(id int, i int, ii float64) []Cont_ER3 {
	po := ii*2
	ii = ii*3
	rows0, err0 := db.Query("select count(sa.id) from student_answers sa where sa.mark_student>=$1 and sa.mark_student<$2 and sa.id_student in (select id from students where id_group=$3) and sa.id_topic=$4",po, ii, id, i)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Cont_ER3{}
	for rows0.Next() {
		test0 := Cont_ER3{}    	
		err0 := rows0.Scan(&test0.Id1)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.K = strconv.FormatFloat(ii, 'f', 2, 64)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}

func GroupSelect() []Gryp {
	rows0, err0 := db.Query("select id, number_group from groups" )
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Gryp{}
	for rows0.Next() {
		test0 := Gryp{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Numb_g )
		if err0 != nil {
			log.Fatal(err0)
		}
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func StudRes(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		SS := oneErr()
		tmpl := template.Must(template.ParseFiles("src/templates/admin/results.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func StudRes1(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 {
	id_g := r.FormValue("sel5") 
		id_s1 := "sel6"+id_g
		id_s2 := r.FormValue(id_s1)
		id_s, err := strconv.Atoi(id_s2)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		SS := oneErrQ(id_s)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/results2.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func StudRes2(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		id_s2 := r.FormValue("sel1")
		id_s, err := strconv.Atoi(id_s2)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		SS := oneErrQ2(id_s)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/results.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func StudRes3(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		id_s2 := r.FormValue("sel2")
		id_s, err := strconv.Atoi(id_s2)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		SS := oneErrQ3(id_s)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/results2.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func StudRes4(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		id_s2 := r.FormValue("sel3")
		id_s, err := strconv.Atoi(id_s2)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		id_t2 := r.FormValue("sel4")
		id_t, err := strconv.Atoi(id_t2)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		SS := oneErrQ4(id_s, id_t)
		tmpl := template.Must(template.ParseFiles("src/templates/admin/results2.html"))
			tmpl.Execute(w, SS)
		} else {
			http.Redirect(w, r, "/admin", 302)
		return
		}
}
func oneErr() []ResTT {
tests00 := []ResTT{}
		test00 := ResTT{}
		test00.Id = 1
		test00.Gro = oneTopic()
		test00.Top = OneTopI()
		tests00 = append(tests00, test00)
		return tests00
	}
	func oneErrQ(id int) []ResTT {
tests00 := []ResTT{}
		test00 := ResTT{}
		test00.Id = 1
		test00.Gro = oneTopicQ(id)
		test00.Top = OneTopI()
		tests00 = append(tests00, test00)
		return tests00
	}
		func oneErrQ2(id int) []ResTT {
tests00 := []ResTT{}
		test00 := ResTT{}
		test00.Id = 1
		test00.Gro = oneTopicQ2(id)
		test00.Top = OneTopI()
		tests00 = append(tests00, test00)
		return tests00
	}
	func oneErrQ3(id int) []ResTT {
tests00 := []ResTT{}
		test00 := ResTT{}
		test00.Id = 1
		test00.Gro = oneTopicQ3(id)
		test00.Top = OneTopI()
		tests00 = append(tests00, test00)
		return tests00
	}
	func oneErrQ4(id int, idi int) []ResTT {
tests00 := []ResTT{}
		test00 := ResTT{}
		test00.Id = 1
		test00.Gro = oneTopicQ4(id, idi)
		test00.Top = OneTopI()
		tests00 = append(tests00, test00)
		return tests00
	}
func oneTopic() []Groups16 {
	rows0, err0 := db.Query("Select id, number_group from groups")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Groups16{}
	for rows0.Next() {
		test0 := Groups16{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Numb_g)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.Stud = OneStudsel(test0.Id)
		test0.Stud1 = Studsel(test0.Id)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func oneTopicQ(id int) []Groups16 {
	rows0, err0 := db.Query("Select id, number_group from groups")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Groups16{}
	for rows0.Next() {
		test0 := Groups16{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Numb_g)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.Stud = OneStudselQ(test0.Id, id)
		test0.Stud1 = Studsel(test0.Id)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func oneTopicQ2(id int) []Groups16 {
	rows0, err0 := db.Query("Select id, number_group from groups")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Groups16{}
	for rows0.Next() {
		test0 := Groups16{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Numb_g)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.Stud = OneStudselQ2(test0.Id, id)
		test0.Stud1 = Studsel(test0.Id)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func oneTopicQ3(id int) []Groups16 {
	rows0, err0 := db.Query("Select id, number_group from groups")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Groups16{}
	for rows0.Next() {
		test0 := Groups16{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Numb_g)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.Stud = OneStudselQ3(test0.Id, id)
		test0.Stud1 = Studsel(test0.Id)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func oneTopicQ4(id int, idi int) []Groups16 {
	rows0, err0 := db.Query("Select id, number_group from groups")
	if err0 != nil {
		log.Fatal(err0)
	}
	defer rows0.Close()
	tests0 := []Groups16{}
	for rows0.Next() {
		test0 := Groups16{}    	
		err0 := rows0.Scan(&test0.Id, &test0.Numb_g)
		if err0 != nil {
			log.Fatal(err0)
		}
		test0.Stud = OneStudselQ4(test0.Id, id, idi)
		test0.Stud1 = Studsel(test0.Id)
		defer rows0.Close()
		tests0 = append(tests0, test0)
	}
	return tests0
}
func OneStudsel(id int) []Students0{
rows1, err1 := db.Query(selStudentGroup, id)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows1.Close()
	tests1 := []Students0{}
	for rows1.Next() {
		test1 := Students0{}    	
		err1 := rows1.Scan(&test1.Id, &test1.Fam, &test1.Nam, &test1.Otc, &test1.Cont)
		if err1 != nil {
			log.Fatal(err1)
		}
		test1.Result_st = OneTestSelect(test1.Id)
		defer rows1.Close()
		tests1 = append(tests1, test1)
	}
return tests1
		
}
func Studsel(id int) []Students01{
rows1, err1 := db.Query(selStGroup, id)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows1.Close()
	tests1 := []Students01{}
	for rows1.Next() {
		test1 := Students01{}    	
		err1 := rows1.Scan(&test1.Id, &test1.Fam, &test1.Nam, &test1.Otc)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer rows1.Close()
		tests1 = append(tests1, test1)
	}
return tests1
		
}
func OneStudselQ(id int,id_s int) []Students0{
rows1, err1 := db.Query("select s.id, s.firstname_student, s.name_student, s.otch_student, count(sa.id) from students s, student_answers sa where s.id_group=$1 and  s.id=sa.id_student and s.id=$2 group by s.id", id, id_s)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows1.Close()
	tests1 := []Students0{}
	for rows1.Next() {
		test1 := Students0{}    	
		err1 := rows1.Scan(&test1.Id, &test1.Fam, &test1.Nam, &test1.Otc, &test1.Cont)
		if err1 != nil {
			log.Fatal(err1)
		}
		test1.Result_st = OneTestSelect(test1.Id)
		defer rows1.Close()
		tests1 = append(tests1, test1)
	}
return tests1	
}
func OneStudselQ2(id int,id_s int) []Students0{
rows1, err1 := db.Query("select s.id, s.firstname_student, s.name_student, s.otch_student, count(sa.id) from students s, student_answers sa where s.id_group=$1 and  s.id=sa.id_student and sa.id_topic=$2 group by s.id", id, id_s)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows1.Close()
	tests1 := []Students0{}
	for rows1.Next() {
		test1 := Students0{}    	
		err1 := rows1.Scan(&test1.Id, &test1.Fam, &test1.Nam, &test1.Otc, &test1.Cont)
		if err1 != nil {
			log.Fatal(err1)
		}
		test1.Result_st = OneTestSelect2(test1.Id, id_s)
		defer rows1.Close()
		tests1 = append(tests1, test1)
	}
return tests1
		
}
func OneStudselQ3(id int,id_s int) []Students0{
rows1, err1 := db.Query("select s.id, s.firstname_student, s.name_student, s.otch_student, count(sa.id) from students s, student_answers sa where s.id_group=$1 and s.id_group=$2 and s.id=sa.id_student group by s.id", id, id_s)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows1.Close()
	tests1 := []Students0{}
	for rows1.Next() {
		test1 := Students0{}    	
		err1 := rows1.Scan(&test1.Id, &test1.Fam, &test1.Nam, &test1.Otc, &test1.Cont)
		if err1 != nil {
			log.Fatal(err1)
		}
		test1.Result_st = OneTestSelect(test1.Id)
		defer rows1.Close()
		tests1 = append(tests1, test1)
	}
return tests1
		
}
func OneStudselQ4(id int,id_s int, idi int) []Students0{
rows1, err1 := db.Query("select s.id, s.firstname_student, s.name_student, s.otch_student, count(sa.id) from students s, student_answers sa where s.id_group=$1 and s.id_group=$2 and s.id=sa.id_student and sa.id_topic=$3 group by s.id", id, idi, id_s)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows1.Close()
	tests1 := []Students0{}
	for rows1.Next() {
		test1 := Students0{}    	
		err1 := rows1.Scan(&test1.Id, &test1.Fam, &test1.Nam, &test1.Otc, &test1.Cont)
		if err1 != nil {
			log.Fatal(err1)
		}
		test1.Result_st = OneTestSelect2(test1.Id, id_s)
		defer rows1.Close()
		tests1 = append(tests1, test1)
	}
return tests1
		
}
func OneTestSelect(id int) []Student_resultrs{
rows2, err2 := db.Query(getAllTests2, id)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer rows2.Close()
	tests2 := []Student_resultrs{}
	for rows2.Next() {
		test2 := Student_resultrs{}    	
		err2 := rows2.Scan(&test2.Id, &test2.Date_test, &test2.Mark, &test2.Name_top, &test2.Mark_m)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer rows2.Close()
		tests2 = append(tests2, test2)
	}
	return tests2
}
func OneTestSelect2(id int, id_t int) []Student_resultrs{
rows2, err2 := db.Query("select s.id, s.date_test,  s.mark_student, t.name_topic, sum(aq.mark) from student_answers s, topics t, questions q, answer_question aq  where s.id_student=$1 and s.id_topic=t.id and s.id_topic=$2  and q.id=aq.id_question and q.id_topic=t.id and q.type='test' group by s.id, t.name_topic", id, id_t)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer rows2.Close()
	tests2 := []Student_resultrs{}
	for rows2.Next() {
		test2 := Student_resultrs{}    	
		err2 := rows2.Scan(&test2.Id, &test2.Date_test, &test2.Mark, &test2.Name_top, &test2.Mark_m)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer rows2.Close()
		tests2 = append(tests2, test2)
	}
	return tests2
}
func OneTopI() []Topic{
rows2, err2 := db.Query(selStTop)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer rows2.Close()
	tests2 := []Topic{}
	for rows2.Next() {
		test2 := Topic{}    	
		err2 := rows2.Scan(&test2.ID, &test2.Name_topic, &test2.Section)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer rows2.Close()
		tests2 = append(tests2, test2)
	}
	return tests2
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) { 
session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

Toq := []respondData{}
	toq := respondData{}
	toq.Group1 = getGropusAll()  
		if r.Method == "GET" { 
			toq.Err = ""
	toq.Student = getStudentByIdf(id)
	Toq = append(Toq, toq)
			tmpl := template.Must(template.ParseFiles("src/templates/admin/editStudent.html"))
			tmpl.Execute(w, Toq)
		} else {
			if r.Method == "POST" {
					fa := r.FormValue("firstname")
	n := r.FormValue("names")
	o := r.FormValue("middlen")
	p := r.FormValue("sel1")
	p1, err := strconv.Atoi(p)
	if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
	_, err1 := db.Exec(editStudentScripts, fa, n, o, p1, id)
		if err1 != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		
			}
		}
	} else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
}

func getStudentByIdf(id int) []Student1{
	rows, err := db.Query(getStudentById, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	students := []Student1{}
	for rows.Next() {
		student := Student1{}
		err := rows.Scan(&student.Id, &student.Fnam, &student.Nam, &student.Mnam, &student.Group)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		students = append(students, student)
	}
	return students

}
func AddStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	isAdmin := session.Values["isAdmin"]
	if isAdmin == 1 {
		fa := r.FormValue("firstname")
		n := r.FormValue("names")
		o := r.FormValue("middlen")
		p := r.FormValue("sel1")
		password := r.FormValue("pas")
		login := r.FormValue("log")
		hashPassword := sha256.Sum256([]byte(password))
		row := db.QueryRow("SELECT q.id from students q where login_student=$1 and paswd_student=$2", login,hashPassword[:])
 bk := new(Student)
    err := row.Scan(&bk.ID)
    fmt.Println(err)
    if err != sql.ErrNoRows {
    	y := "Даний логін вже зайнятий іншим студентом. Спробуйте інший логін."
        tests1 := selStud7( y)          
tmpl := template.Must(template.ParseFiles("src/templates/admin/students.html"))
	tmpl.Execute(w, tests1)
     } else {
		_, err := db.Exec(addStudentScripts, fa, n, o, p, login, hashPassword[:])
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		http.Redirect(w, r, "/admin/students", 302)
	} 
}else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	isAdmin := session.Values["isAdmin"]
	if isAdmin == 1 {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		err := deleteStudentById(id)
		if err != nil {
			fmt.Printf("Err: %s", err)
			session.AddFlash(err.Error())
			session.Save(r, w)
		}
		http.Redirect(w, r, "/admin/students", 302)
		return
	} else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
}

func DeleteTest(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	isAdmin := session.Values["isAdmin"]
	if isAdmin == 1 {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		err := deleteTestById(id)
		if err != nil {
			fmt.Printf("Err: %s", err)
			session.AddFlash(err.Error())
			session.Save(r, w)
		}
		http.Redirect(w, r, "/tasktest", 302)
		return
	} else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
}
func DeleteGroups(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	isAdmin := session.Values["isAdmin"]
	if isAdmin == 1 {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		err := deleteGroupById(id)
		if err != nil {
			fmt.Printf("Err: %s", err)
			session.AddFlash(err.Error())
			session.Save(r, w)
		}
		http.Redirect(w, r, "/admin", 302)
		return
	} else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
}
func TestDet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTest, _ := strconv.Atoi(vars["id"])
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 { 
	ResDet := NewResDetails(idTest)  
		tmpl := template.Must(template.ParseFiles("src/templates/result_details.html"))
		tmpl.Execute(w, ResDet) } else {
			http.Redirect(w, r, "/", 302)
		} 
}

func NewResDetails(ID int) []Answrs {
	rows, err := db.Query(getAnswersi, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tests1 := []Answrs{}
	for rows.Next() {
		test1 := Answrs{}    	
		err := rows.Scan(&test1.Mark_s, &test1.Topp)
		if err != nil {
			log.Fatal(err)
		}
		test1.Answer = NewAnswr(ID)
		defer rows.Close()
		tests1 = append(tests1, test1)
	}
	return tests1

}
func NewAnswr(ID int) []Details {
	rows, err := db.Query(getAQTwst, ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tests := []Details{}
	for rows.Next() {
		test := Details{}    	
		err := rows.Scan(&test.Mark, &test.Text_Question, &test.Id_ans, &test.Texa)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		tests = append(tests, test)
	}
	return tests
	
}

func SaveTest(w http.ResponseWriter, r *http.Request) {
session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 { 
	m:= 0.0;
	id_s := 0;
	id_t := 0;
	session, _ := store.Get(r, "testApplication")
		as1 := r.FormValue("1")
			ass1, err := strconv.ParseInt(as1, 10, 64)
rows2, err := db.Query("select id_topic from questions where id=(select id_question from answer_question where id=$1)", ass1)
 if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()
	for rows2.Next() {
		st := R{}
		err := rows2.Scan(&st.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer rows2.Close()
		id_t = st.ID 
rows := db.QueryRow("Select id from student_answers where id_student=$1 and id_topic=$2", session.Values["idStudent"], id_t)
 student := new(Student)
    err7 := rows.Scan(&student.ID)
    if err7 == sql.ErrNoRows {
_, err = db.Exec(addStudentTest, session.Values["idStudent"],  time.Now(), id_t)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
} 
rows99 := db.QueryRow("Select id from student_answers where id_student=$1 and id_topic=$2", session.Values["idStudent"], id_t)
 student1 := new(Student)
    err07 := rows99.Scan(&student1.ID)
if err07 != sql.ErrNoRows {
id_s = student1.ID 
_, err = db.Exec("delete from student_answer_details where id_test=$1", student1.ID)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
	for i :=1; i<30; i++ {
ans := strconv.Itoa(i);
		as := r.FormValue(ans)
		if as !="" {
    	ass, err := strconv.ParseInt(as, 10, 64)
			if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
_, err1 := db.Exec(addStudentTestAns, ass, id_s)
		if err1 != nil {
			fmt.Fprintf(w, "Err: %s", err1)
			return
		}
} 
	} 
rows1, err := db.Query("select aq.mark from answer_question aq where id in (select id_answer_question from student_answer_details where id_test=$1)", id_s)
 if err != nil {
		log.Fatal(err)
	}
	defer rows1.Close()
	for rows1.Next() {
mark := Marks{}
		err := rows1.Scan(&mark.Id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows1.Close()
		m = m + mark.Id
		fmt.Println(m)
      _, err = db.Exec(UpdStudentTest, m, time.Now(),  id_s)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}

}

    	}
    }
    	http.Redirect(w, r, "/user-account", 302)
    } else {
http.Redirect(w, r, "/index", 302)
		return
}

}



func UpdateUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 { 
	if r.Method == "POST" {
		login := r.FormValue("logg")
		password := r.FormValue("pass1")
		hashPassword := sha256.Sum256([]byte(password))
		password2 := r.FormValue("pass2")
		hashPassword2 := sha256.Sum256([]byte(password2))
row0 := db.QueryRow("SELECT q.id from students q where id=$1 and paswd_student=$2", session.Values["idStudent"], hashPassword[:])
 bk0 := new(Student0)
    err0 := row0.Scan(&bk0.ID)
    if err0 == sql.ErrNoRows {
    	id_p := session.Values["idStudent"]
iAreaId := id_p.(int)
y := "Ви ввели невірний пароль. Спробуйте ще раз."
        tests1 := NewUsErr1(iAreaId, y)          
tmpl := template.Must(template.ParseFiles("src/templates/user-account.html"))
	tmpl.Execute(w, tests1)
} else {
		 row := db.QueryRow("SELECT q.id from students q where login_student=$1 and paswd_student=$2", login, hashPassword2[:])
 bk := new(Student)
    err := row.Scan(&bk.ID)
    if err != sql.ErrNoRows {
    	    	id_p := session.Values["idStudent"]
iAreaId := id_p.(int)
y := "Спробуйте обрати новий пароль."
        tests1 := NewUsErr1(iAreaId, y)          
tmpl := template.Must(template.ParseFiles("src/templates/user-account.html"))
	tmpl.Execute(w, tests1)
    }  else {
    	_, err := db.Exec("UPDATE students SET login_student=$1, paswd_student=$2 WHERE id=$3;", login, hashPassword2[:], session.Values["idStudent"])
if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/user-account", 302)
}
		} } } else {
			http.Redirect(w, r, "/user-account", 302)
 }
}


func UsAcc(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
id_p := session.Values["idStudent"]
iAreaId := id_p.(int)
        tests1 := NewUsErr(iAreaId)           
tmpl := template.Must(template.ParseFiles("src/templates/user-account.html"))
	tmpl.Execute(w, tests1)

    } else {
    	http.Redirect(w, r, "/", 302)
    }
}

func logSt(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	Topic1 := NewTopics()
	tmpl := template.Must(template.ParseFiles("src/templates/student/index.html"))
    tmpl.Execute(w, Topic1)
 } else {
type responsData struct {
				Error  string
			}
			tmpl := template.Must(template.ParseFiles("src/templates/login.html"))
			tmpl.Execute(w, responsData{""})
}
}

func loginStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	http.Redirect(w, r, "/", 302)
 } else 
	if r.Method == "POST" {
		fmt.Println("post")
		login := r.FormValue("login")
		password := r.FormValue("password")
		hashPassword := sha256.Sum256([]byte(password))
		 row := db.QueryRow("SELECT q.id from students q where login_student=$1 and paswd_student=$2", login, hashPassword[:])
 bk := new(Student)
    err := row.Scan(&bk.ID)
    fmt.Println(err)
    if err == sql.ErrNoRows {
    	fmt.Println("norow")
       type responsData struct {
				Error  string
			}
			tmpl := template.Must(template.ParseFiles("src/templates/login.html"))
			tmpl.Execute(w, responsData{"Невірний логін чи пароль"}) 
    } else {
    	fmt.Println(err)
        session.Values["isStudent"] = 1
        session.Values["idStudent"] = bk.ID;
			err := session.Save(r, w)
			if err != nil {
				fmt.Fprintf(w, "%s\n", err)
			}
			http.Redirect(w, r, "/", 302)
    }
	} else {
		fmt.Println("get")
		type responsData struct {
			Error string
		}
		tmpl, _ := template.ParseFiles("templates/student/login.html")
		tmpl.Execute(w, responsData{""})
	}
}

func logoutStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	session.Values["isStudent"] = 0
	session.Values["idStudent"] = 0
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}


func SelectTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
		idTopic, _ := strconv.Atoi(vars["id"])
		tests1 := getTextById(idTopic)
		session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	tmpl := template.Must(template.ParseFiles("src/templates/student/answer.html"))
	tmpl.Execute(w, tests1)
 } else {
tmpl := template.Must(template.ParseFiles("src/templates/answer.html"))
tmpl.Execute(w, tests1)
 }
}
func SelectTest1(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	vars := mux.Vars(r)
		idTopic, _ := strconv.Atoi(vars["id"])
		tests1 := getTextById2(idTopic)
		session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	tmpl := template.Must(template.ParseFiles("src/templates/answer1.html"))
	tmpl.Execute(w, tests1)
 } else {
tmpl := template.Must(template.ParseFiles("src/templates/index.html"))
tmpl.Execute(w, tests1)
 } } else {
 	http.Redirect(w, r, "/", 302)
		return
 }
}
func index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	Topic1 := NewTopics()
	tmpl := template.Must(template.ParseFiles("src/templates/student/index.html"))
    tmpl.Execute(w, Topic1)
	} else {
		Topic1 := NewTopics()  
		tmpl := template.Must(template.ParseFiles("src/templates/index.html"))
		tmpl.Execute(w, Topic1)
	}
   
}

func Train(w http.ResponseWriter, r *http.Request) {
session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
		Topic1 := NewTopics()  
		tmpl := template.Must(template.ParseFiles("src/templates/student/training.html"))
		tmpl.Execute(w, Topic1)
		} else {
			Topic1 := NewTopics()  
		tmpl := template.Must(template.ParseFiles("src/templates/training.html"))
		tmpl.Execute(w, Topic1)
		}
   
}
func Test(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isStudent"] == 1 {
	Topic1 := NewTopics()  
		tmpl := template.Must(template.ParseFiles("src/templates/test.html"))
		tmpl.Execute(w, Topic1)
	} else {
type responsData struct {
			Error string
		}
		tmpl := template.Must(template.ParseFiles("src/templates/login.html"))
		tmpl.Execute(w, responsData{""})
	}
   
}
func NewTopics() ([]Topic) {
		rows, err := db.Query("SELECT * FROM topics;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Topics := []Topic{}
	for rows.Next() {
		topic := Topic{}
		err := rows.Scan(&topic.ID, &topic.Name_topic, &topic.Section)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		Topics = append(Topics, topic)

	}
	return Topics
}
func admin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	isAdmin := session.Values["isAdmin"]
	if isAdmin == 1 {
		Groups := getGGll()
		tmpl := template.Must(template.ParseFiles("src/templates/admin/index.html"))
		tmpl.Execute(w, Groups)
	} else {
		http.Redirect(w, r, "/login1", 302)
		return
	}
}

func loginAdmin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 {
		http.Redirect(w, r, "/admin", 302)
		return
	}
	if r.Method == "POST" {
		login := r.FormValue("login")
		password := r.FormValue("password")

		if login == os.Getenv("testApp-login") && password == os.Getenv("testApp-password") {
			session.Values["isAdmin"] = 1
			err := session.Save(r, w)
			if err != nil {
				fmt.Fprintf(w, "%s\n", err)
			}
			http.Redirect(w, r, "/admin", 302)
			return
		} else {
			responsData := struct {
				Error string
				Login string
			}{
				Error: "Неверный логин или пароль!",
				Login: login,
			}
			tmpl, _ := template.ParseFiles("src/templates/admin/login.html")
			tmpl.Execute(w, responsData)
		}
	} else if r.Method == "GET" {
		responsData := struct {
			Error string
			Login string
		}{
			Error: "",
			Login: "",
		}
		tmpl, _ := template.ParseFiles("src/templates/admin/login.html")
		tmpl.Execute(w, responsData)
	}
}

func logoutAdmin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
		session.Values["isAdmin"] = 0
	session.Save(r, w)
	http.Redirect(w, r, "/login1", 302)
}
