package repositorys

import "github.com/sonhineboy/gsadmin/service/app/models"

type TestRepository struct {
	BaseRepository
}

func (r *TestRepository) SetModel() {
	r.Model = models.Role{}
}

func NewTestRepository() *TestRepository {
	var re TestRepository
	re.SetModel()
	return &re
}
