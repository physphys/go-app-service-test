package dicontainer

import "errors"

var (
	errDefinitionNotFound = errors.New("build definition not found")
	CastError             = errors.New("casting error")
)

const (
	DefNameUserRepository    = "UserRepository"
	DefNameUserDomainService = "UserDomainService"
	DefNameUserAppService    = "UserAppService"
)

type (
	Builder func(con *Container) (interface{}, error)

	Definition struct {
		Name    string
		Builder Builder
	}

	Container struct {
		store map[string]Builder
	}
)

func NewContainer() *Container {
	return &Container{
		store: map[string]Builder{},
	}
}

func (c *Container) Register(d *Definition) {
	c.store[d.Name] = d.Builder
}

func (c *Container) Get(key string) (interface{}, error) {
	builder, ok := c.store[key]
	if !ok {
		return nil, errDefinitionNotFound
	}

	instance, err := builder(c)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
