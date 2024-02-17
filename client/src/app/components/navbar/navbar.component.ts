import { Component } from '@angular/core';
import {AuthService} from "../../services/auth.service";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent {
  isLoggedIn: boolean = false;
  constructor(private authService: AuthService) { }

  ngOnInit(): void {
    this.authService.isUserLoggedIn$.subscribe((isLoggedIn) => {
      this.isLoggedIn = isLoggedIn;
    });
  }
  menuItems = [
    { path: '/dashboard', label: 'Dashboard' },
    { path: '/users', label: 'Users' },
    // Добавьте другие пункты меню
  ];
}
