import { makeStyles } from '@material-ui/core';
import Paper from '@material-ui/core/Paper';
import React, { useEffect, useState } from 'react';
import { useParams, Switch, Route } from 'react-router-dom';
import Annotation from '../components/Annotation';
import BookBelongs from '../components/BookBelongs';
import Download from '../components/Download';
import InfoTable from '../components/InfoTable';
import EditBook from './EditBook';
import { getBookInfo } from '../requests/books';
import {
  addToCollection,
  deleteFromCollection,
} from './../requests/collections';
import Reviews from '../components/Reviews';
import BookRating from '../components/BookRating';
import PrivateRoute from '../components/PrivateRoute';

const useStyles = makeStyles(theme => ({
  bookDetails: {
    width: '100%',
    display: 'flex',
    color: theme.palette.text.primary,
    backgroundColor: theme.palette.background.default,
  },
  imgBox: {
    flexGrow: 0,
  },
  img: {
    marginTop: theme.spacing(2),
    marginRight: theme.spacing(2),
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    display: 'none',
    [theme.breakpoints.up('sm')]: {
      display: 'block',
      height: 384,
    },
    [theme.breakpoints.down('sm')]: {
      display: 'block',
      height: 256,
    },
  },
  infoBox: {
    flexGrow: 1,
    paddingLeft: theme.spacing(3),
    paddingRight: theme.spacing(3),
    paddingTop: theme.spacing(2),
  },
  name: {},
}));

export default function BookDetails(props) {
  const classes = useStyles();
  const [info, setInfo] = useState(null);
  const [belongs, setBelongs] = useState({
    'Буду читать': false,
    Избранное: false,
  });
  const { bookId } = useParams();
  const { token, admin } = props;
  const fserver = 'http://localhost:8100/data/';

  function handleBelongsState(collection, doesBelong) {
    setBelongs({
      ...belongs,
      [collection]: doesBelong,
    });
    if (doesBelong) {
      addToCollection(token, bookId, collection);
    } else {
      deleteFromCollection(token, bookId, collection);
    }
  }

  useEffect(() => {
    getBookInfo(bookId, token)
    .then(data => {
      const { books, belongs } = data;
      if (belongs) {
        setBelongs(belongs);
      }
      if (books) {
        setInfo(books[0]);
      }
    });
  }, [bookId, token]);

  return (
    <Switch>
      <PrivateRoute path="/book/:bookId/edit" checkAdmin={true} admin={admin}>
        {info && (
          <EditBook className={classes.mainInfo} info={info} token={token} />
        )}
      </PrivateRoute>

      <Route path="/book/:bookId">
        <Paper className={classes.bookDetails}>
          <div className={classes.imgBox}>
            {info && (
              <img
                className={classes.img}
                src={fserver + info.id + '.png'}
                alt="Book"
              />
            )}
            <BookBelongs
              belongs={belongs}
              handleBelongsState={handleBelongsState}
            />
            <Download bookId={bookId} />
            <BookRating bookId={bookId} token={token} />
          </div>
          <div className={classes.infoBox}>
            {info && <InfoTable className={classes.mainInfo} info={info} admin={admin} />}
          </div>
        </Paper>
        {info && <Annotation annotation={info.annotation} />}
        <Reviews bookId={bookId} token={token} />
      </Route>
    </Switch>
  );
}
