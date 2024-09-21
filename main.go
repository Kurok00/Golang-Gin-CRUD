package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	// Email string `json:"email"`
}

var students []Student
var nextID = 1

func main() {
    r := gin.Default()
    
	// Route để lấy danh sách sinh viên
	r.GET("/get-students", getStudents)

	// Route để lấy thông tin chi tiết một sinh viên
	r.GET("/get-student-detail/:id", getStudentDetail)

	// Route để thêm sinh viên
	r.POST("/add-student", addStudent )

	// Route để cập nhật sinh viên
	r.PUT("/update-student/:id", updateStudent)

	// Route để xóa sinh viên
	r.DELETE("/delete-student/:id", deleteStudent)
    
	fmt.Println("API is running at: http://localhost:8080/hello")

    r.Run(":8080")
}

// Hàm lấy danh sách sinh viên
func getStudents(c *gin.Context){
	c.JSON(http.StatusOK, students)
}

// Hàm lấy chi tiết một sinh viên
func getStudentDetail(c *gin.Context){
	id := c.Param("id")
	for _, student := range students{
		if fmt.Sprint(student.ID) == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Studen not found !"})
}

// Hàm thêm sinh viên
func addStudent(c *gin.Context){
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err == nil {
		newStudent.ID = nextID
		nextID ++
		students =  append(students, newStudent)
		c.JSON(http.StatusCreated, newStudent)
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "loi"})
	}
	
}

// Hàm cập nhật sinh viên
func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var updatedStudent Student

	if err := c.ShouldBindJSON(&updatedStudent); err == nil {
		for i, student := range students {
			if fmt.Sprint(student.ID) == id {
				students[i].Name = updatedStudent.Name
				students[i].Age = updatedStudent.Age
				c.JSON(http.StatusOK, students[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Hàm xóa sinh viên
func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	for i, student := range students {
		if fmt.Sprint(student.ID) == id {
			students = append(students[:i], students[i+1:]...) // Xóa sinh viên
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}



