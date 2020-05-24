import { Component, OnInit, Input } from '@angular/core';
import { BackendService, User, Player, Position } from 'src/app/services/backend.service';
import { Board } from '../grid/grid.component';

@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css']
})
export class BoardComponent implements OnInit {

  @Input('myID') myID: string;
  @Input('players') players: Player[];

  private user: User;
  public board: Board;
  public usermap: {};
  public me: number;

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

    console.log("usermap", this.usermap);
    console.log("me", this.me);
    console.log("myId", this.myID);

    this.backend.getUser(this.myID).subscribe((user: User) => {
      this.user = user;
      this.backend.listenBoard(this.user.room).subscribe((board: Board) => {
        this.board = board;
        console.log("new board", board);
      });
    })
  }

  onSelected(pos: Position) {
    this.backend.play(this.myID, pos);
  }
}
