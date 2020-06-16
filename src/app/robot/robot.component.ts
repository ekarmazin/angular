import { Component, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";

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
    this.http.get("https://data.assets.staging.sweet.io/s3/logs").toPromise().then(data => {
      this.items = data;
    });

    this.http.get("https://data.assets.staging.sweet.io/schedule/cron").toPromise().then(data => {
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
    const headers = new HttpHeaders()
      .set("Content-Type", "application/json");

    // Send Http request
    this.http.post("https://data.assets.staging.sweet.io/run/manual",
      {
        "email": this.userEmail ,
      },
      {headers})
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

