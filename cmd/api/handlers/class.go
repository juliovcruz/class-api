package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main/cmd/api/request"
	"main/cmd/api/response"
	"main/internal/account_service"
	"main/internal/service"
	"net/http"
)

type ClassHandler struct {
	Service             service.ClassService
	AccService          account_service.AccountService
	StudentClassService service.StudentClassService
	validator           *validator.Validate
}

func NewClassHandler(db *gorm.DB, accountService account_service.AccountService, validator *validator.Validate) ClassHandler {
	return ClassHandler{
		Service: service.ClassService{
			Db: db,
		},
		AccService: accountService,
		validator:  validator,
		StudentClassService: service.StudentClassService{
			Db: db,
		},
	}
}

// @Summary Create Class
// @Description Only the class teacher can create it
// @Tags         classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        class  body      request.ClassRequest  true  "Add class"
// @Success 201 {object} response.ClassResponse
// @Router /classes [post]
func (h *ClassHandler) CreateHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	auth, err := h.AccService.Auth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req request.ClassRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if auth.ID != req.TeacherID && !auth.Role.IsRoleAuthorized([]account_service.Role{account_service.TeacherRole}) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	class, err := req.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.AccService.ExistsTeacher(token, class.TeacherID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.Service.Create(class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.NewClassResponse(class))
}

// @Summary Update Class
// @Tags         classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        class  body      request.ClassRequest  true  "Update class"
// @Param        id   path      string  true  "ClassID"
// @Success 200 {object} response.ClassResponse
// @Router /classes/{id} [put]
func (h *ClassHandler) UpdateHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	auth, err := h.AccService.Auth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var req request.ClassRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if !auth.Role.IsRoleAuthorized([]account_service.Role{account_service.TeacherRole}) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	class, err := req.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	class.ID = id

	if err := h.Service.Update(class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewClassResponse(class))
}

// @Summary Get Class By ID
// @Tags         classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        id   path      string  true  "ClassID"
// @Success 200 {object} response.ClassResponse
// @Router /classes/{id} [get]
func (h *ClassHandler) GetByIDHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if _, err := h.AccService.Auth(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	class, err := h.Service.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewClassResponse(class))
}

// @Summary Get All Classes
// @Tags         classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {array} response.ClassResponse
// @Router /classes [get]
func (h *ClassHandler) GetAllHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if _, err := h.AccService.Auth(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	classes, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewClassesResponse(classes))
}

// @Summary Get All Student Classes By Class
// @Tags         classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        id   path      string  true  "ClassID"
// @Success 200 {array} response.StudentClassResponse
// @Router /classes/{id}/student-classes [get]
func (h *ClassHandler) GetAllStudentClassesByHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if _, err := h.AccService.Auth(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	studentsClass, err := h.StudentClassService.GetAllByClassID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewStudentsClassResponse(studentsClass))
}

// @Summary Delete Class
// @Tags         classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        id   path      string  true  "ClassID"
// @Success 204 {object} response.ClassResponse
// @Router /classes/{id} [delete]
func (h *ClassHandler) DeleteHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	auth, err := h.AccService.Auth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	class, err := h.Service.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if auth.ID != class.TeacherID.String() &&
		!auth.Role.IsRoleAuthorized([]account_service.Role{account_service.TeacherRole}) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if err := h.Service.Delete(class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, 0)
}
