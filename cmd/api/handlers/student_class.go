package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"main/cmd/api/request"
	"main/cmd/api/response"
	"main/internal/account_service"
	"main/internal/service"
	"net/http"
)

type StudentClassHandler struct {
	Service    service.StudentClassService
	AccService account_service.AccountService
	validator  *validator.Validate
	SkipAuth   bool
}

func NewStudentClassHandler(db *gorm.DB, accountService account_service.AccountService, _validator *validator.Validate, SkipAuth bool) StudentClassHandler {
	return StudentClassHandler{
		Service: service.StudentClassService{
			Db: db,
		},
		AccService: accountService,
		validator:  _validator,
		SkipAuth:   SkipAuth,
	}
}

// @Summary Create StudentClass
// @Tags         student-classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        class  body      request.StudentClassRequest  true  "Add studentClass"
// @Success 201 {object} response.StudentClassResponse
// @Router /student-classes [post]
func (h *StudentClassHandler) CreateHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	auth, err := h.AccService.Auth(token, h.SkipAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req request.StudentClassRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !auth.Role.IsRoleAuthorized([]account_service.Role{account_service.TeacherRole}, h.SkipAuth) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	studentClass, err := req.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.AccService.ExistsStudent(token, studentClass.StudentID, h.SkipAuth); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := h.Service.Create(studentClass); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.NewStudentClassResponse(studentClass))
}

// @Summary Delete StudentClass
// @Tags         student-classes
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        class  body      request.StudentClassRequest  true  "Delete studentClass"
// @Success 204 {object} response.StudentClassResponse
// @Router /student-classes [delete]
func (h *StudentClassHandler) DeleteHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	auth, err := h.AccService.Auth(token, h.SkipAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req request.StudentClassRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !auth.Role.IsRoleAuthorized([]account_service.Role{account_service.TeacherRole}, h.SkipAuth) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	studentClass, err := req.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.Service.Delete(studentClass.StudentID, studentClass.ClassID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, 0)
}
