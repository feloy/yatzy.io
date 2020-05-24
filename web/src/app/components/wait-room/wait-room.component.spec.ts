import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WaitRoomComponent } from './wait-room.component';

describe('WaitRoomComponent', () => {
  let component: WaitRoomComponent;
  let fixture: ComponentFixture<WaitRoomComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WaitRoomComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WaitRoomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
