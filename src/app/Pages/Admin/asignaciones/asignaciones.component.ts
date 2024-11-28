import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { AsignacionesService } from '../../../service/asignacion/asignaciones.service';
import { BusesService } from '../../../service/bus/buses.service';
import { RutasService } from '../../../service/ruta/rutas.service';
import { ConductoresService } from '../../../service/conductor/conductores.service';

@Component({
  selector: 'app-asignaciones',
  templateUrl: './asignaciones.component.html',
  styleUrl: './asignaciones.component.css'
})
export class AsignacionesComponent implements OnInit {
  AsignacionesList: any;
  BusesList: any;
  RutasList: any;
  ConductoresList: any;

  AsigacionForm: FormGroup;
  isEditMode: boolean = false; 
  editingAsignacionId?: number ; 

  constructor(
    private asignacionesServices: AsignacionesService,
    private busesServices: BusesService,
    private rutasServices: RutasService,
    private conductoresServices: ConductoresService
  ) {
    this.AsigacionForm = new FormGroup({
      id_conductor: new FormControl('', [Validators.required]),
      id_bus: new FormControl('', [Validators.required]),
      id_ruta: new FormControl('', [Validators.required]),
    });
  }

  ngOnInit() {
    this.obtenerAsignaciones();
    this.obtenerBuses();
    this.obtenerRutas();
    this.obtenerConductores();
  }

  obtenerAsignaciones() {
    this.asignacionesServices.getAsignaciones().subscribe((asignaciones: any) => {
      this.AsignacionesList = asignaciones;
    });
  }

  obtenerBuses() {
    this.busesServices.getBuses().subscribe((res) => {
      this.BusesList = res;
    });
  }

  obtenerRutas() {
    this.rutasServices.getRutas().subscribe((res) => {
      this.RutasList = res;
    });
  }

  obtenerConductores() {
    this.conductoresServices.getConductores().subscribe((res) => {
      this.ConductoresList = res;
    });
  }

  editarAsignacion(id: number) {
    this.isEditMode = true;
    this.editingAsignacionId = id;

    const asignacion = this.AsignacionesList.find((a: any) => a.id === id);
    if (asignacion) {
      this.AsigacionForm.patchValue({
        id_conductor: asignacion.conductor.id,
        id_bus: asignacion.bus.id,
        id_ruta: asignacion.ruta.id,
      });
    }
  }

  guardarAsignacion() {
    if (this.AsigacionForm.valid) {
      const asignacion = {
        id_bus: +this.AsigacionForm.value.id_bus,
        id_ruta: +this.AsigacionForm.value.id_ruta,
        id_conductor: +this.AsigacionForm.value.id_conductor,
      };

      if (this.isEditMode) {
        const updatedAsignacion = { ...asignacion, id: this.editingAsignacionId };

        this.asignacionesServices.updateAsignacion(updatedAsignacion).subscribe({
          next: (res) => {
            console.log('Asignación actualizada:', res);
            this.obtenerAsignaciones();
            this.resetForm();
          },
          error: (err) => {
            console.error('Error al actualizar asignación:', err);
          },
        });
      } else {
        this.asignacionesServices.createAsignacion(asignacion).subscribe({
          next: (res) => {
            console.log('Asignación creada:', res);
            this.obtenerAsignaciones();
            this.resetForm();
          },
          error: (err) => {
            console.error('Error al crear asignación:', err);
          },
        });
      }
    } else {
      console.error('Formulario inválido');
    }
  }

  eliminarAsignacion(id: number) {
    if (confirm('¿Estás seguro de que deseas eliminar esta asignación?')) {
      this.asignacionesServices.deleteAsignacion(id).subscribe({
        next: (res) => {
          console.log('Asignación eliminada:', res);
          this.obtenerAsignaciones();
        },
        error: (err) => {
          console.error('Error al eliminar asignación:', err);
        },
      });
    }
  }

  resetForm() {
    this.AsigacionForm.reset();
    this.isEditMode = false;
    this.editingAsignacionId = undefined;
  }
}
