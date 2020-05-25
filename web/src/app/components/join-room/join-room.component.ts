import { Component, OnInit, Output, EventEmitter } from '@angular/core';

import { RoomConfig } from 'src/app/models/room-config';

@Component({
  selector: 'app-join-room',
  templateUrl: './join-room.component.html',
  styleUrls: ['./join-room.component.css']
})
export class JoinRoomComponent implements OnInit {

  @Output('configChange') msg = new EventEmitter<RoomConfig>();

  size: string;

  constructor() { }

  ngOnInit() {
  }

  play() {
    let size = parseInt(this.size, 10);
    let botsInvites = 0;
    if (size > 100) {
      size -= 100;
      botsInvites = size - 1;
    }
    this.msg.emit({
      roomSize: size,
      botsInvites: botsInvites
    })
  }
}
