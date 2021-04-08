import React, { useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { auth, register } from '../requests/account';
import { useHistory } from 'react-router-dom';
import { colors } from '../auxiliary/constants';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  main: {
    display: 'flex',
    width: '100%',
    height: '90vh',
    color: theme.palette.text.primary,
    backgroundColor: theme.palette.background.default,
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  green: {
    color: colors.green,
  },
}));

export default function Register(props) {
  const classes = useStyles();
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const { handleChangeToken } = props;
  const [error, setError] = useState('');
  const history = useHistory();

  function handleClick() {
    register(login, password)
      .then(data => {
        console.log(data);
        if (data.error) {
          setError(data.error.text);
          return true;
        }
      })
      .then(err => {
        if (err) {
          return true;
        }
        auth(login, password)
          .then(data => {
            console.log(data);
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
      });
  }

  return (
    <div className={classes.main}>
      <Container maxWidth="xs">
        <div className={classes.paper}>
          <Avatar className={classes.avatar}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Регистрация
          </Typography>
          <form className={classes.form} noValidate>
            <Grid container spacing={2}>
              {error && (
                <Grid item xs={12}>
                  <Typography variant="body2" color="error">
                    {error}
                  </Typography>
                </Grid>
              )}
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  id="login"
                  label="Логин"
                  name="login"
                  autoComplete="username"
                  value={login}
                  onChange={e => setLogin(e.currentTarget.value)}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
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
              </Grid>
            </Grid>
            <Button
              fullWidth
              variant="contained"
              color="primary"
              className={classes.submit}
              onClick={handleClick}
            >
              Зарегистрироваться
            </Button>
            <Grid container justify="flex-end">
              <Grid item>
                <Link href="/account/auth" variant="body2" className={classes.green}>
                  У вас уже есть аккаунт? Авторизуйтесь!
                </Link>
              </Grid>
            </Grid>
          </form>
        </div>
      </Container>
    </div>
  );
}
