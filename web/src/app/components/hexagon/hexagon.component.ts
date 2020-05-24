import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

import { formulas, formulaNames } from '../grid/grid.component';

export const colors: string[] = [
  '#F44336',
  '#2196F3',
  '#FFEB3B',
  '#4CAF50',
  '#FF9800',
  '#795548'
];

export const colorsUnplayed: string[] = [
'#EF9A9A',
'#90CAF9',
'#FFF59D',
'#A5D6A7',
'#FFCC80',
'#BCAAA4'
];

@Component({
  selector: 'app-hexagon',
  templateUrl: './hexagon.component.html',
  styleUrls: ['./hexagon.component.scss']
})
export class HexagonComponent implements OnInit {
  @Input() p; // order for positioning
  @Input() x; // x pos in grid
  @Input() y; // y pos in grid
  @Input()
  set label(value: string) {
    this._label = value;
    this.title = formulaNames[formulas.indexOf(value)];
  } // content
  @Input() w; // grid width, for pos in px
  @Input() h; // grid height, for pos in px
  @Input() c; // color/user index
  @Input() played;

  @Output() selected = new EventEmitter();

  X: number; // x pos in px
  Y: number; // y pos in px
  public _label;
  public title = '';

  constructor() { }

  ngOnInit() {
    this.X = this.w / 2 - 25 + 50 * this.x + Math.abs(this.y % 2) * 25;
    this.Y = this.h / 2 - 28.84 - (57.7 - 15) * this.y - 43.3 * this.p;
  }

  onclick() {
    this.selected.emit();
  }
}
