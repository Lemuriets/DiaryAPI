package usecase

import (
	"github.com/Lemuriets/diary/internal/users"
	"github.com/Lemuriets/diary/model"
)

type UseCase struct {
	Repository users.Repository
}

func NewUseCase(repo users.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

func (uc *UseCase) GetById(id uint) model.User {
	user, err := uc.Repository.GetById(id)

	if err != nil {
		return model.User{}
	}

	return user
}

func (uc *UseCase) Update(id uint, updateFields map[string]interface{}) error {
	return uc.Repository.Update(id, updateFields)
}

func (uc *UseCase) Delete(id uint) error {
	return uc.Repository.Delete(id)
}

func (uc *UseCase) MultipleDelete(ids []uint) error {
	return uc.Repository.MultipleDelete(ids)
}

func (uc *UseCase) AddGrade(userId uint, gradeId uint) error {
	if gradeId == 0 {
		return users.SchoolIsNotExists
	} else if userId == 0 {
		return users.UndefinedUser
	}
	uc.Repository.Update(userId, map[string]interface{}{
		"grade_id": gradeId,
	})

	return nil
}
