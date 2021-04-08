import React, { useState } from 'react';
import Paper from '@material-ui/core/Paper';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { Typography } from '@material-ui/core';
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

export default function Annotation(props) {
  const classes = useStyles();
  const [value, setValue] = useState(0);

  const { annotation } = props;
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
        <Typography variant="body1">
          {annotation}
        </Typography>
      </TabPanel>
    </Paper>
  );
}
