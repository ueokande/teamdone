import React, { Component } from 'react';
import { List, ListItem } from 'material-ui/List';
import PropTypes from 'prop-types';

export default class TaskList extends Component {
  render() {
    return (
      <List>
      {
        this.props.entries.map((e) => {
          return (
            <ListItem
              key={e.id}
              primaryText={e.name}
              secondaryText={e.due}
            />
          );
        })
      }
      </List>
    );
  }
}

TaskList.defaultProps = {
  entries: []
};

TaskList.propTypes = {
  entries: PropTypes.arrayOf(PropTypes.object)
};
