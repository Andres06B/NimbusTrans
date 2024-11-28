import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Rutas } from '../../interfaces/interfaces';

@Injectable({
  providedIn: 'root'
})
export class RutasService {

  private url = 'http://localhost:8080/ruta';

  constructor(private http: HttpClient) { }

  getRutas(): Observable<Rutas[]>{
    return this.http.get<Rutas[]>(this.url);
  }

  getRuta(id: number): Observable<Rutas>{
    return this.http.get<Rutas>(this.url + '/' + id);
  }

  createRuta(ruta: Rutas): Observable<Rutas>{
    return this.http.post<Rutas>(this.url, ruta);
  }

  updateRuta(ruta: Rutas): Observable<Rutas>{
    return this.http.put<Rutas>(this.url, ruta);
  }

  deleteRuta(id: number): Observable<Rutas>{
    return this.http.delete<Rutas>(this.url + '/' + id);
  }
}
