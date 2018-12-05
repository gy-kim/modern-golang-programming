package databaselayer

type MongodbHandler struct{}

func NewMongodbHandler() *MongodbHandler {
	return nil
}

func (handler *MongodbHandler) GetAvailableDynos() ([]Animal, error) {
	return nil, nil
}

func (handler *MongodbHandler) GetDynoByNickname(string) (Animal, error) {
	return Animal{}, nil
}
func (handler *MongodbHandler) GetDynosByType(string) (Animal, error) {
	return Animal{}, nil
}
func (handler *MongodbHandler) AddAnimal(Animal) error {
	return nil
}
func (handler *MongodbHandler) UpdateAnumal(Animal, string) error {
	return nil
}
