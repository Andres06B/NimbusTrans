import { Component } from '@angular/core';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css'
})
export class DashboardComponent {
  seccionActiva: string = 'dashboard'; // Sección activa por defecto

  // Cambiar la sección activa
  mostrarSeccion(seccion: string): void {
    this.seccionActiva = seccion;
  }
}
