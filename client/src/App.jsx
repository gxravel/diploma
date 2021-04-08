import React from 'react';
import {
  createMuiTheme,
  makeStyles,
  ThemeProvider,
} from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';

import BookDetails from './layouts/BookDetails';
import Register from './layouts/Register';
import TopBar from './components/TopBar';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from 'react-router-dom';
import { useStateWithSessionStorage } from './auxiliary/hooks';
import { tokenKey } from './auxiliary/constants';
import MyCollections from './layouts/MyCollections';
import { logout } from './requests/account';
import Authenticate from './layouts/Authenticate';
import UserManager from './layouts/UserManager';
import PrivateRoute from './components/PrivateRoute';
import axios from 'axios';
import { host } from './auxiliary/constants';
import EditBook from './layouts/EditBook';
import { useEffect, useState } from 'react';

axios.defaults.baseURL = host;

axios.defaults.headers.post['Content-Type'] =
  'application/x-www-form-urlencoded';

axios.interceptors.response.use(
  response => {
    let data = response.data;
    if (data.data) {
      console.log('data: ', data.data);
      return data.data;
    }
    if (data.response) {
      console.log('response: ', data.response);
      return data.response;
    }
    console.log('error: ', data.error);
    return data;
  },
  error => {
    console.log('critical error: ', error);
  },
);

const theme = createMuiTheme({
  palette: {
    type: 'dark',
    primary: {
      main: '#212121',
    },
    secondary: {
      main: '#616161',
    },
  },
});

const useStyles = makeStyles(() => ({
  app: {
    width: '100%',
    height: '100%',
    backgroundColor: '#424242',
    boxSizing: 'inherit',
  },
  crop: {
    width: '75%',
    margin: 'auto',
    boxSizing: 'inherit',
  },
}));

export default function App() {
  const classes = useStyles();
  const [token, setToken] = useStateWithSessionStorage(tokenKey);
  const [admin, setAdmin] = useState(sessionStorage.getItem('admin') == 'true');
  useEffect(() => {
    sessionStorage.setItem('admin', admin);
  }, [admin]);

  const handleChangeToken = (token, admin) => {
    setToken(token);
    setAdmin(admin);
  };

  return (
    <ThemeProvider theme={theme}>
      <Router>
        <div className={classes.app}>
          <Box className={classes.crop}>
            <TopBar token={token} admin={admin} />
            <Switch>
              <Route path="/book/:bookId">
                <BookDetails token={token} admin={admin} />
              </Route>
              <PrivateRoute path="/book/" checkAdmin={true} admin={admin}>
                <EditBook token={token} />
              </PrivateRoute>
              <PrivateRoute path="/collections" token={token}>
                <MyCollections token={token} />
              </PrivateRoute>
              <Route path="/account/register">
                <Register handleChangeToken={handleChangeToken} />
              </Route>
              <Route path="/account/auth">
                <Authenticate handleChangeToken={handleChangeToken} />
              </Route>
              <Route
                path="/account/logout"
                render={({ location }) => (
                  <>
                    {logout(token, handleChangeToken)}
                    <Redirect
                      to={{
                        pathname: '/',
                        state: { from: location },
                      }}
                    />
                  </>
                )}
              ></Route>
              <PrivateRoute
                path="/manage/users/"
                checkAdmin={true}
                admin={admin}
              >
                <UserManager token={token} />
              </PrivateRoute>
            </Switch>
          </Box>
        </div>
      </Router>
    </ThemeProvider>
  );
}
