package main

const getTestByIdScripts = "SELECT q.id, q.text_question, q.picture, q.type_question, q.id_topic from questions q where id_topic=$1"

const SelectQuestionsByIdTopic = "SELECT q.*, a_q.correctness, a.text_answer, a.picture from questions q, answer_question a_q, answer_options a  where  q.id=(select min(q.id) from questions q where q.id_topic=$1 and q.id>$2) and q.id=a_q.id_question and a.id=a_q.id_answer"

const getAllAQ  = "SELECT a.id, a.id_question, a.id_answer, a.correctness from answer_question a where id_question=$1"
const getAllAQ8  = "SELECT a.id, a.id_question, a.id_answer, a.correctness, a.type_t from answer_question a where id_question=$1"
const getAllAQ1  = "SELECT a.id, a.id_question, a.id_answer, a.correctness, a.mark from answer_question a where id_question=$1"
const getAllAnsw = "SELECT a.id, a.text_answer, a.table_type, a.table_name, a.correct from answer_options a where id=$1"

const getAllStudent = "SELECT a.firstname_student, a.name_student, a.login_student, a.paswd_student, g.number_group from students a, groups g where a.id_group=g.id and a.id=$1"
const getAllTopics = "SELECT a.id, a.name_topic from topics a"
const getAQTwst = "select a.mark, q.text_question,  a.id_question, aa.text_answer from answer_question a, questions q, answer_options aa where a.id in (select id_answer_question from student_answer_details where id_test=$1) and a.id_question=q.id and a.id_answer=aa.id"
const getAnswersi = "select d.mark_student, t.name_topic from student_answers d, topics t where d.id=$1 and d.id_topic=t.id"

const getAllQuestionsById = "SELECT a.id from questions a where id_topic=$1" 
const getAllAnswersById = "SELECT a.id, a.mark from answer_question a where id_question=$1"

const addStudentTest = "Insert INTO student_answers(id_student, date_test, id_topic) VALUES ($1, $2, $3)"
const addStudentTestAns = "Insert INTO student_answer_details(id_answer_question, id_test) VALUES ($1, $2)"
const UpdStudentTest = "UPDATE student_answers SET mark_student=$1, date_test=$2 WHERE id=$3;"

const AllGrop = "select id, number_group from groups;"
const AllStud = "select s.id, s.firstname_student, s.name_student, s.otch_student, g.number_group from students s, groups g where s.id_group=g.id"

const addGroupid = "insert into groups(number_group) VALUES ($1)"
const deleteGroupScripts  = "DELETE FROM groups WHERE id = $1"
const deleteStudentScripts = "DELETE FROM students WHERE id = $1"
const deleteTestTScripts = "DELETE FROM questions WHERE id = $1"
const addStudentScripts = "INSERT INTO students(firstname_student, name_student, otch_student, id_group, login_student, paswd_student) VALUES ($1,$2,$3,$4,$5,$6);"
const getStudentById = "select s.id, s.firstname_student, s.name_student, s.otch_student, g.number_group from students s, groups g where s.id_group=g.id and s.id=$1"
const editStudentScripts = "UPDATE students SET firstname_student=$1,name_student=$2, otch_student=$3, id_group=$4  WHERE id=$5;"
const selStGroup = "select s.id, s.firstname_student, s.name_student, s.otch_student from students s where id_group=$1"
const selStTop = "select id, name_topic, section from topics"

const getAllTests2 = "select s.id, s.date_test,  s.mark_student, t.name_topic, sum(aq.mark) from student_answers s, topics t, questions q, answer_question aq where s.id_student=$1 and s.id_topic=t.id and q.id=aq.id_question and q.id_topic=t.id and q.type='test' group by s.id, t.name_topic"
const selStudentGroup = "select s.id, s.firstname_student, s.name_student, s.otch_student, count(sa.id) from students s, student_answers sa where s.id_group=$1 and s.id=sa.id_student group by s.id"
const getAllTests = "select s.id, s.date_test,  s.mark_student, t.name_topic, sum(aq.mark) from student_answers s, topics t, questions q, answer_question aq where s.id_student=$1 and s.id_topic=t.id and q.id=aq.id_question and s.id_topic=$2 and q.id_topic=t.id and q.type='test' group by s.id, t.name_topic"

const getGropusAnsOpt ="select aq.mark, a.text_answer, q.text_question, q.type_question from answer_question aq, answer_options a, questions q where a.id=aq.id_answer and aq.id_question=q.id and aq.id=$1"
const editExamScripts = "UPDATE answer_question SET mark=$1 WHERE id=$2;"