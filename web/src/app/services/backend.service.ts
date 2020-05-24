import { Injectable } from '@angular/core';
import { AuthenticatedUser } from '../components/login/login.component';
import { AngularFirestore } from '@angular/fire/firestore';
import { filter, take, takeWhile, mergeMap } from 'rxjs/operators';
import { Observable } from 'rxjs';

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

interface Player {
  bot?: boolean;
  name?: string;
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
                return actions.map(action => action.payload.doc.data().name);
              })
            );
        })
      );
  }
}
