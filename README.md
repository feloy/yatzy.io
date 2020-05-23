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

Copy the Firebase configuration files to a bucket for CI/CD:

```shell
gsutil mb gs://$GOOGLE_CLOUD_PROJECT
gsutil cp \
    web/src/firebase-config.ts \
    web/firebase.json \
    web/.firebaserc \
    gs://$GOOGLE_CLOUD_PROJECT/firebase-config.ts
```

Build the firebase Docker image needed for the web build:

```shell
gcloud builds submit --config=tools/firebase/cloudbuild.yaml tools/firebase
```

Build and deploy the web app:

```shell
gcloud builds submit --config=web/cloudbuild.yaml web
```
