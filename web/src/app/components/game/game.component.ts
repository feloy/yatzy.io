import { Component, OnInit, Input } from '@angular/core';
import { RoomConfig } from '../join-room/join-room.component';
import { AuthenticatedUser } from '../login/login.component';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {

  // Firebase auth user
  @Input() user: AuthenticatedUser;

  roomConfig: RoomConfig;

  constructor() { }

  ngOnInit() {
  }

  configChange(config: RoomConfig) {
    this.roomConfig = config;
  }
}
