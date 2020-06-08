import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-robot',
  templateUrl: './robot.component.html',
  styleUrls: ['./robot.component.css']
})
export class RobotComponent implements OnInit {
  items: any = [];

  constructor(private http: HttpClient) {


    this.http.get("https://data.karmazin.me/s3").toPromise().then(data => {
      console.log(data);
      this.items = data;
      // console.log(this.products[0]);
    });


  }
  ngOnInit(): void {
  }

}
