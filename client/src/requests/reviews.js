import qs from 'qs';
import axios from 'axios';

const options = {
  transformResponse: [
    data => {
      return JSON.parse(data);
    },
  ],
};

export function getReviews(id) {
  return axios.get(`book/reviews/${id}`, options).then(data => {
    if (!data.error) {
      return data.reviews;
    }
  });
}

export function addReview(review, token) {
  const { book_id, header, review_text } = review;
  return axios
    .post(
      `book/reviews/${book_id}`,
      qs.stringify({
        header: header,
        review_text: review_text,
        token: token,
      }),
      options,
    );
}
