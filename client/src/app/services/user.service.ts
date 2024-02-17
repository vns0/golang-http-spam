import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import { first, catchError, tap } from 'rxjs/operators';
import {environment} from "../environments/environment";
import {AuthSuccessResponse} from "../models/auth";
import {user} from "../models/users";

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private url = `${environment.apiUrl}`;

  httpOptions: { headers: HttpHeaders } = {
    headers: new HttpHeaders({ "Content-Type": "application/json"}),
  }

  constructor(
    private http: HttpClient,
  ) {}

  getUsers(): Observable<{}> {
    return this.http
      .get(`${this.url}/users`, this.httpOptions)
      .pipe(
        first(),
        tap((response: Array<user>) => {
          return response
        }),
        catchError(error => {
          return throwError(error.error);
        })
      );
  }
  createUser(UserID: number): Observable<{}> {
    return this.http
      .post(`${this.url}/users`, {UserID: UserID}, this.httpOptions)
      .pipe(
        first(),
        tap((response: Array<user>) => {
          return response
        }),
        catchError(error => {
          return throwError(error.error);
        })
      );
  }
  deleteUser(UserID: number): Observable<{}> {
    return this.http
      .post(`${this.url}/users/delete`, {UserID: UserID}, this.httpOptions)
      .pipe(
        first(),
        tap((response: Array<user>) => {
          return response
        }),
        catchError(error => {
          return throwError(error.error);
        })
      );
  }
}
