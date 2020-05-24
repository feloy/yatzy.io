import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-die',
  templateUrl: './die.component.html',
  styleUrls: ['./die.component.css']
})
export class DieComponent implements OnInit {

  @Input()
  set die(value: { die: { dice: number, i: number }[], shots: (0 | 1 | 2) }) {
    if (value) {
      this._die = value.die;
      this._shots = value.shots;
      if (value.shots === 2) {
        this.keep = [false, false, false, false, false];
      }
      this.waiting = [false, false, false, false, false];
    }
  }

  @Output() replay = new EventEmitter();

  _shots: 0 | 1 | 2;
  _die: { dice: number, i: number }[];

  keep: boolean[] = [false, false, false, false, false];
  waiting: boolean[] = [false, false, false, false, false];

  constructor() { }

  ngOnInit() {
  }

  onKeep(i: number, keep: boolean) {
    this.keep[i] = keep;
  }

  onReplay() {
    const replay = this.keep.reduce((acc, keep, ind) => keep ? acc : [...acc, ind], []);
    this.replay.emit(replay);
    this.waiting = [false, false, false, false, false];
    for (let i = 0; i < this._die.length; i++) {
      if (replay.indexOf(i) >= 0) {
        this.waiting[i] = true;
      }
    }
  }
}
