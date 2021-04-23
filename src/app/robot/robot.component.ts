import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-robot',
  templateUrl: './robot.component.html',
  styleUrls: ['./robot.component.css']
})



export class RobotComponent implements OnInit {

  items: any = [];
  userEmail= '';
  cron:  any = [];
  constructor(private http: HttpClient) {
    this.http.get(environment.apiEndPoint + "/s3/logs").toPromise().then(data => {
      this.items = data;
    });

    this.http.get(environment.apiEndPoint + "/schedule/cron").toPromise().then(data => {
      this.cron = data;
    });

  }

  actionMethod(event: any) {
    event.target.disabled = true;
  }

  ngOnInit(): void {
    this.userEmail = sessionStorage.getItem('loggedUser');
  }

  onCreatePost() {
    // Send Http request
    this.http.post(environment.apiEndPoint + "/run/manual",
      {
        "email": this.userEmail ,
      })
      .subscribe(
        (val) => {
          console.log("POST call successful value returned in body",
            val);
        },
        response => {
          console.log("POST call in error", response);
        },
        () => {
          console.log("The POST observable is now completed.");
        });
  }
}

