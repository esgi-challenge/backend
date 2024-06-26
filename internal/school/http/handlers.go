package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/email"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type schoolHandlers struct {
	cfg           *config.Config
	schoolUseCase school.UseCase
	logger        logger.Logger
}

func NewSchoolHandlers(cfg *config.Config, schoolUseCase school.UseCase, logger logger.Logger) school.Handlers {
	return &schoolHandlers{cfg: cfg, schoolUseCase: schoolUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new school
//	@Description	create new school
//	@Tags			School
//	@Accept			json
//	@Produce		json
//	@Param			school	body		models.SchoolCreate	true	"School infos"
//	@Success		201		{object}	models.School
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/schools [post]
func (u *schoolHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.SchoolCreate

		schoolCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		school := &models.SchoolCreate{
			Name: schoolCreate.Name,
		}
		schoolDb, err := u.schoolUseCase.Create(user, school)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, schoolDb)
	}
}

// Invite
//
//	@Summary		Invite a student to the school
//	@Description	Invite a student to the school
//	@Tags			School
//	@Accept			json
//	@Produce		json
//	@Param			school	body		models.SchoolInvite	true	"School infos"
//	@Success		201		{object}	models.SchoolInvite
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/schools/{id}/invite [post]
func (u *schoolHandlers) Invite() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var body models.SchoolInvite

		schoolInvite, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		school := &models.SchoolInvite{
			Firstname: schoolInvite.Firstname,
			Lastname:  schoolInvite.Lastname,
			Email:     schoolInvite.Email,
			SchoolId:  uint(idInt),
			Type:      schoolInvite.Type,
		}

		invitedUser, err := u.schoolUseCase.Invite(user, school)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		emailM := email.InitEmailManager(u.cfg.Smtp.Username, u.cfg.Smtp.Password, u.cfg.Smtp.Host)
		err = emailM.SendInvitationEmail([]string{invitedUser.Email}, invitedUser.Firstname, invitedUser.Lastname, invitedUser.InvitationCode)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, invitedUser)
	}
}

// Read
//
//	@Summary		Get by user
//	@Description	Get by user
//	@Tags			School
//	@Produce		json
//	@Success		200	{object}	models.School
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schools [get]
func (u *schoolHandlers) GetByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
      u.logger.Warnf("Request: Unauthorized")
			return
		}

    var school *models.School

    if user.UserKind == models.ADMINISTRATOR {
		  school, err = u.schoolUseCase.GetByUser(user)
      if err != nil {
        ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
        u.logger.Warnf("Request: %v", err.Error())
        return
      }
    } else if user.UserKind == models.STUDENT {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
      //implement getting school for student
    } else {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
      u.logger.Warnf("Request: Unauthorized")
			return
    }

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, school)
	}
}

// Read
//
//	@Summary		Get school by id
//	@Description	Get school by id
//	@Tags			School
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.School
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schools/{id} [get]
func (u *schoolHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.STUDENT)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		school, err := u.schoolUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, school)
	}
}

// Get students
//
//	@Summary		Get school students
//	@Description	Get school students
//	@Tags			School
//	@Produce		json
//	@Success		200	{object}	[]models.User
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/schools/students [get]
func (u *schoolHandlers) GetSchoolStudents() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		students, err := u.schoolUseCase.GetSchoolStudents(user.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, students)
	}
}

// Delete
//
//	@Summary		Delete school by id
//	@Description	Delete school by id
//	@Tags			School
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/schools/{id} [delete]
func (u *schoolHandlers) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		err = u.schoolUseCase.Delete(user, uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
