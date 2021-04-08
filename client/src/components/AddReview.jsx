import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import { Button, Typography } from '@material-ui/core';
import { addReview } from '../requests/reviews'

const useStyles = makeStyles(theme => ({
  container: {
    width: '100%',
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: '100%',
  },
  publish: {
    display: 'flex',
    justifyContent: 'flex-end',
  },
  reviewText: {
    marginTop: theme.spacing(2),
  },
}));

export default function AddReview(props) {
  const classes = useStyles();
  const [header, setHeader] = useState('');
  const [text, setText] = useState('');
  const {bookId, token, update, backToReviews} = props;

  function handleClick() {
    let review = {
      book_id: bookId,
      header: header,
      review_text: text,
    }
    addReview(review, token).then(() => {
      setHeader('');
      setText('');
      update().then(() => {
        backToReviews();
      });
    });
  }

  return (
    <form className={classes.container} noValidate autoComplete="off">
      <Typography>Заголовок рецензии</Typography>
      <TextField
          id="add-review-header"
          className={classes.textField}
          value={header}
          margin="normal"
          variant="outlined"
          onChange={(e) => setHeader(e.currentTarget.value)}
      />
      <Typography className={classes.reviewText}>Текст рецензии</Typography>
      <TextField
          id="add-review-text"
          multiline
          value={text}
          rows="6"
          className={classes.textField}
          placeholder="Расскажите, чем вам понравилась книга"
          margin="normal"
          variant="outlined"
          onChange={(e) => setText(e.currentTarget.value)}
      />
      <div className={classes.publish}>
        <Button onClick={handleClick}>Опубликовать</Button>
      </div>
    </form>
  );
}