import { makeStyles, Input } from '@material-ui/core';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import EditAnnotation from '../components/EditAnnotation';
import EditInfoTable from '../components/EditInfoTable';
import { putBook, postBook } from '../requests/books';

const useStyles = makeStyles(theme => ({
  bookDetails: {
    width: '100%',
    display: 'flex',
    color: theme.palette.text.primary,
    backgroundColor: theme.palette.background.default,
  },
  imgBox: {
    minWidth: '25%',
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
  // save: {
  //   position: 'sticky',
  //   bottom: 0,
  //   right: 0,
  // },
  buttons: {
    display: 'flex',
    flexDirection: 'column',
    position: 'sticky',
    bottom: 0,
    right: 0,
  },
}));

export default function EditBook(props) {
  const classes = useStyles();
  const { info, token } = props;
  const [result, setResult] = useState(info);

  const history = useHistory();

  const handleSetResult = (field, value) => {
    setResult(prev => {
      return {
        ...prev,
        [field]: value,
      };
    });
  };

  const handleSave = e => {
    const file = document.getElementById('upload-file').files[0];
    let method;
    if (info) {
      method = putBook;
    } else {
      method = postBook;
    }
    method(result, token, file).then(data => {
      let id;
      if (data) {
        id = data;
      } else {
        id = result.id;
      }
      setInterval(() => {
        history.push('/book/' + id);
      }, 1000);
    });
  };

  return (
    <>
      <Paper className={classes.bookDetails}>
        <div className={classes.imgBox}>
          <form>
            <Typography>Загрузить файл</Typography>
            <Input id="upload-file" name="upload-file" type="file" />
          </form>
        </div>
        <div className={classes.infoBox}>
          {result && (
            <EditInfoTable
              className={classes.mainInfo}
              info={result}
              handleSetResult={handleSetResult}
            />
          )}
        </div>
      </Paper>
      {result && (
        <EditAnnotation
          annotation={result.annotation}
          handleSetResult={handleSetResult}
        />
      )}
      <div className={classes.buttons}>
        <Button
          className={classes.save}
          variant="contained"
          color="primary"
          onClick={handleSave}
        >
          Сохранить изменения
        </Button>
        <Button
          className={classes.save}
          variant="contained"
          color="primary"
          onClick={e => {
            history.push('/book/' + result.id);
          }}
        >
          Отмена
        </Button>
      </div>
    </>
  );
}
