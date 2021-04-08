import React, { useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import { auth } from '../requests/account';
import Container from '@material-ui/core/Container';
import { colors } from '../auxiliary/constants';
import { useHistory } from 'react-router-dom';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  green: {
    color: colors.green,
  },
  main: {
    display: 'flex',
    width: '100%',
    height: '90vh',
    color: theme.palette.text.primary,
    backgroundColor: theme.palette.background.default,
  },
}));

export default function Authenticate(props) {
  const classes = useStyles();
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const { handleChangeToken } = props;
  const [error, setError] = useState('');
  const history = useHistory();

  function handleClick() {
    auth(login, password)
      .then(data => {
        if (data.error) {
          setError(data.error.text);
          return true;
        }
        handleChangeToken(data.token, data.admin);
      })
      .then(err => {
        if (!err) {
          history.push('/');
        }
      });
  }

  return (
    <div className={classes.main}>

    <Container maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Авторизация
        </Typography>
        <form className={classes.form} noValidate>
          {error && (
            <Typography variant="body2" color="error">
              {error}
            </Typography>
          )}
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="login"
            label="Логин"
            name="login"
            autoComplete="username"
            autoFocus
            value={login}
            onChange={e => setLogin(e.currentTarget.value)}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Пароль"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={e => setPassword(e.currentTarget.value)}
          />
          <Button
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleClick}
          >
            Авторизоваться
          </Button>
          <Grid container justify="flex-end">
            <Grid item>
              <Link href="/account/register" variant="body2" className={classes.green}>
                {"Нет аккаунта? Зарегистрируйтесь!"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
    </Container>
    </div>
  );
}
