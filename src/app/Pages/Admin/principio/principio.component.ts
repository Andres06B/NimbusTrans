import { Component, OnInit } from '@angular/core';
import { BusesService } from '../../../service/bus/buses.service';
import { RutasService } from '../../../service/ruta/rutas.service';
import { ConductoresService } from '../../../service/conductor/conductores.service';
import { AsignacionesService } from '../../../service/asignacion/asignaciones.service';

@Component({
  selector: 'app-principio',
  templateUrl: './principio.component.html',
  styleUrls: ['./principio.component.css'],
})
export class PrincipioComponent implements OnInit {
  rutasActivas: number = 0;
  busesRegistrados: number = 0;
  conductores: number = 0;
  balance: number = 0; 
  actividades: any[] = []; 

  constructor(
    private busesService: BusesService,
    private rutasService: RutasService,
    private conductoresService: ConductoresService,
    
  ) {}

  ngOnInit() {
    this.cargarDashboard();
  }

  cargarDashboard() {
    
    this.rutasService.getRutas().subscribe((rutas: any[]) => {
      this.rutasActivas = rutas.length;
      this.balance = rutas.reduce((total, ruta) => total + (ruta.precio_viaje || 0), 0);
    });

    
    this.busesService.getBuses().subscribe((buses: any[]) => {
      this.busesRegistrados = buses.length;
    });

    
    this.conductoresService.getConductores().subscribe((conductores: any[]) => {
      this.conductores = conductores.length;
    });

    this.cargarActividadesRecientes();
  }

  cargarActividadesRecientes() {
    
    this.actividades = [
      { id: 1, fecha: '2024-11-27', actividad: 'Registro de nuevo bus', detalles: 'Detalles del registro' },
      { id: 2, fecha: '2024-11-26', actividad: 'Nuevo conductor asignado', detalles: 'Detalles del conductor' },
      { id: 3, fecha: '2024-11-25', actividad: 'Ruta activa actualizada', detalles: 'Detalles de la ruta' },
    ];
  }
}
