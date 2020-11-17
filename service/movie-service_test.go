package service

import (
	"github.com/matiascfgm/Go-rest-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(movie *entity.Movie) (*entity.Movie, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Movie), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Movie, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Movie), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	movie := entity.Movie{
		ID:          "",
		Title:       "Fast And Furious",
		Description: "Fast cars with over 20 gears",
		Year:        "2001",
	}
	// expectations
	mockRepo.On("FindAll").Return([]entity.Movie{movie}, nil)

	testService := NewMovieService(mockRepo)
	result, _ := testService.FindAll()
	// mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//	 Data assertion
	assert.Equal(t, "Fast And Furious", result[0].Title)
	assert.Equal(t, "Fast cars with over 20 gears", result[0].Description)
	assert.Equal(t, "2001", result[0].Year)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	movie := entity.Movie{
		ID:          "420",
		Title:       "Fast And Furious",
		Description: "Fast cars with over 20 gears",
		Year:        "2001",
	}
	// expectations
	mockRepo.On("Create").Return(&movie, nil)

	testService := NewMovieService(mockRepo)
	result, err := testService.Create(&movie)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, "420", result.ID)
	assert.Equal(t, "Fast And Furious", result.Title)
	assert.Equal(t, "Fast cars with over 20 gears", result.Description)
	assert.Equal(t, "2001", result.Year)
	assert.Nil(t, err)

}

func TestValidateEmptyMovie(t *testing.T) {
	testService := NewMovieService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The movie is empty", err.Error())
}

func TestValidateEmptyMovieTitle(t *testing.T) {
	movie := entity.Movie{
		ID:          "",
		Title:       "",
		Description: "Fast cars with over 20 gears",
		Year:        "2001",
	}
	testService := NewMovieService(nil)
	err := testService.Validate(&movie)
	assert.NotNil(t, err)
	assert.Equal(t, "The movie title is empty", err.Error())
}
