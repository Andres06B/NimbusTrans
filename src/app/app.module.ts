import { NgModule } from '@angular/core';
import { BrowserModule, provideClientHydration } from '@angular/platform-browser';


import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './Pages/Web/home/home.component';
import { LogInComponent } from './Pages/Web/log-in/log-in.component';
import { FooterComponent } from './Components/footer/footer.component';
import { DashboardComponent } from './Pages/Admin/dashboard/dashboard.component';
import { PrincipioComponent } from './Pages/Admin/principio/principio.component';
import { RutasComponent } from './Pages/Admin/rutas/rutas.component';
import { BusesComponent } from './Pages/Admin/buses/buses.component';
import { AsignacionesComponent } from './Pages/Admin/asignaciones/asignaciones.component';
import { ConductoresComponent } from './Pages/Admin/conductores/conductores.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LogInComponent,
    FooterComponent,
    DashboardComponent,
    PrincipioComponent,
    RutasComponent,
    BusesComponent,
    AsignacionesComponent,
    ConductoresComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [
    provideClientHydration()
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
