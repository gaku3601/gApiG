package api

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Api struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	AccessPath string `json:"access_path"`
}

func createApi(name string, path string, method string, accessPath string) (*Api, error) {
	con := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable", "postgres", "app", "localhost", "5432")
	db, _ := gorm.Open("postgres", con)
	defer db.Close()
	api := &Api{
		Name:       name,
		Path:       path,
		Method:     method,
		AccessPath: accessPath,
	}
	return api, db.Create(&api).Error
}
