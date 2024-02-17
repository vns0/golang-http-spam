import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthGuardService} from './services/auth-guard.service';
import { LoginComponent } from "./pages/login/login.component";
import { DashboardComponent } from "./pages/dashboard/dashboard.component";
import { AuthRedirectGuard } from "./services/auth-redirect-guard.service";

const routes: Routes = [
 {
   path: "dashboard",
   component: DashboardComponent ,
   canActivate: [AuthGuardService]
 },
 {
   path: "login",
   component: LoginComponent,
   canActivate: [AuthRedirectGuard]
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
