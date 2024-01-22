package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Employee struct {
	Name    string
	Manager *Employee
}

func toEmployeeString(e *Employee) string {
	if e == nil {
		return ""
	}
	return e.Name + "/" + toEmployeeString(e.Manager)
}

func (e *Employee) String() string {
	return toEmployeeString(e)
}

type Organization struct {
	employees      []*Employee
	employeesAsMap map[string]*Employee
}

func NewOrganization() *Organization {
	return &Organization{
		employeesAsMap: make(map[string]*Employee),
	}
}

func (o *Organization) String() string {
	var employeeStrings []string
	for _, employee := range o.employeesAsMap {
		employeeStrings = append(employeeStrings, employee.String())
	}
	slices.Sort(employeeStrings)
	result := ""
	for _, employee := range employeeStrings {
		result = result + fmt.Sprintf("%s\n", employee)
	}
	return result
}

func (o *Organization) CreateEmployee(name string, manager *Employee) *Employee {
	employee, found := o.employeesAsMap[name]
	if !found {
		employee := &Employee{
			Name:    name,
			Manager: manager,
		}
		o.employeesAsMap[name] = employee
		o.employees = append(o.employees, employee)
		return employee
	}
	if manager != nil {
		employee.Manager = manager
	}

	return employee
}

func readOrganization(reader *bufio.Reader) *Organization {
	org := NewOrganization()

	numManagers := readInputInt(reader)
	for j := 0; j < numManagers; j++ {
		managerName := readInputString(reader)
		numEmployees := readInputInt(reader)
		manager := org.CreateEmployee(managerName, nil)
		for k := 0; k < numEmployees; k++ {
			employeeName := readInputString(reader)
			org.CreateEmployee(employeeName, manager)
		}
	}
	return org
}

func readInputString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func readInputInt(reader *bufio.Reader) int {
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if err != nil {
		panic(err)
	}
	textInput, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	return textInput
}

func main() {
	var reader *bufio.Reader
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		f, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		reader = bufio.NewReader(f)
	} else {
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Enter number of test cases: ")
	}
	tc := readInputInt(reader)
	for i := 0; i < tc; i++ {
		org := readOrganization(reader)
		fmt.Println("===================================")
		fmt.Println(org)
	}
}
