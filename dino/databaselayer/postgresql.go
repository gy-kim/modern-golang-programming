package databaselayer

type PQHandler struct{}

func NewPQHandler() *PQHandler {
	return nil
}

func (handler *PQHandler) GetAvailableDynos() ([]Animal, error) {
	return nil, nil
}

func (handler *PQHandler) GetDynoByNickname(string) (Animal, error) {
	return Animal{}, nil
}

func (handler *PQHandler) GetDynosByType(string) (Animal, error) {
	return Animal{}, nil
}

func (handler *PQHandler) AddAnimal(Animal) error {
	return nil
}

func (handler *PQHandler) UpdateAnumal(Animal, string) error {
	return nil
}
