package system

type LogComponent interface {
	Log(string)
}

//Сигнатура функции логгирования
type LogStr func(string)

//Метод удовлетворяющий интерфейсу Компонента, который в дальнейшем будет декорироваться
func (f LogStr) Log(str string) {
	f(str)
}
