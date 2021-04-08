import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Typography } from '@material-ui/core';
import { useEffect } from 'react';

const useStyles = makeStyles(theme => ({
  review: {
    padding: theme.spacing(1),
    backgroundColor: theme.palette.background.paper,
  },
  first: {
    display: 'flex',
    justifyContent: 'space-between',
  },
  header: {},
  text: {
    marginTop: theme.spacing(2),
  },
}));

export default function Review(props) {
  const classes = useStyles();
  const { review } = props;
  const { header, review_text, username, date_added } = review;
  const [date, setDate] = useState('');

  useEffect(() => {
    let temp = new Date(Date.parse(date_added));
    setDate(temp.toLocaleString());
  }, [date_added])

  return (
    <div>
      <div className={classes.first}>
        <Typography variant="body1" className={classes.user}>
          {username}
        </Typography>
        <Typography variant="overline" className={classes.date}>
          {date}
        </Typography>
      </div>
      <div className={classes.review}>
        <Typography variant="body2" className={classes.header}>
          {header}
        </Typography>
        <Typography variant="body2" className={classes.text}>
          {review_text}
        </Typography>
      </div>
    </div>
  );
}
