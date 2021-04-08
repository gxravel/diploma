import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import uuidv1 from 'uuid/v1';

const useStyles = makeStyles(theme => ({
  folders: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
  },
}));

export default function FoldersList(props) {
  const classes = useStyles();
  const [selectedIndex, setSelectedIndex] = React.useState(0);
  const { folders, handleChangeFolder } = props;

  const handleListItemClick = (event, index) => {
    setSelectedIndex(index);
    handleChangeFolder(event.target.innerText);
  };

  return (
    <div className={classes.folders}>
      <List component="nav" aria-label="secondary mailbox folder">
        {folders.map((folder, i) => (
          <ListItem
            key={uuidv1()}
            button
            selected={selectedIndex === i}
            onClick={event => handleListItemClick(event, i)}
          >
            <ListItemText primary={folder} />
          </ListItem>
        ))}
      </List>
    </div>
  );
}
