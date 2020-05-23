import { Component, OnInit, Output, EventEmitter } from '@angular/core';

export interface RoomConfig {
  // my nickname
  me: string;
  // Total number of players
  roomSize: number;
}

@Component({
  selector: 'app-join-room',
  templateUrl: './join-room.component.html',
  styleUrls: ['./join-room.component.css']
})
export class JoinRoomComponent implements OnInit {

  @Output('configChange') msg = new EventEmitter<RoomConfig>();

  name: string;
  size: string;

  constructor() { }

  ngOnInit() {
  }

  play() {
    this.msg.emit({
      me: this.name, 
      roomSize: parseInt(this.size, 10)
    })
  }
}
