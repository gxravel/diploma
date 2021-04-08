import React, { useEffect, useState } from 'react';
import { fade, makeStyles } from '@material-ui/core/styles';
import SearchIcon from '@material-ui/icons/Search';
import { search } from '../requests/books';
import { useDebounce } from '../auxiliary/hooks';
import Autocomplete from '@material-ui/lab/Autocomplete';
import TextField from '@material-ui/core/TextField';
import { useHistory } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles(theme => ({
  search: {
    position: 'relative',
    right: 0,
    borderRadius: theme.shape.borderRadius,
    backgroundColor: fade(theme.palette.common.white, 0.15),
    '&:hover': {
      backgroundColor: fade(theme.palette.common.white, 0.25),
    },
    marginRight: theme.spacing(2),
    marginLeft: 'auto',
    width: '200px',
    [theme.breakpoints.up('sm')]: {
      marginLeft: theme.spacing(3),
      width: '300px',
    },
    [theme.breakpoints.up('md')]: {
      marginLeft: theme.spacing(3),
      width: '500px',
    },
  },
  searchIcon: {
    width: theme.spacing(7),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  inputRoot: {
    color: 'inherit',
  },
  inputInput: {
    padding: theme.spacing(1, 1, 1, 7),
    width: '120px',
    [theme.breakpoints.up('sm')]: {
      width: '220px',
    },
    [theme.breakpoints.up('md')]: {
      width: '420px',
    },
  },
  autocomplete: {
    width: '200px',
  },
}));

export default function Search() {
  const classes = useStyles();
  const [searchTerm, setSearchTerm] = useState('');
  const [searchList, setSearchList] = useState([]);
  const [isSearching, setIsSearching] = useState(false);
  const debouncedSearchTerm = useDebounce(searchTerm, 500);

  const [open, setOpen] = React.useState(false);
  const history = useHistory();
  
  const goToBook = (id) => {
    history.push(`/book/${id}`);
  }

  const handleChangeTags = (e, v) => {
    if (v) {
      goToBook(v.id);
    }
  }

  useEffect(() => {
    if (debouncedSearchTerm) {
      setIsSearching(true);
      search(debouncedSearchTerm).then(results => {
        setIsSearching(false);
        setSearchList(results);
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
    <div className={classes.search}>
      <div className={classes.searchIcon}>
        <SearchIcon />
      </div>
      <div className={classes.autocomplete}>
        <Autocomplete
          className={classes.autocomplete}
          id={'search-field'}
          noOptionsText={''}
          freeSolo
          open={open}
          onOpen={() => {
            setOpen(true);
          }}
          onClose={() => {
            setOpen(false);
          }}
          onChange={handleChangeTags}
          clearOnEscape
          loading={isSearching}
          options={searchList}
          renderInput={params => (
            <TextField
              {...params}
              fullWidth
              onChange={e => setSearchTerm(e.target.value)}
              className={classes.inputInput}
            />
          )}
          getOptionLabel={book => book.title}
          renderOption={book => (
              <div className={classes.inputInput}
              >
                <Typography display='block' variant="body1">{book.title} ({book.publication})</Typography>
                {book.title !== book.original_title && (
                  <Typography display='block' variant="subtitle2">{book.original_title}</Typography>
                )}
                <Typography display='block' variant="body2">{book.author}</Typography>
                
              </div>
          )}
        />
      </div>
    </div>
  );
}
