package main

// Section type specifies details relevant to each section of a class
// here i am using struct for grouping data together to form records.
type Section struct {
	department     string   // name of the department
	course_number  string   // course number of the subject
	section_number string   // section number of the subject
	student_names  []string // slice of students enrolled in a particular section
}

//AlreadyEnrolled is used for connecting a student using name to a specific section
type AlreadyEnrolled struct {
	student_name string   // name of the student
	section      *Section // pointer connection to the struct section to connect student name with the section enrolled
}

//creating an interface for AlreadyEnrolled to display the error as a statment for better understing in case of any error occurs
func (student AlreadyEnrolled) Error() string {
	return student.student_name + " is already in " + student.section.department + " " + student.section.course_number + "-" + student.section.section_number
}

// function to add and check the student name
// the function will intially check wether the student name is present in the list or not
// if the name is already in the list an error message is displayed
// if the name is new it is added to the list of the students enrolled

func (sec *Section) register(name string) (bool, error) {
	//checking wether the given student name is already present in the list or not
	for _, j := range sec.student_names {
		//if the student name is already in the section list then return false as a error statement
		if j == name {
			return false, AlreadyEnrolled{name, sec}
		}
	}
	//if the student name is not present in the section then add it to the list and return as given
	sec.student_names = append(sec.student_names, name)
	return true, nil
}

// function to return all the student names for the particular section if the program is ended
func (sec *Section) Students() string {
	names := ""
	// print all the student names in the list
	for i, name := range sec.student_names {
		_ = i
		names = names + name + "\n"

	}
	return names
}
