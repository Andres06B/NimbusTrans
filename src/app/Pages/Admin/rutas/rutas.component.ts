import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { RutasService } from '../../../service/ruta/rutas.service';
import { Rutas } from '../../../interfaces/interfaces';

@Component({
  selector: 'app-rutas',
  templateUrl: './rutas.component.html',
  styleUrl: './rutas.component.css'
})
export class RutasComponent {
  listaRutas: Rutas[] = [];
  rutasForm: FormGroup;
  isEditMode: boolean = false;
  editingRutaId?: number;

  constructor(private rutasService: RutasService) { 
    this.rutasForm = new FormGroup({
      nombre: new FormControl('',[Validators.required]),
      origen: new FormControl('',[Validators.required]),
      destino: new FormControl('',[Validators.required]),
      precio_viaje: new FormControl('',[Validators.required]),
      distancia: new FormControl('',[Validators.required]),
    });
  }

  ngOnInit() { 
    this.obtenerRutas();
  }

  obtenerRutas() {
    this.rutasService.getRutas().subscribe((res) => {
      this.listaRutas = res;
    });
  }

  
  editarRuta(id:number){
    this.isEditMode = true;
    this.editingRutaId = id;
    const ruta = this.listaRutas.find((a: any)=> a.id === id);
    if (ruta) {
      this.rutasForm.setValue({
        nombre: ruta.nombre,
        origen: ruta.origen,
        destino: ruta.destino,
        precio_viaje: ruta.precio_viaje,
        distancia: ruta.distancia,
      });
    }
  }

  guardarRuta(){
    if (this.rutasForm.valid){
      const ruta = {
        ...this.rutasForm.value,
      }
      if (this.isEditMode){
        const updateRuta = {
          ...ruta,
          id: this.editingRutaId
        }

        this.rutasService.updateRuta(updateRuta).subscribe({
          next: (res) => {
            console.log('Ruta actualizada:', res);
            this.obtenerRutas();
            this.resetForm();
          },
          error: (err) => {
            console.error('Error al actualizar ruta:', err);
          },
        })
      } else {
        this.rutasService.createRuta(ruta).subscribe({
          next: (res) => {
            console.log('Ruta creada:', res);
            this.obtenerRutas();
            this.resetForm();
          },
          error: (err) => {
            console.error('Error al crear ruta:', err);
          },
        })
      }
    }
  }

  eliminarRuta(id: number) {
    if (confirm('¿Estás seguro de que deseas eliminar esta ruta?')) {
      this.rutasService.deleteRuta(id).subscribe({
        next: (res) => {
          console.log('Ruta eliminada:', res);
          this.obtenerRutas();
        },
        error: (err) => {
          console.error('Error al eliminar ruta:', err);
        },
      });
    }
  }

  resetForm() {
    this.rutasForm.reset();
    this.isEditMode = false;
    this.editingRutaId = undefined;
  }

}
