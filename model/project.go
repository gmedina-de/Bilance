package model

type Project struct {
	Id          int64
	Name        string
	Description string
	Payments    []Payment
	Categories  []Category
	Users       []User
}

//func (r Project) NewFromQuery(row *sql.Rows) *Project {
//	scanAndPanic(row, &r.Id, &r.Name, &r.Description)
//	return &r
//}
//
//func (r Project) NewFromRequest(request *http.Request, id int64) *Project {
//	users := strings.Join(request.Form["Users"], ",")
//	r.Id = id
//	r.Name = request.Form.Get("Name")
//	r.Description = request.Form.Get("Description")
//	return &model.Project{
//		,
//		r.paymentRepository.List("WHERE ProjectId = " + idString),
//		r.categoryRepository.List("WHERE ProjectId = " + idString),
//		r.userRepository.List("WHERE Id IN (" + users + ")"),
//	}
//}

type ProjectUser struct {
	Id        int64
	ProjectId int64
	UserId    int64
}
