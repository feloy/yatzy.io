import { Component, OnInit, Input } from '@angular/core';
import { RoomConfig } from '../join-room/join-room.component';

@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css']
})
export class BoardComponent implements OnInit {

  @Input('roomConfig') roomConfig: RoomConfig;

  public waitingPlayers: string[] = [];

  constructor() { }

  ngOnInit() {
    this.waitingPlayers = Array(this.roomConfig.roomSize).fill("...");
    this.waitingPlayers[0] = this.roomConfig.me;
  }
}
