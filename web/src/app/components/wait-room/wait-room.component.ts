import { Component, OnInit, Input } from '@angular/core';

import { BackendService } from 'src/app/services/backend.service';

import { RoomConfig } from 'src/app/models/room-config';
import { AuthenticatedUser } from 'src/app/models/authenticated-user';

@Component({
  selector: 'app-wait-room',
  templateUrl: './wait-room.component.html',
  styleUrls: ['./wait-room.component.css']
})
export class WaitRoomComponent implements OnInit {

  @Input('roomConfig') roomConfig: RoomConfig;
  @Input('user') me: AuthenticatedUser;

  public waitingPlayers: any[] = [];

  public myID: string;
  public complete = false;

  constructor(public backend: BackendService) { }

  ngOnInit() {
    this.join().then((docId: string) => {
      this.myID = docId;
      this.waitRoomComplete(docId);
    });
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
          this.complete = true;
        }
      });
  }
}
