import axios from 'axios';
import { host } from '../auxiliary/constants';
import qs from 'qs';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function search(value) {
  return axios
    .get(
      'book',
      {
        params: {
          search: value,
        },
      },
      options,
    )
    .then(data => {
      if (!data.error) {
        const books = data.books;
        if (books !== null) {
          return books;
        }
      }
      return [];
    });
}

export function getBookInfo(id, token) {
  return axios
    .get('book/' + id, { params: { token: token } }, options)
    .then(data => {
      if (!data.error) {
        return data;
      }
      return [];
    });
}

export function putBook(book, token, file) {
  var formData = new FormData();
  if (file) {
    formData.append("file", file);
  }
  formData.append("token", token);
  formData.append("book", JSON.stringify(book));
  return axios
    .put('book/' + book.id, formData , options, {  headers: {
      'Content-Type': 'multipart/form-data'
    }})
    .then(data => {
      if (!data.error) {
        return '';
      }
    });
}

export function postBook(book, token, file) {
  var formData = new FormData();
  if (file) {
    formData.append("file", file);
  }
  formData.append("token", token);
  formData.append("book", JSON.stringify(book));
  return axios
    .post('book/', formData , options, {  headers: {
      'Content-Type': 'multipart/form-data'
    }})
    .then(data => {
      if (!data.error) {
        return '';
      }
    });
}
const instance = axios.create();
export function downloadBook(id, ext) {
  instance
    .get(`${host}book/download/${id}`, {
      params: {
        ext: ext,
      },
      responseType: 'blob',
    })
    .then(response => {
      const data = response.data;
      const re = /filename=(.*)/;
      const match = response.request.getResponseHeader('Content-Disposition').match(re);
      const filename = decodeURIComponent(match[1]) + ext;
      if (!data.error) {
        console.log('db response: ', data);
        const url = window.URL.createObjectURL(new Blob([data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', filename);
        document.body.appendChild(link);
        link.click();
      }
    })
}
