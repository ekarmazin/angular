import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';

// import { DataStorageService } from '../shared/data-storage.service';
import { AuthService } from '../auth/auth.service';
import { AuthComponent } from '../auth/auth.component';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit, OnDestroy {
  isAuthenticated = false;
  private userSub: Subscription;
  userEmail= '';

  constructor(
    // private dataStorageService: DataStorageService,
    private authService: AuthService,
    // private authComponent: AuthComponent
  ) {}



  ngOnInit() {
    this.userEmail = sessionStorage.getItem('loggedUser');
    this.userSub = this.authService.user.subscribe(user => {
      this.isAuthenticated = !!user;
      console.log(!user);
      console.log(!!user);
    });
    // this.userEmail = this.authComponent.username;
    // console.log("done ", this.userEmail);
  }


  ngOnDestroy() {
    this.userSub.unsubscribe();
  }
}
