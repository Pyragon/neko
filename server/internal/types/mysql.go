package types

type MySQLHandler interface {
	Start() error
	Connect() error
	Select() error
	Shutdown() error
}
