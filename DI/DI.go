package DI

type ServiceImpl struct {}

type OrderService interface {
	Apply(int) error
}

type Application struct {
	os OrderService
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

//todo objectInitializeDI
func (s *ServiceImpl) Apply(id int) error {
	return nil
}

func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}

//todo setterDI
func (app *Application) SetService(os OrderService) {
	app.os = os
}

//todo calledMthodDI
// func (app *Application) Apply(os OrderService, id int) error {
// 	return os.Apply(id)
// }

func main() {
	//todo objectInitializeDI
	app := NewApplication(&ServiceImpl{})
	app.Apply(19)

	//todo setterDI
	appset := &Application{}
	appset.SetService(&ServiceImpl{})
	app.Apply(19)

	//todo calledMthodDI
	// appMethod := &Application{}
	// appMethod.Apply(&ServiceImpl{}, 19)
}

