package di

import (
	"reflect"
	"runtime"

	"go.uber.org/dig"
)

var container = dig.New()
var componentRegistry = make([]interface{}, 0)
var isInitialized = false

func init() {
	for _, target := range componentRegistry {
		provide(target)
	}
	isInitialized = true
}

func RegisterBean(target interface{}) {
	if isInitialized {
		provide(target)
	} else {
		componentRegistry = append(componentRegistry, target)
	}
}

func Invoke(function interface{}) {
	if err := container.Invoke(function); err != nil {
		//logger.Default().Fatal("container.Invoke err - target : %s err : %s", reflect.ValueOf(function).String(), err)
	}
	//logger.Default().Debug("container.Invoke end - target : %s", reflect.ValueOf(function).String())
}

func GetContainer() *dig.Container {
	return container
}

func provide(target interface{}) {
	if err := container.Provide(target); err != nil {
		var _ = runtime.FuncForPC(reflect.ValueOf(target).Pointer()).Name()
		//logger.Default().Fatal("container.Provide err - func : %s err : %s", funcName, err)
	}
	//logger.Default().Debug("container.Provide - func : %s", reflect.ValueOf(target).String())
}
