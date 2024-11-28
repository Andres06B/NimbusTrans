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

  crearConductor(conductor: Conductores) {
    this.conductoresService.createConductor(conductor).subscribe((res) => {
      console.log('Conductor creado');
      this.ConductoresList.push(res);
    });
  }
}
