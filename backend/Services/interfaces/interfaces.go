package interfaces

import models "proyecto/Models"

type BusService interface {
	Create(bus models.Bus) map[string]interface{}
	GetAll() []models.Bus
	GetById(id int) (models.Bus, map[string]interface{})
	Update(bus models.Bus) map[string]interface{}
	Delete(id int) map[string]interface{}
}

type ConductorService interface {
	Create(conductor models.Conductor) map[string]interface{}
	GetAll() []models.Conductor
	GetById(id int) (models.Conductor, map[string]interface{})
	Update(conductor models.Conductor) map[string]interface{}
	Delete(id int) map[string]interface{}
}

type RutaService interface {
	Create(ruta models.Ruta) map[string]interface{}
	GetAll() []models.Ruta
	GetById(id int) (models.Ruta, map[string]interface{})
	Update(ruta models.Ruta) map[string]interface{}
	Delete(id int) map[string]interface{}
}

type AsignacionService interface {
	Create(asignacion models.Asignacion) map[string]interface{}
	GetAll() []models.Asignacion
	GetById(id int) (models.Asignacion, map[string]interface{})
	Update(asignacion models.Asignacion) map[string]interface{}
	Delete(id int) map[string]interface{}
}
