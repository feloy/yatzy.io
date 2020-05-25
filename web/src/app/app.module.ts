import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { AngularFireAuthModule } from '@angular/fire/auth';
import { AngularFireModule } from '@angular/fire';
import { firebaseConfig } from 'src/firebase-config';
import { WaitRoomComponent } from './components/wait-room/wait-room.component';
import { GameComponent } from './components/game/game.component';
import { JoinRoomComponent } from './components/join-room/join-room.component';
import { AngularFirestoreModule } from '@angular/fire/firestore';
import { BoardComponent } from './components/board/board.component';
import { GridComponent } from './components/grid/grid.component';
import { HexagonComponent } from './components/hexagon/hexagon.component';
import { DieComponent } from './components/die/die.component';
import { DiceComponent } from './components/dice/dice.component';
import { RoomInfoComponent } from './components/room-info/room-info.component';
import { YatzyScorePipe } from './pipes/yatzy-score.pipe';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    WaitRoomComponent,
    GameComponent,
    JoinRoomComponent,
    BoardComponent,
    GridComponent,
    HexagonComponent,
    DieComponent,
    DiceComponent,
    RoomInfoComponent,
    YatzyScorePipe
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AngularFireModule.initializeApp(firebaseConfig),
    AngularFireAuthModule,
    AngularFirestoreModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
