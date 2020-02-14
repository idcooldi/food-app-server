package interfaces

import "food-app/domain/entity"

var (
	saveUserApp func(*entity.User) (*entity.User, map[string]string)
	getUsersApp func() ([]entity.User, error)
	getUserApp func(uint64) (*entity.User, error)
	getUserEmailPasswordApp func(*entity.User) (*entity.User, map[string]string)

	//redisCreateAuth
	saveFoodApp func(*entity.Food) (*entity.Food, error)
	//getUsersApp func() ([]entity.User, error)
	//getUserApp func(uint64) (*entity.User, error)
	//getUserEmailPasswordApp func(*entity.User) (*entity.User, map[string]string)
)

type fakeUserApp struct {}
type fakeFoodApp struct {}


func (fa *fakeUserApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return getUserEmailPasswordApp(user)
}

func (fa *fakeUserApp) GetUsers() ([]entity.User, error) {
	return getUsersApp()
}
func (fa *fakeUserApp) GetUser(userId uint64) (*entity.User, error) {
	return getUserApp(userId)
}
func (fa *fakeUserApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return saveUserApp(user)
}


func (f *fakeFoodApp) SaveFood(food *entity.Food) (*entity.Food, error) {
	return saveFoodApp(food)
}

func (f *fakeFoodApp) GetAllFood() ([]entity.Food, error) {
	panic("implement me")
}

func (f *fakeFoodApp) GetFood(uint64) (*entity.Food, error) {
	panic("implement me")
}

func (f *fakeFoodApp) UpdateFood(*entity.Food) (*entity.Food, error) {
	panic("implement me")
}

func (f *fakeFoodApp) DeleteFood(uint64) error {
	panic("implement me")
}