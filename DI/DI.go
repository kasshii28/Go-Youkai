package DI

// ServiceImpl は OrderService インターフェースの実装
type ServiceImpl struct {}

// OrderService は注文処理のための抽象インターフェース
type OrderService interface {
    Apply(int) error
}

// Application はアプリケーションのメインロジックを保持する構造体
// os フィールドに OrderService の実装を注入可能
type Application struct {
    os OrderService
}

// コンストラクタインジェクション:
// NewApplication は OrderService を受け取り、Application を初期化
func NewApplication(os OrderService) *Application {
    return &Application{os: os}
}

// セッターインジェクション:
// SetService は実行時に OrderService の実装を切り替え可能
func (app *Application) SetService(os OrderService) {
    app.os = os
}

// Apply は注入された OrderService の Apply メソッドを実行
func (app *Application) Apply(id int) error {
    return app.os.Apply(id)
}

// ServiceImpl の Apply 実装
func (s *ServiceImpl) Apply(id int) error {
    return nil
}

// メソッドインジェクション（コメントアウト中）:
// 各メソッド呼び出し時に依存を注入する方式
// func (app *Application) Apply(os OrderService, id int) error {
//     return os.Apply(id)
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

