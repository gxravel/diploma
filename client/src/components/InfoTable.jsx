/* eslint-disable no-console */
/* eslint-disable react/prop-types */
/* eslint-disable object-curly-newline */
import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import uuidv1 from 'uuid/v1';
import { useState, useEffect } from 'react';
import EditIcon from '@material-ui/icons/Edit';
import { IconButton } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing(3),
    overflowX: 'auto',
  },
  header: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
}));

function createData(column, value) {
  return {
    column,
    value,
  };
}

export default function InfoTable(props) {
  const classes = useStyles();

  const { info, admin } = props;
  const {
    id,
    title,
    original_title,
    original_language,
    publication,
    writing_end,
    writing_start,
    genres,
    author,
  } = info;
  const { name, original_name } = author;
  const [rows, setRows] = useState([]);

  function extractYear(date) {
    date = date.substring(0, 4);
    if (date === '0001') {
      return '';
    }
    return date;
  }

  useEffect(() => {
    const rows = [];
    let rGenres = [];
    if (genres) {
      genres.forEach(function(genre) {
        rGenres.push(genre.name);
      });
    }
    rows.push(createData('Жанр:', rGenres.join(', ')));
    const rPublication = extractYear(publication);
    const rStart = extractYear(writing_start);
    const rEnd = extractYear(writing_end);
    let rWriting = '';
    if (rStart) {
      rWriting = rStart;
      if (rEnd && rWriting !== rEnd) {
        rWriting += '-' + rEnd;
      }
    } else {
      rWriting = rEnd;
    }
    if (rWriting !== '') {
      rows.push(createData('Дата написания:', rWriting));
    }
    if (rPublication !== '') {
      rows.push(createData('Дата первой публикации:', rPublication));
    }
    rows.push(createData('Язык оригинала:', original_language));
    setRows(rows);
  }, [genres, original_language, publication, writing_end, writing_start]);

  return (
    <div>
      <div className={classes.header}>
        <div className={classes.title}>
          <Typography variant="h4">{title}</Typography>
          {title !== original_title && (
            <Typography color="textSecondary" variant="body2">
              {original_title}
            </Typography>
          )}
        </div>
        { admin && (
          <IconButton href={"/book/"+id+"/edit"}>
            <EditIcon />
          </IconButton>

        )}
      </div>

      <Paper className={classes.root}>
        <Table className={classes.table}>
          <TableBody>
            <TableRow key={uuidv1()}>
              <TableCell component="th" size="small">
                Автор:
              </TableCell>
              <TableCell size="medium" align="left">
                <div>{name}</div>
                {name !== original_name && (
                  <Typography color="textSecondary" variant="caption">
                    {original_name}
                  </Typography>
                )}
              </TableCell>
            </TableRow>
            {rows &&
              rows.map(row => (
                <TableRow key={uuidv1()}>
                  <TableCell
                    style={{ width: '30%' }}
                    component="th"
                    scope="row"
                    size="small"
                  >
                    {row.column}
                  </TableCell>
                  <TableCell align="left" size="medium">
                    {row.value}
                  </TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </Paper>
    </div>
  );
}
