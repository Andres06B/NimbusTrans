import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './Pages/Web/home/home.component';
import { LogInComponent } from './Pages/Web/log-in/log-in.component';
import { DashboardComponent } from './Pages/Admin/dashboard/dashboard.component';
import path from 'path';
import { PrincipioComponent } from './Pages/Admin/principio/principio.component';
import { RutasComponent } from './Pages/Admin/rutas/rutas.component';
import { BusesComponent } from './Pages/Admin/buses/buses.component';
import { ConductoresComponent } from './Pages/Admin/conductores/conductores.component';
import { AsignacionesComponent } from './Pages/Admin/asignaciones/asignaciones.component';


const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LogInComponent },
  { path: 'dash', component: DashboardComponent },
  { path: 'principal', component: PrincipioComponent },
  { path: 'rutas', component: RutasComponent },
  { path: 'buses', component: BusesComponent },
  { path: 'conductores', component: ConductoresComponent },
  { path: 'asignaciones', component: AsignacionesComponent },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
