import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import { HttpClientModule, HttpClientJsonpModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { AuthComponent } from './auth/auth.component';
import { RobotComponent } from './robot/robot.component';

const appRoutes: Routes = [
  { path: '', component: AppComponent}, // <-- default route to HOME page
  { path: 'auth', component: AuthComponent},
  { path: 'robot', component: RobotComponent},

]

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    RobotComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    HttpClientJsonpModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
