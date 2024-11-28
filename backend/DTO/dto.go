package dto

type ConductorDTO struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Cedula      string `json:"cedula"`
	NumLicencia int    `json:"num_licencia"`
	Estado      string `json:"estado"`
}

type BusDTO struct {
	Id        int    `json:"id"`
	Placa     string `json:"placa"`
	Modelo    string `json:"modelo"`
	Capacidad int    `json:"capacidad"`
	Estado    string `json:"estado"`
}

type RutaDTO struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Origen      string `json:"origen"`
	Destino     string `json:"destino"`
	PrecioViaje int    `json:"precio_viaje"`
	Distancia   string `json:"distancia"`
}

type AsignacionDTO struct {
	Id           int `json:"id"`
	Id_bus       int `json:"id_bus"`
	Id_conductor int `json:"id_conductor"`
	Id_ruta      int `json:"id_ruta"`
}
