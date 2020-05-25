import { Component } from '@angular/core';

import { AuthenticatedUser } from './models/authenticated-user';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  user: AuthenticatedUser;

  userChange(user: AuthenticatedUser) {
    this.user = user;
  }
}
