package ecs

// SystemInterface - interface that represents a system, world is an interface and should be cast to whatever data
// structure the game is currently using or that the system cares about.
type SystemInterface interface {
	Update(data interface{}, entity *Entity) error
}

// SystemManager - contains a list of systems and is responsible for calling their update functions on entities.
type SystemManager struct {
	systems []SystemInterface
}

func (s *SystemManager) AddSystem(system SystemInterface) {
	if s.systems == nil {
		s.systems = make([]SystemInterface, 0)
	}

	s.systems = append(s.systems, system)
}

// UpdateSystemsForEntity - Iterates through the systems for the specific entity
func (s *SystemManager) UpdateSystemsForEntity(world interface{}, entity *Entity) error {
	for system := range s.systems {
		err := s.systems[system].Update(world, entity)
		if err != nil {
			return err
		}
	}
	return nil
}
