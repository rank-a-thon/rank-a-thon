import firebase from 'firebase/app';
import 'firebase/auth';

try {
  // Cannot destructure here
  // https://nextjs.org/docs#build-time-configuration
  firebase.initializeApp({
    apiKey: process.env.firebaseApiKey,
    authDomain: process.env.firebaseAuthDomain,
    projectId: process.env.firebaseProjectId,
  });
} catch (err) {
  // https://github.com/NickDelfino/nextjs-with-redux-and-cloud-firestore/blob/49c9cd560318508c03ee6aa8c6a90174b41f4e22/lib/db.js#L16
  // Ignore the "already exists" message which is not an actual error when
  // we're hot-reloading.
  if (!/already exists/.test(err.message)) {
    console.error('Firebase initialization error', err.stack);
  }
}

const sendEmailVerificationToUser = (user) => {
  return user
    .sendEmailVerification({
      url: `${process.env.THIS_URL}/register`,
    })
    .then(() => {
      return true;
    });
};

// const facebookAuthProvider = new firebase.auth.FacebookAuthProvider();
const googleAuthProvider = new firebase.auth.GoogleAuthProvider();
// const githubAuthProvider = new firebase.auth.GithubAuthProvider();
const auth = firebase.auth;

export { googleAuthProvider, auth, sendEmailVerificationToUser };
