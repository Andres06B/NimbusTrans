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

  crearRutas(){
    if(this.rutasForm.valid){
      this.rutasService.createRuta(this.rutasForm.value).subscribe((res) => {
        console.log('Ruta creada: ', res);
      })
    }
  }

}
