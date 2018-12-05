package databaselayer

type MySQLHandler struct{}

func NewMySQLHandler() *MySQLHandler {
	return nil
}

func (handler *MySQLHandler) GetAvailableDynos() ([]Animal, error) {
	return nil, nil
}

func (handler *MySQLHandler) GetDynoByNickname(string) (Animal, error) {
	return Animal{}, nil
}
func (handler *MySQLHandler) GetDynosByType(string) (Animal, error) {
	return Animal{}, nil
}
func (handler *MySQLHandler) AddAnimal(Animal) error {
	return nil
}
func (handler *MySQLHandler) UpdateAnumal(Animal, string) error {
	return nil
}
