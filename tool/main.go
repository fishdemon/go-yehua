package main

import (
	"fmt"
	"strings"
)

type TypeDescription struct {
	Name string
	Fields []*TypeField
}

type TypeField struct {
	Name string
	Type string
	Tag  string
}

var data = `
type User struct {
	Id      int32        
	Name    string         
	Email   string         
	Phone   string         
	Parents []*User_Parent 
}
`

func StructToJsonModel(data string) *TypeDescription {
	strs := strings.Split(data, "\n")
	td := &TypeDescription{}
	for _, row := range strs {
		row = strings.TrimSpace(row)
		if row == "" || row =="}" {
			continue
		}

		if strings.Contains(row, "type") {
			begin := 4
			if strings.Contains(row, "struct") {
				end := strings.Index(row, "struct")
				name := strings.TrimSpace(string([]byte(row)[begin:end]))
				td.Name = name
				fmt.Println(name)
				continue
			}
		}

		fields := strings.Fields(row)
		tag := ""
		if len(fields) >= 3 {
			tag = fields[2]
		}
		field := &TypeField{
			Name: fields[0],
			Type: fields[1],
			Tag:  tag,
		}
		td.Fields = append(td.Fields, field)
	}

	return td
}


func main() {
	fmt.Println(data)
	StructToJsonModel(data)
	
}
