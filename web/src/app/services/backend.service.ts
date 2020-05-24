import { Injectable } from '@angular/core';
import { AuthenticatedUser } from '../components/login/login.component';
import { AngularFirestore, DocumentSnapshot, DocumentData } from '@angular/fire/firestore';
import { filter, take, mergeMap, tap, first, map } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { Board } from '../components/grid/grid.component';

export interface User {
  name?: string;
  size?: number; // number of players required
  tokenId: string;
  room?: string;
  die?: number[];
  shots?: 0 | 1 | 2;
}

export interface Room {
  full?: boolean;
  users?: string[];
  board?: string;
}

export interface Player {
  id?: string;
  bot?: boolean;
  name?: string;
}

export interface Position {
  x: number;
  y: number;
}

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor(public db: AngularFirestore) { }

  // joins user to a new room of size roomSize
  join(user: AuthenticatedUser, roomSize: number): Promise<string> {
    return this.db.collection<User>('users').add({
      name: user.name,
      size: roomSize,
      tokenId: user.token
    }).then((doc: firebase.firestore.DocumentReference) => {
      return doc.id;
    });
  }

  waitRoomComplete(userId: string): Observable<any> {
    return this.db.collection<User>('users').doc(userId).valueChanges()
      .pipe(
        filter((userDoc: User) => userDoc && 'room' in userDoc),
        take(1),
        mergeMap((userDoc: User) => {
          // user is affected to a room
          const roomID = userDoc.room;
          return this.db.collection('rooms').doc(roomID).collection('players').stateChanges()
            .pipe(
              mergeMap(actions => {
                return actions.map(action => {
                  return { name: action.payload.doc.data().name, id: action.payload.doc.id };
                });
              })
            );
        })
      );
  }

  getUser(docId: string): Observable<User> {
    return this.db.collection<User>('users').doc(docId).valueChanges()
    .pipe(
      filter((userDoc: User) => userDoc && 'room' in userDoc),
      first()
    );
  }

  listenBoard(roomId: string): Observable<Board> {
    return this.db.collection<Room>('rooms').doc(roomId).valueChanges()
    .pipe(
      filter((room: Room) => 'board' in room),
      map((room: Room) => JSON.parse(room.board))
    )
  }

  play(userId: string, pos: Position): Promise<void> {
    return this.db.collection<User>('users').doc(userId).update({
      click: pos
    });
  }
}
