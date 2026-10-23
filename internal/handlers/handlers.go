package handlers

import "github.com/DrummDaddy/buildservice/internal/models"

func GetProjectName() string {
	p := models.Project{ID: 1, Name: "ЖК солнечный"}
	return p.Name
}
