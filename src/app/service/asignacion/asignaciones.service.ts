import { Injectable } from '@angular/core';
import { Asignaciones } from '../../interfaces/interfaces';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AsignacionesService {

  private url = 'http://localhost:8080/asignacion';

  constructor(private http: HttpClient) { }

  getAsignaciones(): Observable<Asignaciones[]>{
    return this.http.get<Asignaciones[]>(this.url);
  }

  getAsignacion(id: number): Observable<Asignaciones>{
    return this.http.get<Asignaciones>(this.url + '/' + id);
  }

  createAsignacion(asignacion: Asignaciones): Observable<Asignaciones>{
    return this.http.post<Asignaciones>(this.url, asignacion);
  }

  updateAsignacion(asignacion: Asignaciones): Observable<Asignaciones> {
    return this.http.put<Asignaciones>(`${this.url}/${asignacion.id}`, asignacion);
  }
  
  deleteAsignacion(id: number): Observable<Asignaciones>{
    return this.http.delete<Asignaciones>(this.url + '/' + id);
  }
}
