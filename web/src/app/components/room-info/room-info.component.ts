import { Component, OnInit, Input, Output, ChangeDetectionStrategy, EventEmitter } from '@angular/core';

import { Player } from 'src/app/models/player';
import { colors } from 'src/app/values/colors';
import { Board } from 'src/app/models/board';

@Component({
  selector: 'app-room-info',
  templateUrl: './room-info.component.html',
  styleUrls: ['./room-info.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class RoomInfoComponent implements OnInit {

  _board: Board[];

  @Input() set board(value: Board[]) {
    this._board = value;
    this.computeScores();
  }

  _players: Player[];

  @Input() set players(value: Player[]) {
    this._players = value;
    this.computeScores();
  }

  @Input() userId: string;
  @Output() userScore = new EventEmitter<number>();

  scores: { string?: number };
  colors = colors;

  constructor() { }

  ngOnInit() { }

  private computeScores() {
    if (!this._players) {
      return;
    }

    this.scores = {};
    this._players.forEach((player: Player) => {
      this.scores[player.id] = 0;
    });

    if (!this._board) {
      return;
    }

    this._board.forEach((board: Board) => {
      if (board.userId in this.scores && 'points' in board) {
        this.scores[board.userId] += board.points;
      }
    });

    if (this.userId in this.scores) {
      this.userScore.next(this.scores[this.userId]);
    }
  }
}
