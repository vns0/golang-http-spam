import { Component } from '@angular/core';
import {AuthService} from "../../services/auth.service";
import {CookieService} from "../../utils/cookie.service";
import {Router} from "@angular/router";
import {AuthSuccessResponse} from "../../models/auth";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  authCode: string = '';
  errorMessage: string = ''

  constructor(
    private authService: AuthService,
    private router: Router,
    private cookie: CookieService
    ) {}


  validationCode() {
    return this.authCode.length === 6
  }

  submitForm() {
    if (!this.validationCode()) {
      this.errorMessage = "Invalid or expired code"
    } else {
      this.authService.login(this.authCode)
        .subscribe(
          (response: AuthSuccessResponse) => {
            this.cookie.setCookie('access',response.access_token, 1)
            this.cookie.setCookie('refresh',response.refresh_token, 1)
            this.router.navigate(["dashboard"]);
          },
          ({error} )=> {
            this.errorMessage = error
          }
        )
    }
  }
}
