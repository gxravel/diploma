import React, { useEffect } from 'react';
import ToggleButton from '@material-ui/lab/ToggleButton';
import { makeStyles } from '@material-ui/core/styles';
import { Typography } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  toggle: {
    marginTop: theme.spacing(2),
    width: '90%',
    display: 'block',
    marginLeft: 'auto',
    textTransform: 'none',
  },
}));

export default function BookBelongs(props) {
  const classes = useStyles();
  const [selected, setSelected] = React.useState(false);
  const { belongs, handleBelongsState } = props;

  useEffect(() => {
    setSelected(belongs["Буду читать"]);
  }, [belongs])

  return (
    <ToggleButton
      value="Буду читать"
      selected={selected}
      onChange={(e) => {
        handleBelongsState(e.currentTarget.value, !selected);
      }}
      className={classes.toggle}
    >
      <Typography className={classes.title} component="div">
        Буду читать
      </Typography>
    </ToggleButton>
  );
}
