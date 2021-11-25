package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

type projectRepository struct {
	baseRepository
	paymentRepository Repository
	userRepository    Repository
	typeRepository    Repository
}

func ProjectRepository(
	database service.Database,
	paymentRepository Repository,
	userRepository Repository,
	typeRepository Repository,
) Repository {
	return &projectRepository{
		baseRepository{database: database},
		paymentRepository,
		userRepository,
		typeRepository,
	}
}

func (r *projectRepository) NewEmpty() interface{} {
	return &model.Project{}
}

func (r *projectRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var description string
	row.Scan(&id, &name, &description)
	idString := strconv.FormatInt(id, 10)
	project := model.Project{
		id,
		name,
		description,
		r.paymentRepository.List("WHERE ProjectId = " + idString).([]model.Payment),
		r.typeRepository.List("WHERE ProjectId = " + idString).([]model.Category),
		r.userRepository.List("WHERE Id IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + idString + ")").([]model.User),
	}
	return &project
}

func (r *projectRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	users := strings.Join(request.Form["Users"], ",")
	idString := strconv.FormatInt(id, 10)
	return &model.Project{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Description"),
		r.paymentRepository.List("WHERE ProjectId = " + idString).([]model.Payment),
		r.typeRepository.List("WHERE ProjectId = " + idString).([]model.Category),
		r.userRepository.List("WHERE Id IN (" + users + ")").([]model.User),
	}
}

func (r *projectRepository) Find(id int64) interface{} {
	var result []model.Project
	r.database.Select(&result, r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *projectRepository) List(conditions ...string) interface{} {
	var result []model.Project
	conditions = append(conditions, "ORDER BY Id")
	r.database.Select(&result, r.NewFromQuery, conditions...)
	return result
}

func (r *projectRepository) Insert(entity interface{}) {
	project := entity.(*model.Project)
	result := r.database.Insert(project)
	projectId, _ := result.LastInsertId()
	for _, user := range project.Users {
		r.database.Insert(&model.ProjectUser{0, projectId, user.Id})
	}
}

func (r *projectRepository) Update(entity interface{}) {
	newProject := entity.(*model.Project)
	var newProjectUserIds idRange
	for _, user := range newProject.Users {
		newProjectUserIds = append(newProjectUserIds, user.Id)
	}

	oldProject := r.Find(newProject.Id).(*model.Project)
	var oldProjectUserIds idRange
	for _, user := range oldProject.Users {
		oldProjectUserIds = append(oldProjectUserIds, user.Id)
	}

	for _, newUser := range newProject.Users {
		if !oldProjectUserIds.contains(newUser.Id) {
			r.database.Insert(&model.ProjectUser{0, newProject.Id, newUser.Id})
		}
	}

	for _, oldUser := range oldProject.Users {
		if !newProjectUserIds.contains(oldUser.Id) {
			r.database.MultipleDelete("ProjectUser",
				"WHERE ProjectId = "+strconv.FormatInt(newProject.Id, 10),
				"AND UserId = "+strconv.FormatInt(oldUser.Id, 10),
			)
		}
	}

	r.database.Update(newProject)
}

func (r *projectRepository) Delete(entity interface{}) {
	r.database.Delete(entity)
}
