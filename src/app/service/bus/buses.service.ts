import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Buses } from '../../interfaces/interfaces';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BusesService {

  private url = 'http://localhost:8080/bus';

  constructor(private http: HttpClient) { }

  getBuses(): Observable<Buses[]>{
    return this.http.get<Buses[]>(this.url);
  }

  getBus(id: number): Observable<Buses>{
    return this.http.get<Buses>(this.url + '/' + id);
  }

  createBus(bus: Buses): Observable<Buses>{
    return this.http.post<Buses>(this.url, bus);
  }

  updateBus(bus: Buses): Observable<Buses>{
    return this.http.put<Buses>(`${this.url}/${bus.id}`, bus);
  }

  deleteBus(id: number): Observable<Buses>{
    return this.http.delete<Buses>(`${this.url}/${id}`);
  }
}
