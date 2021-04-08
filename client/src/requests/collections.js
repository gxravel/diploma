import qs from 'qs';
import axios from 'axios';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function getCollections(token) {
  return axios
    .get(
      'my-collections',
      {
        params: {
          token: token,
        },
      },
      options,
    )
    .then(data => {
      if (!data.error) {
        return data.collections;
      }
    });
}

export function getCollection(token, collection, onPage, page, sorting) {
  return axios
    .get(
      'my-collections/' + encodeURIComponent(collection),
      {
        params: {
          token: token,
          onpage: onPage,
          page: page + 1,
          column: sorting,
        },
      },
      options,
    )
    .then(data => {
      if (!data.error) {
        return data.books;
      }
    });
}

export function addToCollection(token, id, collection) {
  return axios.post(
    'my-collections/' + id,
    qs.stringify({
      collection: collection,
      token: token,
    }),
    options,
  );
}

export function deleteFromCollection(token, id, collection) {
  return axios.delete(
    'my-collections/' + id,
    {
      params: {
        collection: collection,
        token: token,
      },
    },
    options,
  );
}
