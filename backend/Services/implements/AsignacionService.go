package implements

import (
	database "proyecto/Database"
	models "proyecto/Models"
)

type AsignacionService struct {
}

func (a *AsignacionService) Create(asignacion models.Asignacion) map[string]interface{} {
	result := database.Database.Create(&asignacion)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Creado correctamente",
	}
}

func (a *AsignacionService) GetAll() []models.Asignacion {
	var asignaciones []models.Asignacion
	result := database.Database.Preload("Conductor").Preload("Bus").Preload("Ruta").Find(&asignaciones)
	if result.Error != nil {
		return nil
	}
	return asignaciones
}

func (a *AsignacionService) GetById(id int) (models.Asignacion, map[string]interface{}) {
	var asignacion models.Asignacion
	result := database.Database.Preload("Conductor").Preload("Bus").Preload("Ruta").First(&asignacion, id)
	if result.Error != nil {
		return asignacion, nil
	}
	return asignacion, map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Obtenido correctamente",
	}
}

func (a *AsignacionService) Update(asignacion models.Asignacion) map[string]interface{} {
	result := database.Database.Model(&asignacion).Updates(asignacion)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Actualizado correctamente",
	}
}

func (a *AsignacionService) Delete(id int) map[string]interface{} {
	var asignacion models.Asignacion
	result := database.Database.First(&asignacion, id)
	if result.Error != nil {
		return nil
	}
	result = database.Database.Delete(&asignacion)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Eliminado correctamente",
	}
}
