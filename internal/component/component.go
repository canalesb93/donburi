package component

import (
	"fmt"
	"reflect"
	"unsafe"
)

type ComponentTypeId int

// CompnentType represents a type of component. It is used to identify
// a component when getting or setting components of an entity.
type ComponentType struct {
	id         ComponentTypeId
	typ        reflect.Type
	name       string
	defaultVal interface{}
}

var nextComponentTypeId ComponentTypeId = 1

// NewComponentType creates a new component type.
// The argument is a struct that represents a data of the component.
func NewComponentType(s interface{}, defaultVal interface{}) *ComponentType {
	componentType := &ComponentType{
		id:         nextComponentTypeId,
		typ:        reflect.TypeOf(s),
		name:       reflect.TypeOf(s).Name(),
		defaultVal: defaultVal,
	}
	if defaultVal != nil {
		componentType.validateDefaultVal()
	}
	nextComponentTypeId++
	return componentType
}

// String returns the component type name.
func (c *ComponentType) String() string {
	return c.name
}

// SetName sets the component type name.
func (c *ComponentType) SetName(name string) *ComponentType {
	c.name = name
	return c
}

// Name returns the component type name.
func (c *ComponentType) Name() string {
	return c.name
}

// Id returns the component type id.
func (c *ComponentType) Id() ComponentTypeId {
	return c.id
}

func (c *ComponentType) New() unsafe.Pointer {
	val := reflect.New(c.typ)
	v := reflect.Indirect(val)
	ptr := unsafe.Pointer(v.UnsafeAddr())
	if c.defaultVal != nil {
		c.setDefaultVal(ptr)
	}
	return ptr
}

func (c *ComponentType) setDefaultVal(ptr unsafe.Pointer) {
	v := reflect.Indirect(reflect.ValueOf(c.defaultVal))
	reflect.NewAt(c.typ, ptr).Elem().Set(v)
}

func (c *ComponentType) validateDefaultVal() {
	if !reflect.TypeOf(c.defaultVal).AssignableTo(c.typ) {
		err := fmt.Sprintf("default value is not assignable to component type: %s", c.name)
		panic(err)
	}
}
