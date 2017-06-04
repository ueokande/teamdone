import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import { connect } from 'react-redux';
import { Provider } from "react-redux"

import createStoreWithMiddleware from './store/configure';
import TaskList from './component/task-list';

class Org extends Component {
  constructor() {
    super()
  }

  render() {
    return (
      <TaskList entries={[
          {
            id: 0,
            name: 'Hello',
            due: '2017-10-11',
          }, {
            id: 1,
            name: 'World',
            due: '2017-10-15',
          }
        ]}
      />
    )
  }
}

const store = createStoreWithMiddleware();

window.addEventListener('load', () => {
  var parent = document.getElementById('teamdone-org');
  if (parent !== null) {
    ReactDOM.render(
      <Provider store={store}>
        <MuiThemeProvider>
          <Org />
        </MuiThemeProvider>
      </Provider>,
      parent
    );
  }
});
