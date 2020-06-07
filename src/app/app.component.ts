import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import 'rxjs/add/operator/map'

// export type Item = { "Keys": string };

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app-dev01';
  items: any = [];

  constructor(private http: HttpClient) {


    this.http.get("https://data.karmazin.me/s3").toPromise().then(data => {
      console.log(data);
      this.items = data;
      // console.log(this.products[0]);

    });


  }
}
