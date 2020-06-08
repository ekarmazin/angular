import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule, Routes } from "@angular/router";
import { HttpClientModule, HttpClientJsonpModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { AuthComponent } from './auth/auth.component';
import { RobotComponent } from './robot/robot.component';
import { HeaderComponent } from './header/header.component';
import { AuthGuard } from './auth/auth.guard';

const appRoutes: Routes = [
  { path: '', redirectTo: '/auth', pathMatch: 'full' }, // <-- default route to HOME page
  { path: 'auth', component: AuthComponent},
  { path: 'robot', canActivate: [AuthGuard],  component: RobotComponent},
]

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    RobotComponent,
    HeaderComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    HttpClientJsonpModule,
    FormsModule,
    ReactiveFormsModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
