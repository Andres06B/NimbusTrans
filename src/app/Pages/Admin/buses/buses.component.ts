import { Component } from '@angular/core';
import { BusesService } from '../../../service/bus/buses.service';
import { Buses } from '../../../interfaces/interfaces';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-buses',
  templateUrl: './buses.component.html',
  styleUrl: './buses.component.css',
})
export class BusesComponent {
  BusesList: Buses[] = [];
  BusForm: FormGroup;
  isEditMode: boolean = false;
  editingBusId?: number;

  constructor(private busesServices: BusesService) {
    this.BusForm = new FormGroup({
      placa: new FormControl('', [Validators.required]),
      modelo: new FormControl('', [Validators.required]),
      capacidad: new FormControl('', [Validators.required]),
      estado: new FormControl('', [Validators.required]),
    });
  }

  ngOnInit() {
    this.obtenerBuses();
  }

  obtenerBuses() {
    this.busesServices.getBuses().subscribe((res) => {
      this.BusesList = res;
    });
  }

  
  editarBus(id: number) {
    this.isEditMode = true;  
    this.editingBusId = id;  
    const bus = this.BusesList.find((b:any) => b.id === id);

    if (bus) {
      this.BusForm.setValue({
        placa: bus.placa,
        modelo: bus.modelo,
        capacidad: bus.capacidad,
        estado: bus.estado,
      });
    }
  }

  guardarBus() {
    if (this.BusForm.valid) {
      const bus = {
        ...this.BusForm.value,
      }
      if (this.isEditMode) {
        const updatedBus = { ...bus, id: this.editingBusId };
        this.busesServices.updateBus(updatedBus).subscribe({
          next: (res) => {
            console.log('Bus actualizado:', res);
            this.obtenerBuses();  
            this.resetForm();      
          },
          error: (err) => {
            console.error('Error al actualizar bus:', err);
          },
        });
      } else {
        this.busesServices.createBus(bus).subscribe({
          next: (res) => {
            console.log('Bus creado:', res);
            this.obtenerBuses();  
            this.resetForm();      
          },
          error: (err) => {
            console.error('Error al crear bus:', err);
          },
        });
      }
    } else {
      console.error('Formulario inválido');
    }
  }

  eliminarBus(id: number) {
    if (confirm('¿Estás seguro de que deseas eliminar este bus?')) {
      this.busesServices.deleteBus(id).subscribe({
        next: (res) => {
          console.log('Bus eliminado:', res);
          this.obtenerBuses();  
        },
        error: (err) => {
          console.error('Error al eliminar bus:', err);
        },
      });
    }
  }

  resetForm() {
    this.BusForm.reset();
    this.isEditMode = false;
    this.editingBusId = undefined;
  }
}
