package http

import (
	"net/http"
	"strconv"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type courseHandlers struct {
	cfg           *config.Config
	courseUseCase course.UseCase
	schoolUseCase school.UseCase
	logger        logger.Logger
}

func NewCourseHandlers(cfg *config.Config, courseUseCase course.UseCase, schoolUseCase school.UseCase, logger logger.Logger) course.Handlers {
	return &courseHandlers{cfg: cfg, courseUseCase: courseUseCase, schoolUseCase: schoolUseCase, logger: logger}
}

// Create
//
//	@Summary		Create new course
//	@Description	create new course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			course	body		models.CourseCreate	true	"Course infos"
//	@Success		201		{object}	models.Course
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/courses [post]
func (u *courseHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.ADMINISTRATOR)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		var body models.CourseCreate

		courseCreate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		course := &models.Course{
			Name:        courseCreate.Name,
			Description: courseCreate.Description,
			PathId:      *courseCreate.PathId,
			TeacherId:   *courseCreate.TeacherId,
			SchoolId:    school.ID,
		}

		courseDb, err := u.courseUseCase.Create(user, course)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		// Retrieving to get the preload
		courseWithPreload, err := u.courseUseCase.GetById(courseDb.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, courseWithPreload)
	}
}

// Read
//
//	@Summary		Get all course
//	@Description	Get all course
//	@Tags			Course
//	@Produce		json
//	@Success		200	{object}	[]models.Course
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/courses [get]
func (u *courseHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := request.ValidateRole(u.cfg.JwtSecret, ctx, models.TEACHER)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			return
		}

		var schoolId uint

		if *user.UserKind == models.ADMINISTRATOR {
			school, err := u.schoolUseCase.GetByUser(user)
			if err != nil {
				ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
				u.logger.Infof("Request: %v", err.Error())
				return
			}
			schoolId = school.ID
		} else {
			schoolId = *user.SchoolId
		}

		courses, err := u.courseUseCase.GetAllBySchoolId(schoolId)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, courses)
	}
}

// Read
//
//	@Summary		Get course by id
//	@Description	Get course by id
//	@Tags			Course
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	models.Course
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		500	{object}	errorHandler.HttpErr
//	@Router			/courses/{id} [get]
func (u *courseHandlers) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.UrlParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		course, err := u.courseUseCase.GetById(uint(idInt))

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, course)
	}
}

// Update
//
//	@Summary		Update course
//	@Description	Update course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"id"
//	@Param			course	body		models.CourseUpdate	true	"Course infos"
//	@Success		201		{object}	models.Course
//	@Failure		400		{object}	errorHandler.HttpErr
//	@Failure		500		{object}	errorHandler.HttpErr
//	@Router			/courses/{id} [put]
func (u *courseHandlers) Update() gin.HandlerFunc {
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

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		var body models.CourseUpdate

		courseUpdate, err := request.ValidateJSON(body, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.BodyParamsErrorResponse())
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		courseDb, err := u.courseUseCase.GetById(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		if courseDb.SchoolId != school.ID {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: Not allowed to update course not on your school")
			return
		}

		course := &models.Course{
			Name:        courseUpdate.Name,
			Description: courseUpdate.Description,
			PathId:      *courseUpdate.PathId,
			TeacherId:   *courseUpdate.TeacherId,
			SchoolId:    courseDb.SchoolId,
		}
		courseDb, err = u.courseUseCase.Update(uint(idInt), course)

		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		// Retrieving to get the preload
		courseWithPreload, err := u.courseUseCase.GetById(courseDb.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, courseWithPreload)
	}
}

// Delete
//
//	@Summary		Delete course by id
//	@Description	Delete course by id
//	@Tags			Course
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200	{object}	nil
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/courses/{id} [delete]
func (u *courseHandlers) Delete() gin.HandlerFunc {
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

		school, err := u.schoolUseCase.GetByUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		courseDb, err := u.courseUseCase.GetById(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		if courseDb.SchoolId != school.ID {
			ctx.AbortWithStatusJSON(errorHandler.UnauthorizedErrorResponse())
			u.logger.Infof("Request: Not allowed to delete course not on your school")
			return
		}

		err = u.courseUseCase.Delete(uint(idInt))
		if err != nil {
			ctx.AbortWithStatusJSON(errorHandler.ErrorResponse(err))
			u.logger.Infof("Request: %v", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
