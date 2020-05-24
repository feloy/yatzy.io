import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-dice',
  templateUrl: './dice.component.html',
  styleUrls: ['./dice.component.css']
})
export class DiceComponent implements OnInit {

  @Input() dice: { dice: number, i: number };
  @Input() initKeep;
  @Input() set waiting(value: boolean) {
    this._waiting = value;
    if (value) {
      this.startAnim();
    } else {
      clearInterval(this.animId);
      this.anim = 0;
    }
  }
  @Output() keep = new EventEmitter();

  public _waiting: boolean;
  public _keep;
  public anim = 0;
  private animId;

  constructor() {
  }

  ngOnInit() {
    this._keep = this.initKeep;
  }

  onClicked() {
    this._keep = !this._keep;
    this.keep.emit(this._keep);
  }

  private startAnim() {
    setTimeout(() => {
      this.anim = 1;
      this.animId = setInterval(() => {
        this.anim = (this.anim + 1) % 5 + 1;
      }, 250);
    }, 100 * Math.random());
  }
}
