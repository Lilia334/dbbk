package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"database/sql"
)

type Groupsg struct {
	Id int
	Numb_g string
}
type Groupsgy struct {
	Error2 string
	Gr []Groupsg
}
type Studs1 struct {
	Error2 string
	G []Studs
}
type Studs struct {
	Id int
	Fnam string
	Nam string
	Mnam string
	Group string
	Groups []Groupsg
}

func deleteStudentById(id int) error{
	res, err := db.Exec(deleteStudentScripts, id)
	if err != nil{
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0{
		return fmt.Errorf("Помилка видалення студента!")
	}
	return nil
}
func deleteGroupById(id int) error{
	res, err := db.Exec(deleteGroupScripts, id)
	if err != nil{
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0{
		return fmt.Errorf("Помилка видалення групи!")
	}
	return nil
}
func AddGroups(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	fmt.Println("hi")
	if session.Values["isAdmin"] == 1 { 
		
		gr := r.FormValue("numb_group")
		row := db.QueryRow("SELECT q.id, q.number_group from groups q where number_group=$1", gr)
 bk := new(Groupsg)
    err := row.Scan(&bk.Id, &bk.Numb_g)
    fmt.Println(err)
    if err != sql.ErrNoRows { 
y := "Група з даною назвою вже є в базі"
        tests1 := getGGll2( y)          
tmpl := template.Must(template.ParseFiles("src/templates/admin/index.html"))
	tmpl.Execute(w, tests1)
    }else {
    	_, err = db.Exec(addGroupid, gr)
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err)
			return
		}
		http.Redirect(w, r, "/admin", 302)
    } 
		} else {
			http.Redirect(w, r, "/admin", 302)
		}
}
func getGGll() []Groupsgy {
	Toq := []Groupsgy{}
	toq := Groupsgy{}
	toq.Error2 = ""
	toq.Gr = getGropusAll()
	Toq = append(Toq, toq)
	return Toq	
}
func getGGll2(er string) []Groupsgy {
	Toq := []Groupsgy{}
	toq := Groupsgy{}
	toq.Error2 = er
	toq.Gr = getGropusAll()
	Toq = append(Toq, toq)
	return Toq	
}

func getGropusAll() []Groupsg{
rows, err := db.Query(AllGrop)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	groups := []Groupsg{}
	for rows.Next() {
		group := Groupsg{}
		err := rows.Scan(&group.Id, &group.Numb_g)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		groups = append(groups, group)
	}
	return groups
}
func selStud7(er string) []Studs1 {
	Toq := []Studs1{}
	toq := Studs1{}
	toq.Error2 = er
	toq.G = selStud2()
	Toq = append(Toq, toq)  
	return Toq
}
func selStud(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "testApplication")
	if session.Values["isAdmin"] == 1 { 
	Toq := []Studs1{}
	toq := Studs1{}
	toq.Error2 = ""
	toq.G = selStud2()
	Toq = append(Toq, toq)  
		tmpl := template.Must(template.ParseFiles("src/templates/admin/students.html"))
		tmpl.Execute(w, Toq)
		} else {
		http.Redirect(w, r, "/admin", 302)
		return
	}
}
func selStud2() []Studs{
	rows, err := db.Query(AllStud)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	students := []Studs{}
	for rows.Next() {
		student := Studs{}
		err := rows.Scan(&student.Id, &student.Fnam, &student.Nam, &student.Mnam, &student.Group)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		student.Groups = getGropusAll()
		students = append(students, student)
	}
	return students

}
