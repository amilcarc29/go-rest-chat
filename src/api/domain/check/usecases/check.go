package usecases

func (usecases *UseCases) Check() (bool, error) {
	return usecases.checkRepository.Check()
}
