export interface Buses{
    id: number;
    placa: string;
    modelo: string;
    capacidad: number;
    estado: string;
}

export interface Rutas{
    id: number;
    nombre: string;
    origen: string;
    destino: string;
    precio_viaje: number;
    distancia: number;
}

export interface Conductores{
    id: number;
    nombre: string;
    cedula: string;
    num_licencia: number;
    estado: string;
}

export interface Asignaciones{
    id?: number;
    id_conductor: number;
    id_ruta: number;
    id_bus: number;
}