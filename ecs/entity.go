package ecs

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
