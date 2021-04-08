import React from 'react';
import { Route, Redirect } from 'react-router-dom';
// A wrapper for <Route> that redirects to the login
// screen if you're not yet authenticated.
export default function PrivateRoute({ children, ...rest }) {
  return (
    <Route
      {...rest}
      render={({ location }) =>
      (rest.checkAdmin && rest.admin) || (!rest.checkAdmin && rest.token) ? (
          children
        ) : (
          <Redirect
            to={{
              pathname: '/account/auth',
              state: { from: location },
            }}
          />
        )
      }
    />
  );
}
