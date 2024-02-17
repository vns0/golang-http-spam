import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';

import { Observable } from 'rxjs';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate{

  constructor(
    private authService: AuthService,
    private router: Router) { }

    canActivate(): Observable<boolean> {
      if (!this.authService.isLoggedInSubject.value) {
        this.router.navigate(["login"])
      }
      return this.authService.isUserLoggedIn$;
    }
}
