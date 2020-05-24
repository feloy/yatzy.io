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
    this.join().then((docId: string) => this.waitRoomComplete(docId));
  }

  join(): Promise<string> {
    return this.backend.join(this.me, this.roomConfig.roomSize);
  }

  waitRoomComplete(docId: string) {
    const sub = this.backend.waitRoomComplete(docId)
    .subscribe(
      (name: string) => {
        this.waitingPlayers.push(name);
        if (this.waitingPlayers.length == this.roomConfig.roomSize) {
          sub.unsubscribe();
          console.log("COMPLETE")
        }
      });
  }
}
