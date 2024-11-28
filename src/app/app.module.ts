import { NgModule } from '@angular/core';
import { BrowserModule, provideClientHydration } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { HomeComponent } from './Pages/Web/home/home.component';
import { ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { FooterComponent } from './Components/footer/footer.component';
import { AsignacionesComponent } from './Pages/Admin/asignaciones/asignaciones.component';
import { BusesComponent } from './Pages/Admin/buses/buses.component';
import { ConductoresComponent } from './Pages/Admin/conductores/conductores.component';
import { DashboardComponent } from './Pages/Admin/dashboard/dashboard.component';
import { PrincipioComponent } from './Pages/Admin/principio/principio.component';
import { RutasComponent } from './Pages/Admin/rutas/rutas.component';
import { LogInComponent } from './Pages/Web/log-in/log-in.component';
import { provideHttpClient } from '@angular/common/http';





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
    AppRoutingModule,
    ReactiveFormsModule
  ],
  providers: [
    provideClientHydration(),
    provideHttpClient()
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
