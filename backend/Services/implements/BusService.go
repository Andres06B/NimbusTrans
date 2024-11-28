package implements

import (
	database "proyecto/Database"
	models "proyecto/Models"
)

type BusService struct {
}

func (b *BusService) Create(bus models.Bus) map[string]interface{} {
	result := database.Database.Create(&bus)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Creado correctamente",
	}
}

func (b *BusService) GetAll() []models.Bus {
	var buses []models.Bus
	result := database.Database.Find(&buses)
	if result.Error != nil {
		return nil
	}
	return buses
}

func (b *BusService) GetById(id int) (models.Bus, map[string]interface{}) {
	var bus models.Bus
	result := database.Database.First(&bus, id)
	if result.Error != nil {
		return bus, nil
	}
	return bus, map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Obtenido correctamente",
	}
}

func (b *BusService) Update(bus models.Bus) map[string]interface{} {
	result := database.Database.Model(&bus).Updates(bus)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Actualizado correctamente",
	}
}

func (b *BusService) Delete(id int) map[string]interface{} {
	var bus models.Bus
	result := database.Database.First(&bus, id)
	if result.Error != nil {
		return nil
	}
	result = database.Database.Delete(&bus)
	if result.Error != nil {
		return nil
	}
	return map[string]interface{}{
		"id":      result.RowsAffected,
		"message": "Eliminado correctamente",
	}
}
