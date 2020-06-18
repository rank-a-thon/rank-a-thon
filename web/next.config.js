module.exports = {
  env: {
    // Firebase
    firebaseApiKey: process.env.FIREBASE_API_KEY,
    firebaseAuthDomain: process.env.FIREBASE_AUTH_DOMAIN,
    firebaseProjectId: process.env.FIREBASE_PROJECT_ID,

    API_URL: process.env.API_URL || 'http://localhost:5555/api/',

    THIS_URL: process.env.THIS_URL || 'http://localhost:5555/api/',
  },
};
