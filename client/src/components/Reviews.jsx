import React, { useState, useEffect } from 'react';
import Paper from '@material-ui/core/Paper';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import TabPanel from './TabPanel';
import { makeStyles } from '@material-ui/core/styles';
import Review from './Review';
import AddReview from './AddReview';
import { getReviews } from './../requests/reviews';
import uuidv1 from 'uuid/v1';
import { Divider } from '@material-ui/core';

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

const useStyles = makeStyles(theme => ({
  reviews: {
    paddingTop: theme.spacing(3),
    backgroundColor: theme.palette.background.default,
  },
  divider: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
}));

export default function Reviews(props) {
  const classes = useStyles();
  const [value, setValue] = useState(0);
  const [reviews, setReviews] = useState([]);
  const { bookId, token } = props;

  function handleChange(event, newValue) {
    setValue(newValue);
  }

  function updateReviews() {
    return getReviews(bookId).then(data => {
      console.log('data: ', data);
      setReviews(data);
    });
  }

  function backToReviews() {
    setValue(0);
  }

  useEffect(() => {
    getReviews(bookId).then(data => {
      console.log('data: ', data);
      setReviews(data);
    });
  }, [bookId]);

  return (
    <Paper className={classes.reviews} square>
      <Tabs
        className={classes.tabs}
        value={value}
        indicatorColor="primary"
        textColor="inherit"
        onChange={handleChange}
        aria-label="reviews"
      >
        <Tab label="Рецензии читателей" {...a11yProps(0)} />
        <Tab label="Добавить рецензию" {...a11yProps(1)} />
      </Tabs>
      <TabPanel value={value} index={0}>
        {reviews &&
          reviews.map(r => (
            <div key={uuidv1()}>
              <Review review={r} />
              <Divider className={classes.divider} />
            </div>
          ))}
      </TabPanel>
      <TabPanel value={value} index={1}>
        <AddReview
          bookId={bookId}
          token={token}
          update={updateReviews}
          backToReviews={backToReviews}
        />
      </TabPanel>
    </Paper>
  );
}
