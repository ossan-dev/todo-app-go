package crud

// func TestGetAllEndpoint(t *testing.T) {
// 	server := &TodoServer{}

// 	t.Run("returns todos", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "api/todos", nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		got := response.Body.String()
// 		want := `[{"id": 1, "description": "FirstTodo", "isCompleted": false}]`

// 		if got != want {
// 			t.Errorf("got %q but want %q", got, want)
// 		}
// 	})
// }
