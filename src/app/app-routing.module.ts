import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {AuthComponent} from "./auth/auth.component";
import {AppComponent} from "./app.component";
import {RobotComponent} from "./robot/robot.component";

const appRoutes: Routes = [
  { path: '', component: AppComponent},
  { path: 'robot', component: RobotComponent},
  { path: 'auth', component: AuthComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(appRoutes)],
  exports: [RouterModule]
})

export class AppRoutingModule {

}
