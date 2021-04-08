import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { useHistory } from 'react-router-dom';
import ButtonBase from '@material-ui/core/ButtonBase';

const useStyles = makeStyles(() => ({
  image: {
    width: 224,
    height: 224,
    maxWidth: '100%',
  },
  img: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
  },
  annotation: {
    maxHeight: 100,
    overflow: 'hidden',
  },
}));

export default function BookCard(props) {
  const classes = useStyles();
  const { handleRemove, book } = props;
  const {
    id,
    title,
    original_title,
    publication,
    annotation,
    author,
    genres,
  } = book;

  const history = useHistory();
  const fserver = 'http://localhost:8100/data/';

  function handleClick() {
    history.push(`book/${id}`);
  }

  return (
    <>
      <Grid item xs={3} onClick={handleClick}>
        <ButtonBase className={classes.image}>
          <img
            className={classes.img}
            alt="apple"
            src={fserver + id + '.png'}
          />
        </ButtonBase>
      </Grid>
      <Grid
        style={{ marginLeft: 5 }}
        item
        xs={9}
        container
        direction="column"
        spacing={3}
      >
        <Grid item xs>
          {title !== original_title && (
            <Typography variant="caption" color="textSecondary">
              {`${original_title}`}
            </Typography>
          )}
          <Typography gutterBottom variant="body1">
            {`${title} (${publication}), ${genres}`}
          </Typography>
          <Typography dispay="inline" variant="body2" gutterBottom>
            {author}
          </Typography>
          <Typography
            className={classes.annotation}
            variant="body2"
            color="textSecondary"
            onClick={e => {
              e.persist();
              console.log(e.target.style);
              e.target.style['max-height'] = 'none';
            }}
          >
            {annotation}
          </Typography>
        </Grid>
        <Grid
          item
          onClick={() => {
            handleRemove(id);
          }}
        >
          <Typography variant="body2" style={{ cursor: 'pointer' }}>
            Remove
          </Typography>
        </Grid>
      </Grid>
    </>
  );
}
