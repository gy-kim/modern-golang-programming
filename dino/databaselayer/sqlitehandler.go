package databaselayer

type SQLiteHandler struct{}

func NewSQLiteHandler() *SQLiteHandler {
	return nil
}

func (handler *SQLiteHandler) GetAvailableDynos() ([]Animal, error) {
	return nil, nil
}

func (handler *SQLiteHandler) GetDynoByNickname(string) (Animal, error) {
	return Animal{}, nil
}
func (handler *SQLiteHandler) GetDynosByType(string) (Animal, error) {
	return Animal{}, nil
}
func (handler *SQLiteHandler) AddAnimal(Animal) error {
	return nil
}
func (handler *SQLiteHandler) UpdateAnumal(Animal, string) error {
	return nil
}
