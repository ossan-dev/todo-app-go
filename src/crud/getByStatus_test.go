package crud

// func TestGetByStatus(t *testing.T) {
// 	t.Run("get completed todos return non-empty collections", func(t *testing.T) {
// 		todos := []model.Todo{
// 			model.NewTodo(1, "FirstTodo", false),
// 			model.NewTodo(2, "SecondTodo", false),
// 			model.NewTodo(3, "ThirdTodo", true),
// 		}

// 		todoManager := &TodoManager{todos}

// 		got := todoManager.GetByStatus(true)
// 		want := []model.Todo{
// 			model.NewTodo(3, "ThirdTodo", true),
// 		}
// 		util.AssertCollectionsEqual(t, got, want)
// 	})

// 	t.Run("get completed todos return empty collections", func(t *testing.T) {
// 		todos := []model.Todo{
// 			model.NewTodo(1, "FirstTodo", false),
// 			model.NewTodo(2, "SecondTodo", false),
// 			model.NewTodo(3, "ThirdTodo", false),
// 		}

// 		todoManager := &TodoManager{todos}

// 		got := todoManager.GetByStatus(true)
// 		util.AssertCollectionsEqual(t, got, []model.Todo{})
// 	})
// }
