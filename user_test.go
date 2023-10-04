package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestInviteUser(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserInvite($email:String!$input:UserInput!){userInvite(email: $email input: $input){user{id,email,htmlUrl,name,role},errors{message,path}}}",
	"variables":{
		"email": "kyle@opslevel.com",
		"input": {
			"name": "Kyle Rockman",
			"skipWelcomeEmail": false
		}
	}}`
	response := `{"data": {
	"userInvite": {
		"user": {{ template "user_1" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "user/invite", request, response)
	// Act
	result, err := client.InviteUser("kyle@opslevel.com", ol.UserInput{
		Name: "Kyle Rockman",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1", string(result.Id))
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestGetUser(t *testing.T) {
	// Arrange
	request := `{"query":"query UserGet($input:UserIdentifierInput!){account{user(input: $input){id,email,htmlUrl,name,role}}}",
	"variables":{
		"input": {
			"email": "kyle@opslevel.com"
		}
	}}`
	response := `{"data": {
	"account": {
		"user": {{ template "user_1" }}
	}
	}}`
	client := ABetterTestClient(t, "user/get", request, response)
	// Act
	result, err := client.GetUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1", string(result.Id))
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestGetUserTeams(t *testing.T) {
	// Arrange
	testRequestOne := TestRequest{
		Request:   `"query": "query UserTeamsList($after:String!$first:Int!$user:ID!){account{user(id: $user){teams(after: $after, first: $first){nodes{alias,id},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		Variables: `"variables": { {{ template "first_page_variables" }}, "user": "{{ template "id1"}}" }`,
		Response: `{ "data": {
                    "account": {
                        "user": {
                            "teams": {
                                "nodes": [
                                  {{ template "teamId_1"}},
                                  {{ template "teamId_2"}}
                                ],
                              {{ template "pagination_initial_pageInfo_response" }},
                              "totalCount": 2
                  }}}}}`,
	}
	testRequestTwo := TestRequest{
		Request:   `"query": "query UserTeamsList($after:String!$first:Int!$user:ID!){account{user(id: $user){teams(after: $after, first: $first){nodes{alias,id},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		Variables: `"variables": { {{ template "second_page_variables" }}, "user": "{{ template "id1"}}" }`,
		Response: `{ "data": {
                    "account": {
                        "user": {
                            "teams": {
                                "nodes": [
                                  {{ template "teamId_3"}}
                                ],
                                {{ template "pagination_second_pageInfo_response" }},
                                "totalCount": 1
                            }}}}}`,
	}
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := TmpPaginatedTestClient(t, "user/teams", requests...)
	// Act
	user := ol.User{
		UserId: ol.UserId{
			Id: "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		},
	}
	resp, err := user.Teams(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, "example_3", result[2].Alias)
}

func TestListUser(t *testing.T) {
	// Arrange
	testRequestOne := TestRequest{
		Request:   `"query": "query UserList($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		Variables: `{{ template "pagination_initial_query_variables" }}`,
		Response:  `{ "data": { "account": { "users": { "nodes": [ {{ template "user_1" }}, {{ template "user_2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	}
	testRequestTwo := TestRequest{
		Request:   `"query": "query UserList($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		Variables: `{{ template "pagination_second_query_variables" }}`,
		Response:  `{ "data": { "account": { "users": { "nodes": [ {{ template "user_3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	}
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := TmpPaginatedTestClient(t, "user/list", requests...)
	// Act
	response, err := client.ListUsers(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Edgar Ochoa", result[1].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[1].Role)
	autopilot.Equals(t, "Matthew Brahms", result[2].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[2].Role)
}

func TestUpdateUser(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserUpdate($input:UserInput!$user:UserIdentifierInput!){userUpdate(user: $user input: $input){user{id,email,htmlUrl,name,role},errors{message,path}}}",
	"variables":{
		"input": {
			"role": "admin",
			"skipWelcomeEmail": false
		},
		"user": {
			"email": "kyle@opslevel.com"
		}
	}}`
	response := `{"data": {
	"userUpdate": {
		"user": {{ template "user_1_update" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "user/update", request, response)
	// Act
	result, err := client.UpdateUser("kyle@opslevel.com", ol.UserInput{
		Role: ol.UserRoleAdmin,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1", string(result.Id))
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result.Role)
}

func TestDeleteUser(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserDelete($user:UserIdentifierInput!){userDelete(user: $user){errors{message,path}}}",
	"variables":{
		"user": {
			"email": "kyle@opslevel.com"
		}
	}}`
	response := `{"data": {
	"userDelete": {
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "user/delete", request, response)
	// Act
	err := client.DeleteUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteUserDoesNotExist(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserDelete($user:UserIdentifierInput!){userDelete(user: $user){errors{message,path}}}",
	"variables":{
		"user": {
			"email": "not-found@opslevel.com"
		}
	}}`
	response := `{"data": {
	"userDelete": {
		"errors": [{
			"message": "User with email 'not-found@opslevel.com' does not exist on this account",
			"path": ["user"]
		}]
	}
	}}`
	client := ABetterTestClient(t, "user/delete_not_found", request, response)
	// Act
	err := client.DeleteUser("not-found@opslevel.com")
	// Assert
	autopilot.Equals(t, `OpsLevel API Errors:
	- 'user' User with email 'not-found@opslevel.com' does not exist on this account
`, err.Error())
}

func TestGetUserTags(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			Request: `{"query": "query UserTagsList($after:String!$first:Int!$user:ID!){account{user(id: $user){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
            "variables": {
              {{ template "first_page_variables" }},
              "user": "{{ template "id1"}}"
            }}`,
			Response: `{
        "data": {
          "account": {
            "user": {
              "tags": {
                "nodes": [
                  {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIw",
                    "key": "user-tag-1",
                    "value": "user-value-1"
                  },
                  {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIx",
                    "key": "user-tag-2",
                    "value": "user-value-2"
                  },
                  {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIy",
                    "key": "db",
                    "value": "prod"
                  }
                ],
                {{ template "pagination_initial_pageInfo_response" }},
                "totalCount": 3
              }
            }
          }
        }
      }`,
		},
		{
			Request: `{"query": "query UserTagsList($after:String!$first:Int!$user:ID!){account{user(id: $user){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
            "variables": {
                {{ template "second_page_variables" }},
                "user": "{{ template "id1"}}"
            }
            }`,
			Response: `{
        "data": {
          "account": {
            "user": {
              "tags": {
                "nodes": [
                {
                  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIz",
                  "key": "captain",
                  "value": "tuna"
                }
                ],
                {{ template "pagination_second_pageInfo_response" }},
                "totalCount": 1
              }
          }
          }}}`,
		},
	}
	client := APaginatedTestClient(t, "user/tags", requests...)
	// Act
	user := ol.User{
		UserId: ol.UserId{
			Id: "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		},
	}
	resp, err := user.GetTags(client, nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, resp.TotalCount)
	autopilot.Equals(t, "user-tag-1", result[0].Key)
	autopilot.Equals(t, "user-value-1", result[0].Value)
	autopilot.Equals(t, "user-tag-2", result[1].Key)
	autopilot.Equals(t, "user-value-2", result[1].Value)
	autopilot.Equals(t, "db", result[2].Key)
	autopilot.Equals(t, "prod", result[2].Value)
	autopilot.Equals(t, "captain", result[3].Key)
	autopilot.Equals(t, "tuna", result[3].Value)
}
