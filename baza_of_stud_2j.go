package main

import "fmt"
import "encoding/json"
import "os"
import "bufio"

type Action struct {
	Action string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID string  `json:"id"`
	Salary float64 `json:"salary"`
	Subject string `json:"subject"`
	Classroom []string `json:"classroom"`
	Person struct {
		Name string `json:"name"`
		Surname string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

var obj_sl []GeneralObject 

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

type GeneralObject interface {
	GetCreateAction() DefinedAction
	GetUpdateAction() DefinedAction
	GetReadAction() DefinedAction
	GetDeleteAction() DefinedAction
}

type CreateTeacher struct {
	T Teacher `json:"data"`
}
func (action *CreateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateTeacher) Process() {
	obj_sl = append(obj_sl, action.T)
}

type UpdateTeacher struct {
	T Teacher `json:"data"`
}
func (action *UpdateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateTeacher) Process() {
	
}

type ReadTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
func (action *ReadTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadTeacher) Process() {
	 for i := 0; i < len(obj_sl); i++{
		 if action.Data.ID == GetId(obj_sl[i]){
			 fmt.Println(obj_sl[i])
		}
	} 
}

func GetId(o GeneralObject)string{
		switch o.(type){
			case *Teacher: return o.ID
	
		}
		return o.ID
}

type DeleteTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
func (action *DeleteTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteTeacher) Process() {
	for i := 0; i < len(obj_sl); i++{
		if GetId(obj_sl[i]) == action.Data.ID{
			obj_sl[i]= nil 
		}
	}
}

func main() {
file, _ := os.Open("data.json")
scanner := bufio.NewScanner(file)

for scanner.Scan() {
  var act Action
  err := json.Unmarshal([]byte(scanner.Text()), &act)
  
	var obj GeneralObject
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	}
	var toDo DefinedAction
	switch act.Action {
	case "create":
		toDo = obj.GetCreateAction()
		toDo.Process()
	case "update":
		toDo = obj.GetUpdateAction()
		toDo.Process()
	case "read":
		toDo = obj.GetReadAction()
		toDo.Process()
	}
//	toDo.GetFromJSON(file)
	
	toDo.Process()
}
}
