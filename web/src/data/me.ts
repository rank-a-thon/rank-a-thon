// Hacky way to handle state FOR NOW pls no judge. Will transit to using redux soon to manage auth

// TODO: use redux for Auth
// TODO: use cookies to store JWT

type JWTToken = {
  access_token: string;
  refresh_token: string;
};

type User = {
  ID: number;
  email: string;
  name: string;
  user_type: number;
  team_id: number;
};

export type Me = JWTToken & User;

export function saveMe(token: JWTToken, user: User) {
  window.localStorage.setItem('token', JSON.stringify(token));
  window.localStorage.setItem('user', JSON.stringify(user));
}

export function getMe(): Me | null {
  if (!localStorage.getItem('token')) {
    return null;
  } else {
    return {
      ...JSON.parse(window.localStorage.getItem('token')),
      ...JSON.parse(window.localStorage.getItem('user')),
    };
  }
}

export function clearMe() {
  window.localStorage.clear();
}
