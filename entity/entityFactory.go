package entity

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/mechanical-lich/game-engine/component"
)

type ComponentAddFunction func(params []string) (component.Component, error)

var blueprints = make(map[string][]string)
var componentAddFunctions = make(map[string]ComponentAddFunction)

// LoadBlueprintsFromFile - Loads the blueprints for the factory to construct entities
func LoadBlueprintsFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	return LoadFactoryFromStream(file)
}

// LoadFactorFromStream - Loads the blueprints from an io stream.
func LoadFactoryFromStream(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanLines)

	entityName := ""
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			entityName = ""
			continue
		}
		if entityName == "" {
			entityName = value
			continue
		} else {
			blueprints[entityName] = append(blueprints[entityName], value)
		}
	}

	return nil
}

// RegisterComponentAddFunction - Register a function pointer to point be used whenever the factory goes to build
// an entity from a blueprint using this component.
func RegisterComponentAddFunction(name string, function ComponentAddFunction) {
	componentAddFunctions[name] = function
}

// Create - Creates an entity from the named blueprint.
func Create(name string) (*Entity, error) {
	blueprint := blueprints[name]
	if blueprint != nil {
		entity := Entity{}
		entity.Blueprint = name

		for _, value := range blueprint {
			c := strings.Split(value, ":")
			params := strings.Split(c[1], ",")
			if componentAddFunctions[c[0]] != nil {
				newComp, err := componentAddFunctions[c[0]](params)
				if err != nil {
					return nil, err
				}
				entity.AddComponent(newComp)
			} else {
				return nil, errors.New("no component handler function registered")
			}

		}
		return &entity, nil
	}
	return nil, errors.New("no blueprint found")
}
