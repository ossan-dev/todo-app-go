package crud

// func TestDelete(t *testing.T) {
// 	todos := []models.Todo{
// 		models.NewTodo(1, "FirstTodo", false),
// 		models.NewTodo(2, "SecondTodo", false),
// 		models.NewTodo(3, "ThirdTodo", true),
// 	}

// 	todoManager := &TodoManager{todos}

// 	t.Run("delete todo", func(t *testing.T) {
// 		todoManager.DeleteById(2)
// 		_, err := todoManager.GetById(2)
// 		utils.AssertError(t, err, ErrTodoNotFound)
// 	})

// 	t.Run("not delete todo if not exists", func(t *testing.T) {
// 		err := todoManager.DeleteById(4)
// 		utils.AssertError(t, err, ErrTodoNotFound)
// 	})
// }
