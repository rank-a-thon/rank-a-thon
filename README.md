# Rankathon

Smart, algorithm powered hackathons. Watch [this video](https://www.youtube.com/watch?v=QkSwkuQ1q-U)
and read [this poster](https://i.ibb.co/1GrSnmq/2001.jpg) for an idea of what
this project is about.

## Setting Up for Development

1. Install [Docker](https://www.docker.com/get-started) and
   [docker-compose](https://docs.docker.com/compose/install/).

1. Rename `.env.example` to `.env` in the folder root and fill it up with the
   appropriate secrets and environment variables.

1. Run `docker-compose up --build`.

1. Voila! It should be working! Visit <http://localhost:5555> for the client and
   <http://localhost:5555/api> for the backend.

Note that both the client and backend have hot-reload working. Any changes you
make will (mostly) be hot-reloaded in a few seconds.

## Deployment

TODO: Write this

## Editor Setup for VS Code

1. Obtain the following extensions:

- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode): Code formatter
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint): JS linter, configured via `.eslintrc.js`

2. Update your workspace's `settings.json` (Preferences: Open User Settings > Open settings.json):

```
{
  // Let eslint and prettier format code on save
  "editor.formatOnSave": true,
}
```

3. Reload your editor
