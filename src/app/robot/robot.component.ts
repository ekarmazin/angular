import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {AuthComponent} from "../auth/auth.component";
import {message} from "aws-sdk/clients/sns";

@Component({
  selector: 'app-robot',
  templateUrl: './robot.component.html',
  styleUrls: ['./robot.component.css']
})
export class RobotComponent implements OnInit {

  items: any = [];
  userEmail= '';
  constructor(private http: HttpClient) {
    this.http.get("https://data.assets.staging.sweet.io/s3").toPromise().then(data => {
      this.items = data;
    });


  }
  ngOnInit(): void {
    this.userEmail = sessionStorage.getItem('loggedUser');
    console.log(sessionStorage.getItem('loggedUser'));
  }

  onCreatePost(postData: { content: string }) {
    // Send Http request
    this.http
      .post(
        'https://test-5a7bf.firebaseio.com/posts.json',
        postData
      )
      .subscribe(responseData => {
        console.log(responseData);

      });

  }
}

