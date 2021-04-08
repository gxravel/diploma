import axios from 'axios';
import qs from 'qs';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function auth(login, password) {
  return axios.post(
    'auth',
    qs.stringify({
      login: login,
      password: password,
    }),
    options,
  );
}

export function register(login, password) {
  return axios.post(
    'register',
    qs.stringify({
      login: login,
      password: password,
    }),
    options,
  );
}

export function logout(token, handleChangeToken) {
  axios.delete('logout/' + token).then(() => {
    handleChangeToken('', false);
  });
}
