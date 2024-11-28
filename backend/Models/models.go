package models

type Asignacion struct {
	Id          int       `gorm:"primary_key" json:"id"`
	IDconductor int       `json:"id_conductor"`
	IDbus       int       `json:"id_bus"`
	IDruta      int       `json:"id_ruta"`
	Conductor   Conductor `gorm:"foreignKey:IDconductor" json:"conductor"`
	Bus         Bus       `gorm:"foreignKey:IDbus " json:"bus"`
	Ruta        Ruta      `gorm:"foreignKey:IDruta " json:"ruta"`
}

func (a *Asignacion) TableName() string {
	return "asignacion"
}

type Conductor struct {
	Id           int    `gorm:"primary_key" json:"id"`
	Nombre       string `gorm:"type:varchar(255)" json:"nombre"`
	Cedula       string `gorm:"type:varchar(255)" json:"cedula"`
	Num_licencia int    `gorm:"type:int" json:"num_licencia"`
	Estado       string `gorm:"Enum('activo', 'inactivo')" json:"estado"`
}

func (c *Conductor) TableName() string {
	return "conductor"
}

type Bus struct {
	Id        int    `gorm:"primary_key" json:"id"`
	Placa     string `gorm:"type:varchar(255)" json:"placa"`
	Modelo    string `gorm:"type:varchar(255)" json:"modelo"`
	Capacidad int    `gorm:"type:int" json:"capacidad"`
	Estado    string `gorm:"Enum('activo', 'inactivo')" json:"estado"`
}

func (b *Bus) TableName() string {
	return "bus"
}

type Ruta struct {
	Id           int    `gorm:"primary_key" json:"id"`
	Nombre       string `gorm:"type:varchar(255)" json:"nombre"`
	Origen       string `gorm:"type:varchar(255)" json:"origen"`
	Destino      string `gorm:"type:varchar(255)" json:"destino"`
	Precio_viaje int    `gorm:"type:int" json:"precio_viaje"`
	Distancia    string `gorm:"type:varchar(255)" json:"distancia"`
}

func (r *Ruta) TableName() string {
	return "ruta"
}
