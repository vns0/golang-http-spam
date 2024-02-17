import {Component, OnInit} from '@angular/core';
import {AuthService} from "../../services/auth.service";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  isLoggedIn: boolean = false;
  isFormAttackVisible = false
  newAttackForm: FormGroup;
  constructor(private authService: AuthService, private fb: FormBuilder) {
    this.newAttackForm = this.fb.group({
      url: ['', [Validators.required, this.urlValidator]],
      method: ['GET', [Validators.required]],
      data: ['', [this.validateJson]],
      count: [100, [Validators.required, Validators.min(1)]],
      threads: [1, [Validators.required, Validators.min(1)]],
      queryParams: ['', [this.queryParamsValidator]],
      proxy: [''],
    });
  }
  ngOnInit(): void {
    this.authService.isUserLoggedIn$.subscribe((isLoggedIn) => {
      this.isLoggedIn = isLoggedIn;
    });
  }
  logout(): void {
    this.authService.logout();
  }

  createAttack() {
    if (this.newAttackForm.status === 'VALID') {
      console.log(this.buildSpamCommand(this.newAttackForm.value))
    }
  }

  openForm() {
    this.isFormAttackVisible = true;
  }

  closeForm() {
    this.isFormAttackVisible = false;
    this.newAttackForm.reset();
    console.log('close')
  }

  urlValidator(control) {
    const value = control.value;
    if (!value) {
      return null;
    }
    const urlPattern = /^(https?|http):\/\/([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}(:\d+)?(\/\S*)?$/;
    return urlPattern.test(value) ? null : { 'invalidUrl': { value } };
  }

  validateJson(control) {
    const value = control.value || '';
    if (value) {
      try {
        JSON.parse(value);
        return null;
      } catch (error) {
        return { invalidJson: { value: control.value } };
      }
    }
    return null
  }

  queryParamsValidator(control) {
    const value = control.value || '';
   if (value) {
     const valid = /^[?&A-Za-z0-9_=%-]+$/.test(value);
     return valid ? null : { invalidQueryParams: { value: control.value } };
   }
   return null
  }

  buildSpamCommand(attackData): string {
    let command = `site=${attackData.url} method=${attackData.method}`;

    if (attackData.data) {
      const dataString = JSON.stringify(attackData.data);
      command += ` data=${dataString}`;
    }

    if (attackData.count) {
      command += ` count=${attackData.count}`;
    }

    if (attackData.queryParams) {
      command += ` queryParams=${attackData.queryParams}`;
    }

    if (attackData.proxy) {
      command += ` proxy=${attackData.proxy}`;
    }

    if (attackData.threads) {
      command += ` proxy=${attackData.proxy}`;
    }

    return command;
  }

}
