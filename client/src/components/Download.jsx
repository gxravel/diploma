import React from 'react';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import { downloadBook } from '../requests/books';

const useStyles = makeStyles(theme => ({
  container: {
    margin: 0,
    marginTop: theme.spacing(2),
  },
}));

export default function Download(props) {
  const classes = useStyles();
  const { bookId } = props;

  return (
    <Grid container className={classes.container} justify="center">
      <Grid item>
        <ButtonGroup
          variant="text"
          size="large"
          aria-label="full width outlined button group"
        >
          <Button
            onClick={() => {
              downloadBook(bookId, '.fb2');
            }}
          >
            .fb2
          </Button>
          <Button>.pdf</Button>
          <Button>.txt</Button>
        </ButtonGroup>
      </Grid>
    </Grid>
  );
}
