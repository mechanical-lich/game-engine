package ecs

import (
	"log"
	"strings"
)

// Entity - represents an entity which essentially is an array of components built from a blueprint.
type Entity struct {
	Components map[string]Component
	Shown      bool
	Blueprint  string
}

// AddComponent - Adds the provided component to the entity.
func (entity *Entity) AddComponent(c Component) {
	if entity.Components == nil {
		entity.Components = make(map[string]Component)
	}

	entity.Components[c.GetType()] = c
}

// HasComponent - Returns if the entity has the
func (entity *Entity) HasComponent(name string) bool {
	if entity.Components == nil {
		entity.Components = make(map[string]Component)
	}

	return entity.Components[name] != nil
}

// HasComponents - takes a comma separated string of component names and returns if entity has all of them.
func (entity *Entity) HasComponents(names ...string) bool {
	if entity.Components == nil {
		entity.Components = make(map[string]Component)
	}

	if len(names) == 1 {
		names = strings.Split(names[0], ",")
		if len(names) > 1 {
			log.Print("Deprecated use of HasComponents.")
		}
	}

	for _, name := range names {
		if entity.Components[name] == nil {
			return false
		}
	}

	return true
}

// GetComponent - Gets the component as a component interface.
func (entity *Entity) GetComponent(name string) Component {
	if entity.Components == nil {
		entity.Components = make(map[string]Component)
	}

	return entity.Components[name]
}

// RemoveComponent - Removes the component from the entity.
func (entity *Entity) RemoveComponent(name string) {
	if entity.Components == nil {
		entity.Components = make(map[string]Component)
	}

	entity.Components[name] = nil
}
