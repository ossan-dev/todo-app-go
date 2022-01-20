package crud

// func TestSavingTodosAndRetrievingThem(t *testing.T) {
// 	store := NewInMemoryTodoStore()
// 	server := TodoServer{store}
// 	todo := "Example"

// 	server.ServeHTTP(httptest.NewRecorder(), newPostTodoReq(todo))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostTodoReq(todo))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostTodoReq(todo))

// 	response := httptest.NewRecorder()
// 	server.ServeHTTP(response, newGetToDoByIdReq(1))
// 	util.AssertStatusCode(t, response.Code, http.StatusOK)
// 	util.AssertResponseBody(t, response.Body.String(), "Example")

// 	response = httptest.NewRecorder()
// 	server.ServeHTTP(response, newGetToDoByIdReq(2))
// 	util.AssertStatusCode(t, response.Code, http.StatusOK)
// 	util.AssertResponseBody(t, response.Body.String(), "Example")

// 	response = httptest.NewRecorder()
// 	server.ServeHTTP(response, newGetToDoByIdReq(3))
// 	util.AssertStatusCode(t, response.Code, http.StatusOK)
// 	util.AssertResponseBody(t, response.Body.String(), "Example")
// }
