package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

type projects struct {
	database           service.Database
	paymentRepository  GRepository[model.Payment]
	userRepository     GRepository[model.User]
	categoryRepository GRepository[model.Category]
}

func Projects(
	database service.Database,
	paymentRepository GRepository[model.Payment],
	userRepository GRepository[model.User],
	categoryRepository GRepository[model.Category],
) Repository {
	return &projects{
		database,
		paymentRepository,
		userRepository,
		categoryRepository,
	}
}

func (r *projects) ModelName() string {
	return "project"
}

func (r *projects) ModelNamePlural() string {
	return "projects"
}

func (r *projects) NewEmpty() interface{} {
	return &model.Project{}
}

func (r *projects) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var description string
	scanAndPanic(row, &id, &name, &description)
	idString := strconv.FormatInt(id, 10)
	project := model.Project{
		id,
		name,
		description,
		r.paymentRepository.List("WHERE ProjectId = " + idString),
		r.categoryRepository.List("WHERE ProjectId = " + idString),
		r.userRepository.List("WHERE Id IN (SELECT UserId FROM ProjectUser WHERE ProjectId = " + idString + ")"),
	}
	return &project
}

func (r *projects) NewFromRequest(request *http.Request, id int64) interface{} {
	users := strings.Join(request.Form["Users"], ",")
	idString := strconv.FormatInt(id, 10)
	return &model.Project{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Description"),
		r.paymentRepository.List("WHERE ProjectId = " + idString),
		r.categoryRepository.List("WHERE ProjectId = " + idString),
		r.userRepository.List("WHERE Id IN (" + users + ")"),
	}
}

func (r *projects) Find(id int64) interface{} {
	var result []model.Project
	r.database.Select(r.ModelName(), &result, "*", r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *projects) List(conditions ...string) interface{} {
	var result []model.Project
	r.database.Select(r.ModelName(), &result, "*", r.NewFromQuery, conditions...)
	return result
}

func (r *projects) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *projects) Insert(entity interface{}) {
	project := entity.(*model.Project)
	result := r.database.Insert(r.ModelName(), project)
	projectId, _ := result.LastInsertId()
	for _, user := range project.Users {
		r.database.Insert(r.ModelName(), &model.ProjectUser{0, projectId, user.Id})
	}
}

func (r *projects) Update(entity interface{}) {
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
			r.database.Insert("ProjectUser", &model.ProjectUser{0, newProject.Id, newUser.Id})
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

	r.database.Update(r.ModelName(), newProject)
}

func (r *projects) Delete(entity interface{}) {
	r.database.Delete(r.ModelName(), entity)
}
