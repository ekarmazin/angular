import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-robot',
  templateUrl: './robot.component.html',
  styleUrls: ['./robot.component.css']
})
export class RobotComponent implements OnInit {

  items: any = [];
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
  }

}
