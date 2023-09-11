package injector

import (
	"fmt"
	"github.com/leor-w/utils"
	"reflect"
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
	scope *Scope // 存储所有的 Scope
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
	if err := scope.populate(); err != nil {
		return err
	}
	return nil
}

func (container *Container) Scope(name string) *Scope {
	if name == "" {
		return container.scope
	}
	scope := container.scope.getScope(name)
	if scope != nil {
		return scope
	}
	return container.scope.Scope(name)
}

func New() *Container {
	return &Container{
		scope: &Scope{
			buckets:     make(map[reflect.Type]*bucket),
			childScopes: make([]*Scope, 0),
		},
	}
}
