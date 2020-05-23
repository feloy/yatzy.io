# Yatzy.io: An online and multi-player yatzy game

## Web

Create a Firebase project and add a web application.

### Authentication

[Use anonymous authentication for the web](https://firebase.google.com/docs/auth/web/anonymous-auth)

Get the Firebase SDK Snippet Configuration and create a `web/src/firebase-config.ts` file containing something similar to:

```javascript
export const firebaseConfig = {
    apiKey: 'your-firebase-api-key',
    authDomain: 'your-project-id.firebaseapp.com',
    databaseURL: "https://your-project-id.firebaseio.com",
    projectId: 'your-firebase-project-id',
    [...]
};

```
