package todos

// func TestNewTodoSvc(t *testing.T) {
// 	testHost := "test_json_host_name"
// 	assert.Equal(t, NewTodoSvc(testHost), &TodoSvc{hostname: testHost})
// }

// func TestGetTodos(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		host  string
// 		todos []Todo
// 		err   error
// 	}{
// 		{
// 			name: "Invalid schema",
// 			host: "test_json_host",
// 			err:  fmt.Errorf("Get \"/todos\": unsupported protocol scheme \"\""),
// 		},
// 	}
// 	for _, tt := range tests {
// 		todoSvc := NewTodoSvc(tt.host)
// 		t.Run(tt.name, func(t *testing.T) {
// 			todos, err := todoSvc.GetTodos()
// 			if tt.err == nil {
// 				assert.Equal(t, tt.todos, todos)
// 			} else {
// 				if err == nil {
// 					t.Error("Expected error, got nil")
// 				} else {
// 					assert.Equal(t, tt.err.Error(), err.Error())
// 				}
// 			}
// 		})
// 	}
// }
