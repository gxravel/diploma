import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Autocomplete from '@material-ui/lab/Autocomplete';
import TextField from '@material-ui/core/TextField';
import CheckBoxOutlineBlankIcon from '@material-ui/icons/CheckBoxOutlineBlank';
import CheckBoxIcon from '@material-ui/icons/CheckBox';
import Checkbox from '@material-ui/core/Checkbox';
import { getGenres } from '../requests/genres';

const icon = <CheckBoxOutlineBlankIcon fontSize="small" />;
const checkedIcon = <CheckBoxIcon fontSize="small" />;

const useStyles = makeStyles(theme => ({
  inputRoot: {
    color: 'inherit',
  },
}));

export default function GetGenres(props) {
  const classes = useStyles();
  const [list, setList] = useState([]);
  const { genres, handleSetResult } = props;
  const [stringGenres, setStringGenres] = useState('');
  const [open, setOpen] = useState(false);
  const [checked, setChecked] = useState([]);

  const handleChangeTags = (e, v) => {
    setChecked(v);
  };

  useEffect(() => {
    let formattedGenres = [];
    if (genres) {
      genres.forEach(genre => {
        formattedGenres.push(genre.name);
      });
    }
    setStringGenres(formattedGenres.join(', '));
    console.log(genres);
    getGenres().then(genres => {
      setList(genres);
    });
  }, []);

  useEffect(() => {
    if (!open && checked.length > 0) {
      handleSetResult('genres', checked);
    }
  }, [open, checked]);

  return (
    <div className={classes.autocomplete}>
      <Autocomplete
        multiple
        className={classes.autocomplete}
        id="search-genre"
        freeSolo
        disableCloseOnSelect
        open={open}
        onOpen={() => {
          setOpen(true);
        }}
        onClose={() => {
          setOpen(false);
        }}
        onChange={handleChangeTags}
        options={list}
        renderInput={params => (
          <TextField
            {...params}
            fullWidth
            variant="outlined"
            placeholder={stringGenres}
            className={classes.inputInput}
          />
        )}
        getOptionLabel={genre => genre.name}
        renderOption={(genre, { selected }) => (
          <React.Fragment>
            <Checkbox
              icon={icon}
              checkedIcon={checkedIcon}
              style={{ marginRight: 8 }}
              checked={selected}
            />
            {genre.name}
          </React.Fragment>
        )}
      />
    </div>
  );
}
