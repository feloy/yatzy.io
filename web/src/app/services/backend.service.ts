import { Injectable } from '@angular/core';
import { AuthenticatedUser } from '../components/login/login.component';
import { AngularFirestore } from '@angular/fire/firestore';

export interface User {
  name?: string;
  size?: number; // number of players required
  tokenId: string;
  room?: string;
  die?: number[];
  shots?: 0 | 1 | 2;
}

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor(public db: AngularFirestore) { }

  // joins user to a new room of size roomSize
  join(user: AuthenticatedUser, roomSize: number) {
    this.db.collection<User>('users').add({
      name: user.name,
      size: roomSize,
      tokenId: user.token
    }).then((doc: firebase.firestore.DocumentReference) => {
      console.log("docId", doc.id);
    });
  }
}
