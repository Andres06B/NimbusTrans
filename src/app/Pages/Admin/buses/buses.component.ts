import { Component } from '@angular/core';
import { BusesService } from '../../../service/bus/buses.service';
import { Buses } from '../../../interfaces/interfaces';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-buses',
  templateUrl: './buses.component.html',
  styleUrl: './buses.component.css'
})
export class BusesComponent {
[x: string]: any;

  BusesList: Buses[] = [];
  BusForm: FormGroup;
  constructor(private busesServices: BusesService){
    this.BusForm = new FormGroup({
      placa: new FormControl('',[Validators.required]),
      modelo: new FormControl('',[Validators.required]),
      capacidad: new FormControl('',[Validators.required]),
      estado: new FormControl('',[Validators.required]),
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

  crearBus() {
    if (this.BusForm.valid) {
      this.busesServices.createBus(this.BusForm.value).subscribe((res) => {
        console.log('Bus creado: ', res);
      });
    }
  }
  
}
