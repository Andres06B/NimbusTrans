import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Conductores } from '../../interfaces/interfaces';

@Injectable({
  providedIn: 'root'
})
export class ConductoresService {

  private url = 'http://localhost:8080/conductor';

  constructor(private http: HttpClient) { }

  getConductores(): Observable<Conductores[]>{
    return this.http.get<Conductores[]>(this.url);
  }

  getConductor(id: number): Observable<Conductores>{
    return this.http.get<Conductores>(this.url + '/' + id);
  }

  createConductor(conductor: Conductores): Observable<Conductores>{
    return this.http.post<Conductores>(this.url, conductor);
  }

  updateConductor(conductor: Conductores): Observable<Conductores>{
    return this.http.put<Conductores>(`${this.url}/${conductor.id}`, conductor);
  }

  deleteConductor(id: number): Observable<Conductores>{
    return this.http.delete<Conductores>(`${this.url}/${id}`);
  }
}
