package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id    int
	Name  string
	Marks float64
}
type StudentManager struct {
	Students map[int]Student
}

// Constructor function: Map ko initialize karne ka sabse safe tareeka
func NewStudentManager() *StudentManager {
	return &StudentManager{
		Students: make(map[int]Student),
	}
}

func (sm *StudentManager) AddStudent(s Student) error {
	if s.Id <= 0 {
		return errors.New("student ID must be positive")
	}
	if strings.TrimSpace(s.Name) == "" {
		return errors.New("student name cannot be empty")

	}
	if s.Marks < 0 || s.Marks > 100 {
		return errors.New("marks must be between 0 and 100")
	}

	if _, exists := sm.Students[s.Id]; exists {
		return errors.New("student already exists")
	}

	sm.Students[s.Id] = s
	fmt.Println("Student added successfully!")
	return nil
}

func (sm *StudentManager) GetStudent() {
	// var Allstudent []Student
	fmt.Println("============ All Students =====================")
	if len(sm.Students) == 0 {
		fmt.Println("No students available.")
		return
	}
	for _, student := range sm.Students {
		// Allstudent = append(Allstudent, student)
		fmt.Printf("ID: %d | Name: %s | Marks: %.2f\n", student.Id, student.Name, student.Marks)
	}
	// return Allstudent
}

func (sm *StudentManager) UpdateStudent(Studentid int, newMarks float64, newName string) error {

	// 1. Check karein ki student exist karta hai ya nahi
	Student, exists := sm.Students[Studentid]
	if !exists {
		return errors.New("student not found")
	}
	// 2. Field update karein
	Student.Marks = newMarks
	Student.Name = newName

	// 3. Updated struct ko wapas map mein daalein
	sm.Students[Studentid] = Student
	fmt.Println("Student updated successfully!")
	return nil

}

func (sm *StudentManager) DeleteStudent(Studentid int) error {

	// 1. Check karein ki student exist karta hai ya nahi
	_, exists := sm.Students[Studentid]
	if !exists {
		return errors.New("student not found")
	}
	//Delete student  from map
	delete(sm.Students, Studentid)

	return nil
}

func (sm *StudentManager) GetTopper() (*Student, error) {
	if len(sm.Students) == 0 {
		return &Student{}, errors.New("no students available")
	}

	// Map ki saari values ko slice mein convert karke MaxFunc lagayein

	// studentSlice := slices.Collect(maps.Values(sm.Students))
	// topper := slices.MaxFunc(sm.Students, func(s1, s2 Student) int {
	// 	return cmp.Compare(s1.Marks, s2.Marks)
	// })

	var topper Student
	first := true

	for _, student := range sm.Students {
		if first || student.Marks > topper.Marks {
			topper = student
			first = false
		}
	}
	return &topper, nil

}

// Helper function terminal se safe input padhne ke liye
func readInput(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("input closed")
	}

	return strings.TrimSpace(scanner.Text()), nil
}

func main() {
	// s1 := Student{
	// 	Id:    1,
	// 	Name:  "Anubhav",
	// 	Marks: 56.9,
	// }
	// s2 := Student{
	// 	Id:    2,
	// 	Name:  "Raj",
	// 	Marks: 45.9,
	// }
	// s3 := Student{
	// 	Id:    3,
	// 	Name:  "Rohit",
	// 	Marks: 76.9,
	// }

	// AddStudent(s1)
	// AddStudent(s2)
	// AddStudent(s3)
	// GetStudent()
	// UpdateStudent(1, 20.9, "Robin")
	// GetStudent()

	// DeleteStudent(1)
	// GetStudent()
	// GetTopper()

	scanner := bufio.NewScanner(os.Stdin)
	s := NewStudentManager()

	for {

		fmt.Println("\n========= Student Management System =========")
		fmt.Println("1. Add Student")
		fmt.Println("2. View All Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Get Topper")
		fmt.Println("6. Exit")

		choiceStr, err := readInput(scanner, "Choose an option (1-6): ")
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice! Please enter a number between 1 and 6.")
			continue

		}

		switch choice {
		case 1:
			idStr, err := readInput(scanner, "Enter Student ID: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID! Must be a number.")
				continue
			}

			name, err := readInput(scanner, "Enter Student Name: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			if name == "" {
				fmt.Println("Name cannot be empty!")
				continue
			}

			marksStr, err := readInput(scanner, "Enter Student Marks: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			marks, err := strconv.ParseFloat(marksStr, 64)
			if err != nil {
				fmt.Println("Invalid Marks! Must be a number.")
				continue
			}

			err = s.AddStudent(Student{Id: id, Name: name, Marks: marks})

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

		case 2:
			s.GetStudent()

		case 3:
			idStr, err := readInput(scanner, "Enter Student ID to update: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID!")
				continue
			}

			newName, err := readInput(scanner, "Enter New Name: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			if newName == "" {
				fmt.Println("Name cannot be empty!")
				continue
			}

			marksStr, err := readInput(scanner, "Enter New Marks: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			newMarks, err := strconv.ParseFloat(marksStr, 64)
			if err != nil {
				fmt.Println("Invalid Marks!")
				continue
			}

			if err := s.UpdateStudent(id, newMarks, newName); err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			idStr, err := readInput(scanner, "Enter Student ID to delete: ")
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID!")
				continue
			}

			if err := s.DeleteStudent(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Student deleted successfully!")
			}

		case 5:
			topper, err := s.GetTopper()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Topper: ID: %d | Name: %s | Marks: %.2f\n", topper.Id, topper.Name, topper.Marks)
			}

		case 6:
			fmt.Println("Exiting program. Goodbye!")
			return

		default:
			fmt.Println("Invalid option! Please choose between 1 and 6.")
		}

	}
}
