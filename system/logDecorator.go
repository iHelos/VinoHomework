package system

type LogDecorator func(LogComponent) LogComponent

func Decorate(c LogComponent, ds ...LogDecorator) LogComponent{
	decorated := c
	for _, decorator := range ds {
		decorated = decorator(decorated)
	}
	return decorated
}