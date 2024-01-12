package managers

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Dependency func(args ...any) []any

const DependencyIndicator = "__"

type DependencyManager struct {
	dependencies map[string]Dependency
}

func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		dependencies: make(map[string]Dependency),
	}
}

func (d *DependencyManager) Get(key string) Dependency {
	return d.dependencies[key]
}

func (d *DependencyManager) Init(modules ...any) {
	for i := 0; i < len(modules); i++ {
		ModuleReflectedValue := reflect.ValueOf(modules[i]).Elem()
		moduleReflectedType := ModuleReflectedValue.Type()

		for j := 0; j < moduleReflectedType.NumMethod(); j++ {

			method := moduleReflectedType.Method(j)
			//logger.Info(fmt.Sprintf("Method Name: %v", method.Name))
			matched, _ := regexp.MatchString(fmt.Sprintf(`\w+%v$`, DependencyIndicator), method.Name)
			if matched {
				d.dependencies[strings.Replace(method.Name, DependencyIndicator, "", -1)] = createFunction(ModuleReflectedValue.MethodByName(method.Name))
			}
		}
	}
}

func (d *DependencyManager) GetAll() map[string]Dependency {
	return d.dependencies
}

func createFunction(reflectedMethod reflect.Value) func(args ...any) []any {
	return func(args ...any) []any {
		rArgs := []reflect.Value{}
		for i := 0; i < len(args); i++ {
			rArgs = append(rArgs, reflect.ValueOf(args[i]))
		}

		rResults := reflectedMethod.Call(rArgs)

		results := []any{}

		for i := 0; i < len(rResults); i++ {
			results = append(results, rResults[i].Interface())
		}
		return results

	}
}
