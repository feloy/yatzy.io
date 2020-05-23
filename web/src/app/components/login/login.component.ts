import { Component, OnInit, Output, EventEmitter, AfterViewInit } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/auth';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  @Output('userChange') msg = new EventEmitter<firebase.User>();

  constructor(public auth: AngularFireAuth) { }

  ngOnInit() {
    const that = this; 
    this.auth.auth.onAuthStateChanged(function onStateChanged(firebaseUser: firebase.User) {
      that.msg.emit(firebaseUser);
    });
  }

  login() {
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
