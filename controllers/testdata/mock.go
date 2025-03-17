package testdata

import "github.com/h-hiwatashi/go-todo-clean-api/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostTodoService(todo models.Todo) (models.Todo, error) {
	return todoTestData[1], nil
}

func (s *serviceMock) GetTodoListService(page int) ([]models.Todo, error) {
	return todoTestData, nil
}

func (s *serviceMock) GetTodoService(todoID int) (models.Todo, error) {
	return todoTestData[0], nil
}

func (s *serviceMock) PostNiceService(todo models.Todo) (models.Todo, error) {
	return todoTestData[0], nil
}
func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}
