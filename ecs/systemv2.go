package ecs

// SystemInterface - interface that represents a system, world is an interface and should be cast to whatever data
// structure the game is currently using or that the system cares about.
type SystemV2Interface interface {
	UpdateSystem(data interface{}) error
	UpdateEntity(data interface{}, entity *Entity) error
}

// SystemManager - contains a list of systems and is responsible for calling their update functions on entities.
type SystemV2Manager struct {
	systems []SystemV2Interface
}

func (s *SystemV2Manager) AddSystem(system SystemV2Interface) {
	if s.systems == nil {
		s.systems = make([]SystemV2Interface, 0)
	}

	s.systems = append(s.systems, system)
}

func (s *SystemV2Manager) UpdateSystems(world interface{}) error {
	for system := range s.systems {
		err := s.systems[system].UpdateSystem(world)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateSystemsForEntity - Iterates through the systems for the specific entity
func (s *SystemV2Manager) UpdateSystemsForEntity(world interface{}, entity *Entity) error {
	for system := range s.systems {
		err := s.systems[system].UpdateEntity(world, entity)
		if err != nil {
			return err
		}
	}
	return nil
}
