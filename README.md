![alt text](https://raw.githubusercontent.com/rank-a-thon/rank-a-thon/master/web/public/img/logo.svg "Rankathon Logo")

# Rankathon

Smart, algorithm powered hackathons. Watch [this video](https://www.youtube.com/watch?v=QkSwkuQ1q-U)
and read [this poster](https://i.ibb.co/1GrSnmq/2001.jpg) for an idea of what
this project is about.

This project is LIVE and deployed on http://rankathon.io.

# Orbital Milestone Submissions

To view our Orbital Milestone 3 submission, [click here](https://hackmd.io/@sunyitao/rankathon-3).

## Setting Up for Development

1. Install [Docker](https://www.docker.com/get-started) and
   [docker-compose](https://docs.docker.com/compose/install/).

1. Rename `.env.example` to `.env` in the folder root and fill it up with the
   appropriate secrets and environment variables.

1. Run `docker-compose up --build`.

1. Enter the PostgreSQL container with

   ```bash
   $ docker exec -ti postgres /bin/sh
   ```

1. Now restart docker-compose by pressing Ctrl+C in your docker-compose shell, then running

   ```bash
   $ docker-compose up
   ```

1. Voila! It should be working! Visit <http://localhost:5555> for the client and
   <http://localhost:5555/api> for the backend.

Note that both the client and backend have hot-reload working. Any changes you
make will be hot-reloaded in a few seconds.

## Deployment to Production

1. SSH to your server / use your local machine.

1. Install [Docker](https://www.docker.com/get-started),
   [docker-compose](https://docs.docker.com/compose/install/), Git, [tmux](https://github.com/tmux/tmux/wiki/Installing) with your package manager of choice.

1. Rename `.env.example` to `.env` in the folder root and fill it up with the appropriate secrets and environment variables.

1. In your shell, run

   ```bash
   $ tmux
   ```

   This starts a `tmux` session, which allows you to detach from the running server without the containers terminating when you are done deploying.

   All commands run afterwards should be run in `tmux`.

1) Run `./deploy-prod.sh`. The containers will build and PostgreSQL will throw some database errors.

1) Start a new `tmux` terminal by pressing Ctrl+B then C.

1) Enter the PostgreSQL container with

   ```bash
   $ docker exec -ti postgres /bin/sh
   ```

1) Go back to your docker-compose shell by pressing Ctrl+B then 0.

1) Now restart docker-compose by pressing Ctrl+C in your docker-compose shell, then running

   ```bash
   $ ./deploy-prod.sh
   ```

1) Voila! It should be working! Visit either <http://localhost> or your server's URL and Rankathon should be up!

Note that this creates a production build of Rankathon and is not suitable for development purposes.

## Running Tests

1. `cd api`
2. `sh test_all.sh`

## Golang Backend Documentation

1. Run `godoc -http=:6060`
2. Navigate to `http://localhost:6060/pkg/github.com/rank-a-thon/rank-a-thon/api/`

## Editor Setup for VS Code

1. Obtain the following extensions:

- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode): Code formatter
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint): JS linter, configured via `.eslintrc.js`

2. Update your workspace's `settings.json` (Preferences: Open User Settings > Open settings.json):

```json
{
  // Let eslint and prettier format code on save
  "editor.formatOnSave": true
}
```

3. Reload your editor

## License

This project uses [Apache License 2.0](https://github.com/rank-a-thon/rank-a-thon/blob/master/LICENSE).
