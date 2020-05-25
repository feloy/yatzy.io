import { Component, OnInit, Input } from '@angular/core';

import { AuthenticatedUser } from 'src/app/models/authenticated-user';
import { RoomConfig } from 'src/app/models/room-config';

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
