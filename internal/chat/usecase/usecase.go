package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/chat"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type chatUseCase struct {
	chatRepo   chat.Repository
	schoolRepo school.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewChatUseCase(cfg *config.Config, chatRepo chat.Repository, schoolRepo school.Repository, logger logger.Logger) chat.UseCase {
	return &chatUseCase{cfg: cfg, chatRepo: chatRepo, schoolRepo: schoolRepo, logger: logger}
}

func (u *chatUseCase) SaveMessage(msg *models.Message) (*models.Message, error) {
	return u.chatRepo.SaveMessage(msg)
}

func (u *chatUseCase) Create(channel *models.Channel) (*models.Channel, error) {
	return u.chatRepo.Create(channel)
}

func (u *chatUseCase) GetAllByUser(userId uint) (*[]models.Channel, error) {
	return u.chatRepo.GetAllByUser(userId)
}

func (u *chatUseCase) GetById(id uint) (*models.Channel, error) {
	return u.chatRepo.GetById(id)
}

func (u *chatUseCase) GetAllPossibleChatStudent(user *models.User) (*[]models.User, error) {
	students, err := u.schoolRepo.GetSchoolStudents(*user.SchoolId)
	if err != nil {
		return nil, err
	}

	existingChannels, err := u.GetAllByUser(user.ID)
	if err != nil {
		return nil, err
	}

	var filteredStudents []models.User

	for _, student := range *students {
		existingChannel := false

		for _, channel := range *existingChannels {
			if channel.FirstUserId == student.ID || channel.SecondUserId == student.ID {
				existingChannel = true
				break
			}
		}

		if !existingChannel {
			filteredStudents = append(filteredStudents, student)
		}
	}

	return &filteredStudents, nil
}

func (u *chatUseCase) GetAllPossibleChatTeacher(user *models.User) (*[]models.User, error) {
	teachers, err := u.schoolRepo.GetSchoolTeachers(*user.SchoolId)
	if err != nil {
		return nil, err
	}

	existingChannels, err := u.GetAllByUser(user.ID)
	if err != nil {
		return nil, err
	}

	var filteredTeachers []models.User

	for _, teacher := range *teachers {
		existingChannel := false

		for _, channel := range *existingChannels {
			if channel.FirstUserId == teacher.ID || channel.SecondUserId == teacher.ID {
				existingChannel = true
				break
			}
		}

		if !existingChannel {
			filteredTeachers = append(filteredTeachers, teacher)
		}
	}

	return &filteredTeachers, nil
}
