package http

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/example"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type exampleHandlers struct {
	exampleUseCase example.UseCase
	cfg         *config.Config
	logger      logger.Logger
}

func NewExampleHandlers(exampleUseCase example.UseCase, cfg *config.Config, logger logger.Logger) example.Handlers {
	return &exampleHandlers{exampleUseCase: exampleUseCase, cfg: cfg, logger: logger}
}

// Create
//	@Summary		Create new example
//	@Description	create new example
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Param			example	body		models.Example	true	"Example infos"
//	@Success		201		{object}	models.Example
//	@Failure		400		{object}	string
//	@Failure		406		{object}	string
//	@Router			/examples [post]
func (u *exampleHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		example := &models.Example{}

    if err := ctx.BindJSON(example); err != nil {
      ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      u.logger.Errorf("Request: %v", err)
      return
    }

		example, err := u.exampleUseCase.Create(example)

    if err != nil {
      ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
      u.logger.Errorf("Request: %v", err)
      return
    }

		ctx.JSON(http.StatusCreated, example)
	}
}

// Create
//	@Summary		Get all example
//	@Description	Get all example
//	@Tags			Example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Example
//	@Failure		500	{object}	string
//	@Router			/examples [get]
func (u *exampleHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		examples, err := u.exampleUseCase.GetAll()

    if err != nil {
      ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
      u.logger.Errorf("Request: %v", err)
      return
    }

		ctx.JSON(http.StatusOK, examples)
	}
}
