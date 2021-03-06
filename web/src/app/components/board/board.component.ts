import { Component, OnInit, Input } from '@angular/core';

import { Observable } from 'rxjs';

import { BackendService } from 'src/app/services/backend.service';

import { User } from 'src/app/models/user';
import { Position } from 'src/app/models/position';
import { Player } from 'src/app/models/player';
import { Board } from 'src/app/models/board';

@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css']
})
export class BoardComponent implements OnInit {

  @Input('myID') myID: string;
  @Input('players') players: Player[];

  private user: User;
  public board: Board[];
  public usermap: {};
  public me: number;

  public die$: Observable<{ die: { dice: number, i: number }[], shots: (0 | 1 | 2) }>;
  public finish: boolean = false;

  constructor(public backend: BackendService) { }

  ngOnInit() {
    this.usermap = {};
    this.players.map((usr: Player) => this.usermap[usr.id] = this.players.indexOf(usr));

    for (var i = 0; i < this.players.length; i++) {
      if (this.players[i].id == this.myID) {
        this.me = i;
        break;
      }
    }

    this.backend.getUser(this.myID).subscribe((user: User) => {
      this.user = user;
      this.backend.listenBoard(this.user.room).subscribe((board: Board[]) => {
        this.board = board;
      });
    });

    this.die$ = this.backend.listenDie(this.myID);

    this.backend.listenFinish(this.myID).subscribe(() => this.finish = true);
  }

  onSelected(pos: Position) {
    this.backend.play(this.myID, pos);
  }

  onReplay(positions: number[]) {
    this.backend.replayDie(this.myID, positions);
  }

  userScore(score: number) {
    // TODO send to leaderboard
  }
}
