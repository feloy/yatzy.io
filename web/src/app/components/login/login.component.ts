import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/auth';

import { AuthenticatedUser } from 'src/app/models/authenticated-user';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  @Output('userChange') msg = new EventEmitter<AuthenticatedUser>();

  readonly nameKey = 'nickname';
  name: string;

  constructor(public auth: AngularFireAuth) { }

  ngOnInit() {
    this.name = window.localStorage.getItem(this.nameKey);
    const that = this; 
    this.auth.auth.onAuthStateChanged(function onStateChanged(firebaseUser: firebase.User) {
      if (firebaseUser) {
        firebaseUser.getIdToken(true).then((token: string) => {
          that.msg.emit({
            name: that.name,
            token: token
          });  
        });  
      } else {
        that.msg.emit(null);
      }
    });
  }

  login() {
    window.localStorage.setItem(this.nameKey, this.name);
    this.auth.auth.signInAnonymously().catch(function (error) {
      // Handle Errors here.
      var errorCode = error.code;
      var errorMessage = error.message;
      console.log("Error login", errorCode, errorMessage);
    });
  }

  logout() {
    this.auth.auth.signOut();
  }
}
