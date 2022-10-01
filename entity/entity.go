package entity

import (
	"github.com/mechanical-lich/game-engine/component"
)

// Entity - represents an entity which essentially is an array of components built from a blueprint.
type Entity struct {
	Components map[string]component.Component
	Shown      bool
	Blueprint  string
}

// AddComponent - Adds the provided component to the entity.
func (entity *Entity) AddComponent(c component.Component) {
	if entity.Components == nil {
		entity.Components = make(map[string]component.Component)
	}

	entity.Components[c.GetType()] = c
}

// HasComponent - Returns if the entity has the component.
func (entity *Entity) HasComponent(name string) bool {
	if entity.Components == nil {
		entity.Components = make(map[string]component.Component)
	}

	return entity.Components[name] != nil
}

// GetComponent - Gets the component as a component interface.
func (entity *Entity) GetComponent(name string) component.Component {
	if entity.Components == nil {
		entity.Components = make(map[string]component.Component)
	}

	return entity.Components[name]
}

// RemoveComponent - Removes the component from the entity.
func (entity *Entity) RemoveComponent(name string) {
	if entity.Components == nil {
		entity.Components = make(map[string]component.Component)
	}

	entity.Components[name] = nil
}
