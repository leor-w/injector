package injector

import (
	"fmt"
	"reflect"

	"github.com/leor-w/kid/utils"
)

var container *Container

func init() {
	container = New()
}

func Provide(val IProvider, opts ...Option) error {
	return container.Provide(val, opts...)
}

func Invoke(fn interface{}, opts ...Option) ([]reflect.Value, error) {
	return container.Invoke(fn, opts...)
}

func Populate(opts ...Option) error {
	return container.Populate(opts...)
}

type Container struct {
	scopes *Scope // 存储所有的 Scope
}

func (container *Container) Provide(val IProvider, opts ...Option) error {
	var options = &Options{}
	for _, o := range opts {
		o(options)
	}
	e := newEntity(val, options)
	if utils.IsNilPointer(e.instance) {
		return fmt.Errorf("container.Provide: 实例必须是一个有效的指针值")
	}
	scope := container.Scope(options.Scope)
	return scope.provide(val, options)
}

func (container *Container) Invoke(fn interface{}, opts ...Option) ([]reflect.Value, error) {
	var options *Options
	for _, o := range opts {
		o(options)
	}
	scope := container.Scope(options.Scope)
	return scope.invoke(fn, options)
}

func (container *Container) Populate(opts ...Option) error {
	var options = &Options{}
	for _, o := range opts {
		o(options)
	}
	scope := container.Scope(options.Scope)
	return scope.populate()
}

func (container *Container) Scope(name string) *Scope {
	if name == "" {
		return container.scopes
	}
	return container.scopes.getScope(name)
}

func New() *Container {
	return &Container{
		scopes: &Scope{
			buckets:     make(map[reflect.Type]*bucket),
			childScopes: make([]*Scope, 0),
		},
	}
}
