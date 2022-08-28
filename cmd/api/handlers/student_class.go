package handlers

import (
	"github.com/gin-gonic/gin"
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
}

func NewStudentClassHandler(db *gorm.DB, accountService account_service.AccountService) StudentClassHandler {
	return StudentClassHandler{
		Service: service.StudentClassService{
			Db: db,
		},
		AccService: accountService,
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
	auth, err := h.AccService.Auth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req request.StudentClassRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if auth.ID != req.StudentID && auth.Role != account_service.TeacherRole && auth.Role != account_service.AdminRole {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	studentClass, err := req.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.AccService.ExistsStudent(token, studentClass.StudentID); err != nil {
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
	auth, err := h.AccService.Auth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req request.StudentClassRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if auth.ID != req.StudentID && auth.Role != account_service.TeacherRole && auth.Role != account_service.AdminRole {
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
