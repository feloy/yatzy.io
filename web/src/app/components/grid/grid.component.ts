import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

import { formulas } from 'src/app/values/formulas';

import { Board } from 'src/app/models/board';
import { Hexagon } from 'src/app/models/hexagon';

@Component({
  selector: 'app-grid',
  templateUrl: './grid.component.html',
  styleUrls: ['./grid.component.css']
})
export class GridComponent implements OnInit {

  private RADIUS = 3;

  @Input()
  set board(value: Board[]) {
    if (value && value.length > 0) {
      this.compute(value);
    }
  }
  @Input() usermap: { string: number };
  @Input() me: number;

  @Output() selected = new EventEmitter();

  public w: number;
  public h: number;
  public hexagons: Hexagon[] = [];
  public index: number[][] = [];

  constructor() { }

  ngOnInit() {
    this.w = (2 * this.RADIUS + 1) * 50;
    this.h = (2 * this.RADIUS + 1) * 42.42 + 14.42;
    this.compute([]);
  }

  private compute(board: Board[]) {
    let p = 0;
    this.index = [];
    this.hexagons = [];
    for (let y = this.RADIUS; y >= -this.RADIUS; y--) {
      this.index[y] = [];
      const Y = Math.abs(y);
      for (let x = -this.RADIUS + Math.floor(Y / 2); x <= this.RADIUS - Math.floor((Y + 1) / 2); x++) {
        this.hexagons.push({ x, y, p, label: '', colorIndex: -1, clickable: false, played: false });
        this.index[y][x] = p;
        p++;
      }
    }
    if (board) {
      board.map((b: Board) => {
        const P = this.index[b.y][b.x];
        if ('points' in b) {
          this.hexagons[P].label = '+' + b.points;
          this.hexagons[P].clickable = false;
          this.hexagons[P].played = true;
        } else {
          this.hexagons[P].label = formulas[b.formula];
          this.hexagons[P].clickable = (this.usermap[b.userId]  === this.me);
        }
        this.hexagons[P].colorIndex = this.usermap[b.userId];
      });
    }
  }

  onSelected(p: number) {
    if (this.hexagons[p].colorIndex !== this.me || !this.hexagons[p].clickable) {
      return;
    }
    this.hexagons[p].label = '...';
    this.selected.emit({x: this.hexagons[p].x, y: this.hexagons[p].y});
  }
}
