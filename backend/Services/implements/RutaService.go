package implements

import (
	database "proyecto/Database"
	models "proyecto/Models"
)

type RutaService struct {
}

func (r *RutaService) Create(ruta models.Ruta) map[string]interface{} {
	result := database.Database.Create(&ruta)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Creado correctamente",
	}
}

func (r *RutaService) GetAll() []models.Ruta {
	var rutas []models.Ruta
	result := database.Database.Find(&rutas)
	if result.Error != nil {
		return nil
	}
	return rutas
}

func (r *RutaService) GetById(id int) (models.Ruta, map[string]interface{}) {
	var ruta models.Ruta
	result := database.Database.First(&ruta, id)
	if result.Error != nil {
		return ruta, nil
	}
	return ruta, map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Obtenido correctamente",
	}
}

func (r *RutaService) Update(ruta models.Ruta) map[string]interface{} {
	result := database.Database.Model(&ruta).Updates(ruta)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Actualizado correctamente",
	}
}

func (r *RutaService) Delete(id int) map[string]interface{} {
	var ruta models.Ruta
	result := database.Database.First(&ruta, id)
	if result.Error != nil {
		return nil
	}
	result = database.Database.Delete(&ruta)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Eliminado correctamente",
	}
}
