import React, { useState, useEffect } from 'react';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core';
import BookCard from '../components/BookCard';
import FoldersList from '../components/FoldersList';
import DisplayOptions from '../components/DisplayOptions';
import Grid from '@material-ui/core/Grid';
import {
  getCollection,
  getCollections,
  deleteFromCollection,
} from '../requests/collections';

const useStyles = makeStyles(theme => ({
  collections: {
    width: '100%',
    display: 'flex',
    color: theme.palette.text.primary,
    backgroundColor: theme.palette.background.default,
  },
  folders: {
    flexGrow: 0,
    width: '25%',
    height: '100%',
  },
  displayOptions: {
    flexGrow: 1,
    paddingLeft: theme.spacing(3),
    paddingRight: theme.spacing(3),
    paddingTop: theme.spacing(2),
    height: '100%',
  },
  grid: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'space-around',
    overflow: 'hidden',
    backgroundColor: theme.palette.background.paper,
  },
  gridList: {
    width: 500,
    height: 450,
  },
  icon: {
    color: 'rgba(255, 255, 255, 0.54)',
  },
  books: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
    margin: 'auto',
  },
}));

export default function MyCollections(props) {
  const classes = useStyles();
  const [folders, setFolders] = useState(['Буду читать']);
  const [page, setPage] = React.useState(0);
  const [rowsPerPage] = React.useState(5);
  const [currentFolder, setCurrentFolder] = useState('Буду читать');
  const [books, setBooks] = useState([]);
  const [sorting, setSorting] = useState('mc.added');

  const { token } = props;

  const handleChangePage = newPage => {
    setPage(newPage);
  };

  // const handleChangeRowsPerPage = event => {
  //   setRowsPerPage(parseInt(event.target.value, 10));
  //   setPage(0);
  // };
  const handleChangeFolder = folder => {
    setCurrentFolder(folder);
  };

  const handleRemove = id => {
    deleteFromCollection(token, id, currentFolder).then(() => {
      getCollection(token, currentFolder, rowsPerPage, page, sorting).then(
        books => {
          setBooks(books);
        },
      );
    });
  };

  const handleChangeSorting = sorting => {
    setSorting(sorting);
  };

  useEffect(() => {
    getCollections(token).then(collections => {
      setFolders(collections);
    });
  }, [token]);

  useEffect(() => {
    console.log('we in effect: ', sorting);
    getCollection(token, currentFolder, rowsPerPage, page, sorting).then(
      books => {
        console.log(books);
        setBooks(books);
      },
    );
  }, [token, currentFolder, rowsPerPage, page, sorting]);

  return (
    <>
      <Paper className={classes.collections}>
        <div className={classes.folders}>
          {folders && (
            <FoldersList
              folders={folders}
              handleChangeFolder={handleChangeFolder}
            />
          )}
        </div>
        <div className={classes.displayOptions}>
          {page !== undefined && rowsPerPage && (
            <DisplayOptions
              handleChangeSorting={handleChangeSorting}
              handleChangePage={handleChangePage}
              page={page}
              rowsPerPage={rowsPerPage}
            />
          )}
        </div>
      </Paper>
      <div className={classes.books}>
        <Paper className={classes.paper}>
          <Grid container spacing={2}>
            {books &&
              books
                .map(book => (
                  <BookCard
                    key={book.id}
                    handleRemove={handleRemove}
                    book={book}
                  />
                ))}
          </Grid>
        </Paper>
      </div>
    </>
  );
}
