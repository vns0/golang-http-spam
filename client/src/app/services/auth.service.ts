import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import {BehaviorSubject, Observable, throwError} from 'rxjs';
import { first, catchError, tap } from 'rxjs/operators';

import { ErrorHandlerService } from './error-handler.service';
import {environment} from "../environments/environment";
import {AuthSuccessResponse} from "../models/auth";
import {CookieService} from "../utils/cookie.service";
import {Router} from "@angular/router";



@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private url = `${environment.apiUrl}/auth`;

  isLoggedInSubject = new BehaviorSubject<boolean>(this.checkLoginStatus());
  isUserLoggedIn$ = this.isLoggedInSubject.asObservable();

  httpOptions: { headers: HttpHeaders } = {
    headers: new HttpHeaders({ "Content-Type": "application/json"}),
  }

  constructor(
    private http: HttpClient,
    private errorHandlerService: ErrorHandlerService,
    private cookieService: CookieService,
    private router: Router
    ) {}

  checkLoginStatus(): boolean {
    return !!(this.cookieService.getCookie('access') && this.cookieService.getCookie('refresh'));
  }

  login(code: string): Observable<{}> {
    return this.http
      .post(`${this.url}/login`, {code}, this.httpOptions)
      .pipe(
        first(),
        tap((response: AuthSuccessResponse) => {
          this.isLoggedInSubject.next(true);
          return response
        }),
        catchError(error => {
          this.errorHandlerService.handleError<{
            access_token: string; code: string
          }>("login")
          return throwError(error.error);
        })
      );
  }
  logout() {
    this.cookieService.deleteCookie('access')
    this.cookieService.deleteCookie('refresh')
    this.isLoggedInSubject.next(false);
    this.router.navigate(['/login']);
  }
}
