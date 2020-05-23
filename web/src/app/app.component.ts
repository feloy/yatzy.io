import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  user: firebase.User;

  userChange(user: firebase.User) {
    this.user = user;
  }
}
