package crud

// func TestAddTodo(t *testing.T) {
// 	store := StubTodoStore{
// 		map[int]model.Todo{},
// 	}

// 	server := &TodoServer{&store}

// 	t.Run("it adds todo on POST", func(t *testing.T) {
// 		todoDescription := "Example"

// 		request := newPostTodoReq(todoDescription)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		util.AssertStatusCode(t, response.Code, http.StatusAccepted)

// 		if len(store.todos) != 1 {
// 			t.Errorf("got %d calls to AddTodo but want %d", len(store.todos), 1)
// 		}

// 		if store.todos[1].Description != todoDescription {
// 			t.Errorf("did not store correct todo got %q but want %q", store.todos[1].Description, todoDescription)
// 		}
// 	})
// }

// func newPostTodoReq(description string) *http.Request {
// 	req, _ := http.NewRequest(http.MethodPost, "/api/todos", strings.NewReader(description))
// 	return req
// }
