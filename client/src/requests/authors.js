import axios from 'axios';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function searchAuthor(value) {
  return axios
    .get(
      'author/',
      {
        params: {
          search: value,
        },
      },
      options,
    )
    .then(data => {
      if (!data.error) {
        const authors = data.authors;
        if (authors !== null) {
          return authors;
        }
      }
      return [];
    });
}
