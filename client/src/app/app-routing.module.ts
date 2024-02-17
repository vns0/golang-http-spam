import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthGuardService} from './services/auth-guard.service';
import { LoginComponent } from "./pages/login/login.component";
import { DashboardComponent } from "./pages/dashboard/dashboard.component";
import { AuthRedirectGuard } from "./services/auth-redirect-guard.service";
import { UsersComponent } from "./pages/users/users.component";

const routes: Routes = [
  {
    path: "login",
    component: LoginComponent,
    canActivate: [AuthRedirectGuard]
  },
  {
    path: "dashboard",
    component: DashboardComponent ,
    canActivate: [AuthGuardService]
  },
  {
    path: "users",
    component: UsersComponent ,
    canActivate: [AuthGuardService]
  },
  {
    path: "**",
    redirectTo: "/login"
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
