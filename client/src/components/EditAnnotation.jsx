import React, { useState } from 'react';
import Paper from '@material-ui/core/Paper';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/core/styles';
import TabPanel from './TabPanel';

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

const useStyles = makeStyles(theme => ({
  annotation: {
    paddingTop: theme.spacing(3),
    backgroundColor: theme.palette.background.default,
  },
  tabs: {
  },
}));

export default function EditAnnotation(props) {
  const classes = useStyles();
  const [value, setValue] = useState(0);
  const { annotation, handleSetResult } = props;

  function handleChange(event, newValue) {
    setValue(newValue);
  }

  return (
    <Paper className={classes.annotation} square>
      <Tabs
        className={classes.tabs}
        value={value}
        indicatorColor="primary"
        textColor="inherit"
        onChange={handleChange}
        aria-label="annotations"
      >
        <Tab label="Аннотация" {...a11yProps(0)} />
      </Tabs>
      <TabPanel value={value} index={0}>
      <TextField
          id="book-annotation"
          variant="outlined"
          multiline
          fullWidth
          value={annotation}
          onChange={e => handleSetResult('annotation', e.currentTarget.value)}
        />
      </TabPanel>
    </Paper>
  );
}
