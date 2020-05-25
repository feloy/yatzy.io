import { Injectable } from '@angular/core';
import { AngularFirestore } from '@angular/fire/firestore';

import { filter, take, mergeMap, tap, first, map } from 'rxjs/operators';
import { Observable } from 'rxjs';

import { User } from '../models/user';
import { Room } from '../models/room';
import { Position } from '../models/position';
import { AuthenticatedUser } from '../models/authenticated-user';
import { Board } from '../models/board';

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

  listenBoard(roomId: string): Observable<Board[]> {
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

  listenDie(docId: string): Observable<any> {
    return this.db.collection<User>('users').doc(docId).valueChanges()
      .pipe(
        filter((usr: User) => usr && 'die' in usr && !('replay' in usr)),
        map((usr: User) => ({
          die: usr.die.reduce((acc, val, ind) => [...acc, { dice: val, i: ind }], []),
          shots: usr.shots
        })
        )
      );
  }

  replayDie(docId: string, positions: number[]) {
    this.db.collection<User>('users').doc(docId).update({
      replay: positions
    });
  }

  listenFinish(docId: string): Observable<boolean> {
    return this.db.collection<User>('users').doc(docId).valueChanges()
      .pipe(
        filter((user: User) => user && 'finish' in user && user['finish'] === true),
        first(),
        map(() => true)
      );
  }
}
