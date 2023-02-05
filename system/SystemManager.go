package system

import "github.com/mechanical-lich/game-engine/entity"

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

// UpdateSystemsForEntity - Iterates through the systems for the specific entity.
func (s *SystemManager) UpdateSystemsForEntity(world interface{}, entity *entity.Entity) error {
	for system := range s.systems {
		err := s.systems[system].Update(world, entity)
		if err != nil {
			return err
		}
	}
	return nil
}
