import { Component } from '@angular/core';
import { AuthenticatedUser } from './components/login/login.component';

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
