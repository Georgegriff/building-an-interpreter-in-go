package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(key string) (Object, bool) {
	obj, ok := e.store[key]
	return obj, ok
}

func (e *Environment) Set(key string, value Object) Object {
	e.store[key] = value
	return value
}
