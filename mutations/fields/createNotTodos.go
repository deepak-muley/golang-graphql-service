package fields

import (
	"context"

	"github.com/graphql-go/graphql"

	mongo "golang-graphql-service/data"
	types "golang-graphql-service/types"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var CreateNotTodo = &graphql.Field{
	Type:        types.NotTodo,
	Description: "Create a not Todo",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// get our params
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)
		notTodoCollection := mongo.Client.GetInternalClient().Database("medium-app").Collection("Not_Todos")
		_, err := notTodoCollection.InsertOne(context.Background(), map[string]string{"name": name, "description": description})
		if err != nil {
			panic(err)
		}
		return todoStruct{name, description}, nil
	},
}
