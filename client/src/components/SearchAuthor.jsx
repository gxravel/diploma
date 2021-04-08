import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { searchAuthor } from '../requests/authors';
import { useDebounce } from '../auxiliary/hooks';
import Autocomplete from '@material-ui/lab/Autocomplete';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles(theme => ({

  inputRoot: {
    color: 'inherit',
  },
}));

export default function SearchAuthor(props) {
  const classes = useStyles();
  const [searchTerm, setSearchTerm] = useState('');
  const [searchList, setSearchList] = useState([]);
  const [isSearching, setIsSearching] = useState(false);
  const debouncedSearchTerm = useDebounce(searchTerm, 500);
  
  const { author, setAuthorId } = props;

  const [open, setOpen] = React.useState(false);

  const handleChangeTags = (e, v) => {
    setAuthorId(v);
  }

  useEffect(() => {
    if (debouncedSearchTerm) {
      setIsSearching(true);
      searchAuthor(debouncedSearchTerm).then(results => {
        setIsSearching(false);
        setSearchList(results);
        console.log(results);
      });
    } else {
      setSearchList([]);
    }
  }, [debouncedSearchTerm]);

  useEffect(() => {
    if (!open) {
      setSearchList([]);
    }
  }, [open]);

  return (
    <div className={classes.autocomplete}>
      <Autocomplete
        className={classes.autocomplete}
        id="search-author"
        noOptionsText={'sadas'}
        freeSolo
        open={open}
        onOpen={() => {
          setOpen(true);
        }}
        onClose={() => {
          setOpen(false);
        }}
        loading={isSearching}
        onChange={handleChangeTags}
        options={searchList}
        renderInput={params => (
          <TextField
            {...params}
            fullWidth
            variant="outlined"
            placeholder={author}
            onChange={e => setSearchTerm(e.target.value)}
            className={classes.inputInput}
          />
        )}
        getOptionLabel={author => author.name}
        renderOption={author => (
          <div
            className={classes.inputInput}
          >
            <Typography display="block" variant="body1">
              {author.name}
            </Typography>
            {author.name !== author.original_name && (
              <Typography display="block" variant="subtitle2">
                {author.original_name}
              </Typography>
            )}
          </div>
        )}
      />
    </div>
  );
}
