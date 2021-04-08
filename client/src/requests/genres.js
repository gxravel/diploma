import axios from 'axios';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function getGenres() {
  return axios
    .get(
      'book/genres/',
      options,
    )
    .then(data => {
      if (!data.error) {
        const genres = data.genres;
        if (genres) {
          return genres;
        }
      }
      return [];
    });
}