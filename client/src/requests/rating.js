import qs from 'qs';
import axios from 'axios';

axios.defaults.headers.put['Content-Type'] =
  'application/x-www-form-urlencoded';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function getRating(id, token) {
  return axios
    .get(
      `book/rating/${id}`,
      {
        params: {
          token: token,
        },
      },
      options,
    )
    .then(data => {
      if (!data.error) {
        return data.rating;
      }
    });
}

export function addRating(userValue, bookId, token) {
  return axios
    .post(
      `book/rating/${bookId}`,
      qs.stringify({
        user_value: userValue,
        token: token,
      }),
      options,
    );
}

export function editRating(userValue, bookId, token) {
  return axios
    .put(
      `book/rating/${bookId}`,
      qs.stringify({
        user_value: userValue,
        token: token,
      }),
      options,
    );
}
