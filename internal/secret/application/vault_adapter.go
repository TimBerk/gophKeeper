package application

// VaultAdapter - структура для работы с нескольки UseCase
type VaultAdapter struct {
	AddUC  *AddUseCase  // use-case “добавить”
	ListUC *ListUseCase // use-case “получить список”
}

// Add - адаптация метода под выполенение в UseCase
func (v *VaultAdapter) Add(uid, typ string, data []byte, meta map[string]string) error {
	return v.AddUC.Exec(uid, typ, data, meta)
}

// List - адаптация метода под выполенение в UseCase
func (v *VaultAdapter) List(uid string) (any, error) {
	return v.ListUC.Exec(uid)
}
