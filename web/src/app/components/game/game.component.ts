import { Component, OnInit, Input } from '@angular/core';
import { RoomConfig } from '../join-room/join-room.component';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {

  // Firebase auth user
  @Input() user: firebase.User;

  roomConfig: RoomConfig;

  constructor() { }

  ngOnInit() {
  }

  configChange(config: RoomConfig) {
    this.roomConfig = config;
  }
}
