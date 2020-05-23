import { Component, OnInit, Input } from '@angular/core';
import { RoomConfig } from '../join-room/join-room.component';
import { AuthenticatedUser } from '../login/login.component';
import { BackendService } from 'src/app/services/backend.service';

@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css']
})
export class BoardComponent implements OnInit {

  @Input('roomConfig') roomConfig: RoomConfig;
  @Input('user') me: AuthenticatedUser;

  public waitingPlayers: string[] = [];

  constructor(public backend: BackendService) { }

  ngOnInit() {
    this.waitingPlayers = Array(this.roomConfig.roomSize).fill("...");
    this.waitingPlayers[0] = this.me.name;

    this.join();
  }

  join() {
    this.backend.join(this.me, this.roomConfig.roomSize);
  }
}
