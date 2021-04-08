import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import MenuItem from '@material-ui/core/MenuItem';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import LibraryAddIcon from '@material-ui/icons/LibraryAdd';
import Menu from '@material-ui/core/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import MoreIcon from '@material-ui/icons/MoreVert';
import { useHistory } from 'react-router-dom';
import { Link } from '@material-ui/core';
import Search from './Search';

const useStyles = makeStyles(theme => ({
  grow: {
    flexGrow: 1,
  },
  main: {
    height: '10vh',
  },
  title: {
    display: 'none',
    [theme.breakpoints.up('sm')]: {
      display: 'block',
    },
  },
  sectionDesktop: {
    display: 'none',
    [theme.breakpoints.up('md')]: {
      display: 'flex',
    },
  },
  sectionMobile: {
    display: 'flex',
    [theme.breakpoints.up('md')]: {
      display: 'none',
    },
  },
}));

export default function TopBar(props) {
  const classes = useStyles();

  const history = useHistory();
  const { token, admin } = props;

  const [anchorEl, setAnchorEl] = React.useState(null);
  const [mobileMoreAnchorEl, setMobileMoreAnchorEl] = React.useState(null);

  const isMenuOpen = Boolean(anchorEl);
  const isMobileMenuOpen = Boolean(mobileMoreAnchorEl);

  function handleProfileMenuOpen(event) {
    setAnchorEl(event.currentTarget);
  }

  function handleMobileMenuClose() {
    setMobileMoreAnchorEl(null);
  }

  function handleMenuClose() {
    setAnchorEl(null);
    handleMobileMenuClose();
  }

  function handleMobileMenuOpen(event) {
    setMobileMoreAnchorEl(event.currentTarget);
  }

  const menuId = 'primary-search-account-menu';
  const renderMenu = (
    <Menu
      anchorEl={anchorEl}
      anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
      id={menuId}
      keepMounted
      transformOrigin={{ vertical: 'top', horizontal: 'right' }}
      open={isMenuOpen}
      onClose={handleMenuClose}
    >
      {!token
        ? [
            <MenuItem
              key={1}
              onClick={() => {
                history.push('/account/register');
                handleMenuClose();
              }}
            >
              Регистрация
            </MenuItem>,
            <MenuItem
              key={2}
              onClick={() => {
                history.push('/account/auth');
                handleMenuClose();
              }}
            >
              Вход
            </MenuItem>,
          ]
        : admin
        ? [
            <MenuItem
              key={1}
              onClick={() => {
                history.push('/manage/users');
                handleMenuClose();
              }}
            >
              Управление пользователями
            </MenuItem>,
            <MenuItem
              key={2}
              onClick={() => {
                history.push('/account/logout');
                handleMenuClose();
              }}
            >
              Выход
            </MenuItem>,
          ]
        : [
            <MenuItem
              key={2}
              onClick={() => {
                history.push('/account/logout');
                handleMenuClose();
              }}
            >
              Выход
            </MenuItem>,
          ]}
    </Menu>
  );

  const mobileMenuId = 'primary-search-account-menu-mobile';
  const renderMobileMenu = (
    <Menu
      anchorEl={mobileMoreAnchorEl}
      anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
      id={mobileMenuId}
      keepMounted
      transformOrigin={{ vertical: 'top', horizontal: 'right' }}
      open={isMobileMenuOpen}
      onClose={handleMobileMenuClose}
    >
      <MenuItem onClick={handleProfileMenuOpen}>
        <IconButton
          aria-label="account of current user"
          aria-controls="primary-search-account-menu"
          aria-haspopup="true"
          color="inherit"
        >
          <AccountCircle />
        </IconButton>
        <p>Мой аккаунт</p>
      </MenuItem>
    </Menu>
  );

  return (
    <div className={classes.main}>
      <AppBar position="static">
        <Toolbar>
          <Typography className={classes.title} variant="h6" noWrap>
            <Link href="/" color="inherit">
              Book-machine
            </Link>
          </Typography>
          <div className={classes.grow} />
          <Search />
          {admin && (
            <IconButton href="/book/" color="inherit">
              <LibraryAddIcon />
            </IconButton>
          )}
          {token && (
            <IconButton href="/collections" color="inherit">
              <LibraryBooksIcon />
            </IconButton>
          )}
          <div className={classes.sectionDesktop}>
            <IconButton
              edge="end"
              aria-label="account of current user"
              aria-controls={menuId}
              aria-haspopup="true"
              onClick={handleProfileMenuOpen}
              color="inherit"
            >
              <AccountCircle />
            </IconButton>
          </div>
          <div className={classes.sectionMobile}>
            <IconButton
              aria-label="show more"
              aria-controls={mobileMenuId}
              aria-haspopup="true"
              onClick={handleMobileMenuOpen}
              color="inherit"
            >
              <MoreIcon />
            </IconButton>
          </div>
        </Toolbar>
      </AppBar>
      {renderMobileMenu}
      {renderMenu}
    </div>
  );
}
