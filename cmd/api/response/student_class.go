package response

import (
	"github.com/google/uuid"
	"main/internal/entity"
)

type StudentClassResponse struct {
	ClassID   uuid.UUID `json:"class_id"`
	StudentID uuid.UUID `json:"student_id"`
}

func NewStudentClassResponse(studentClass entity.StudentClass) StudentClassResponse {
	return StudentClassResponse{
		ClassID:   studentClass.ClassID,
		StudentID: studentClass.StudentID,
	}
}

func NewStudentsClassResponse(studentsClass []entity.StudentClass) []StudentClassResponse {
	res := []StudentClassResponse{}

	for _, studentClass := range studentsClass {
		res = append(res, StudentClassResponse{
			ClassID:   studentClass.ClassID,
			StudentID: studentClass.StudentID,
		})
	}

	return res
}
