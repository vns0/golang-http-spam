import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import { first, catchError, tap } from 'rxjs/operators';
import {environment} from "../environments/environment";
import {AuthSuccessResponse} from "../models/auth";

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  private url = `${environment.apiUrl}`;

  httpOptions: { headers: HttpHeaders } = {
    headers: new HttpHeaders({ "Content-Type": "application/json"}),
  }

  constructor(
    private http: HttpClient,
  ) {}

  getStats(): Observable<{}> {
    return this.http
      .get(`${this.url}/stats`, this.httpOptions)
      .pipe(
        first(),
        tap((response: AuthSuccessResponse) => {
          return response
        }),
        catchError(error => {
          return throwError(error.error);
        })
      );
  }
}
