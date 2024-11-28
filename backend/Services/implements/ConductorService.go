package implements

import (
	database "proyecto/Database"
	models "proyecto/Models"
)

type ConductorService struct {
}

func (c *ConductorService) Create(conductor models.Conductor) map[string]interface{} {
	result := database.Database.Create(&conductor)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Creado correctamente",
	}
}

func (c *ConductorService) GetAll() []models.Conductor {
	var conductors []models.Conductor
	result := database.Database.Find(&conductors)
	if result.Error != nil {
		return nil
	}
	return conductors
}

func (c *ConductorService) GetById(id int) (models.Conductor, map[string]interface{}) {
	var conductor models.Conductor
	result := database.Database.First(&conductor, id)
	if result.Error != nil {
		return conductor, nil
	}
	return conductor, map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Obtenido correctamente",
	}
}

func (c *ConductorService) Update(conductor models.Conductor) map[string]interface{} {
	result := database.Database.Model(&conductor).Updates(conductor)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Actualizado correctamente",
	}
}

func (c *ConductorService) Delete(id int) map[string]interface{} {
	var conductor models.Conductor
	result := database.Database.First(&conductor, id)
	if result.Error != nil {
		return nil
	}
	result = database.Database.Delete(&conductor)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Eliminado correctamente",
	}
}
