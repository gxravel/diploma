import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
import SearchAuthor from './SearchAuthor';
import GetGenres from './GetGenres';

const useStyles = makeStyles(theme => ({
  // root: {
  //   width: '100%',
  //   marginTop: theme.spacing(3),
  //   overflowX: 'auto',
  //   display: 'flex',
  //   flexDirection: 'row',
  // },
  imgSector: {
    width: '30%',
  },
  paper: {
    width: '100%',
    marginTop: theme.spacing(3),
    overflowX: 'auto',
  },
  header: {
    width: 'auto',
    display: 'flex',
    flexDirection: 'column',
  },
  col1: {
    width: '15%',
  },
  writingDate: {
    // display: 'flex',
    // flexDirection: 'row',
  },
}));

export default function InfoTable(props) {
  const classes = useStyles();
  const { info, handleSetResult } = props;
  const {
    title,
    original_title,
    original_language,
    publication,
    writing_end,
    writing_start,
    genres,
    author,
  } = info;
  const { name } = author;

  const setAuthorId = value => {
    handleSetResult('author', value);
    console.log(value);
  };

  return (
    <div className={classes.root}>
      <form className={classes.header} noValidate autoComplete="off">
        <TextField
          id="book-title"
          variant="outlined"
          value={title}
          onChange={e => handleSetResult('title', e.currentTarget.value)}
        />
        <TextField
          id="book-original_title"
          variant="outlined"
          value={original_title}
          onChange={e =>
            handleSetResult('original_title', e.currentTarget.value)
          }
        />
      </form>
      <Paper className={classes.paper}>
        <Table className={classes.table}>
          <TableBody>
            <TableRow>
              <TableCell className={classes.col1} component="th" size="small">
                Автор:
              </TableCell>
              <TableCell className={classes.col2} align="left">
                <div>
                  <SearchAuthor author={name} setAuthorId={setAuthorId} />
                </div>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell className={classes.col1} component="th" size="small">
                Жанры:
              </TableCell>
              <TableCell className={classes.col2} align="left">
                <div>
                  <GetGenres
                    genres={genres}
                    handleSetResult={handleSetResult}
                  />
                </div>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell className={classes.col1} component="th" size="small">
                Дата написания:
              </TableCell>
              <TableCell className={classes.col2} align="left">
                <div className={classes.writingDate}>
                  <TextField
                    id="book-writing_start"
                    variant="outlined"
                    value={writing_start}
                    onChange={e =>
                      handleSetResult('writing_start', e.currentTarget.value)
                    }
                  />
                  <TextField
                    id="book-writing_end"
                    variant="outlined"
                    value={writing_end}
                    onChange={e =>
                      handleSetResult('writing_end', e.currentTarget.value)
                    }
                  />
                </div>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell className={classes.col1} component="th" size="small">
                Дата первой публикации:
              </TableCell>
              <TableCell className={classes.col2} align="left">
                <TextField
                  id="book-publication"
                  variant="outlined"
                  value={publication}
                  onChange={e =>
                    handleSetResult('publication', e.currentTarget.value)
                  }
                />
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell className={classes.col1} component="th" size="small">
                Язык оригинала:
              </TableCell>
              <TableCell className={classes.col2} align="left">
                <TextField
                  id="book-original_language"
                  variant="outlined"
                  value={original_language}
                  onChange={e =>
                    handleSetResult('original_language', e.currentTarget.value)
                  }
                />
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </Paper>
    </div>
  );
}
