import { Component } from '@angular/core';
import { ConductoresService } from '../../../service/conductor/conductores.service';
import { Conductores } from '../../../interfaces/interfaces';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-conductores',
  templateUrl: './conductores.component.html',
  styleUrl: './conductores.component.css'
})
export class ConductoresComponent {


  ConductoresList: Conductores[] = [];
  ConductorForm: FormGroup;
  isEditMode: boolean = false;
  editingConductorId?: number;

  constructor(private conductoresService: ConductoresService) {
    this.ConductorForm = new FormGroup({
      nombre: new FormControl('',[Validators.required]),
      cedula: new FormControl('',[Validators.required]),
      num_licencia: new FormControl('',[Validators.required]),
      estado: new FormControl('',[Validators.required]),
    });
  }

  ngOnInit() {
    this.obtenerConductores();
  }

  obtenerConductores() {
    this.conductoresService.getConductores().subscribe((res) => {
      this.ConductoresList = res;
    });
  }

  editarConductor(id: number) {
    this.isEditMode = true;
    this.editingConductorId = id;

    const conductor = this.ConductoresList.find((a: any) => a.id === id);
    if (conductor) {
      this.ConductorForm.setValue({
        nombre: conductor.nombre,
        cedula: conductor.cedula,
        num_licencia: conductor.num_licencia,
        estado: conductor.estado,
      });
    }
  }

  guardarConductor() {
    if (this.ConductorForm.valid) {
      const conductor = {
        ...this.ConductorForm.value,
      }
      if (this.isEditMode) {
        const updatedConductor = { ...conductor, id: this.editingConductorId };
        this.conductoresService.updateConductor(updatedConductor).subscribe({
          next: (res) => {
            console.log('Conductor actualizado:', res);
            this.obtenerConductores();
            this.resetForm();
          },
          error: (err) => {
            console.error('Error al actualizar conductor:', err);
          },
        });
      } else {
        this.conductoresService.createConductor(conductor).subscribe({
          next: (res) => {
            console.log('Conductor creado:', res);
            this.obtenerConductores();
            this.resetForm();
          },
          error: (err) => {
            console.error('Error al crear conductor:', err);
          },
        });
      }
    } else {
      console.error('Formulario inválido');
    }
  }

  eliminarConductor(id: number) {
    if (confirm('¿Estás seguro de que deseas eliminar este conductor?')) {
      this.conductoresService.deleteConductor(id).subscribe({
        next: (res) => {
          console.log('Conductor eliminado:', res);
          this.obtenerConductores();
        },
        error: (err) => {
          console.error('Error al eliminar conductor:', err);
        },
      });
    }
  }

  resetForm() {
    this.ConductorForm.reset();
    this.isEditMode = false;
    this.editingConductorId = undefined;  
  }




}
