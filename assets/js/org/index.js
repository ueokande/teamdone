import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import { connect } from 'react-redux';
import { Provider } from "react-redux"

import createStoreWithMiddleware from './store/configure';
import TaskList from './component/task-list';
import UserFormDialog from './component/user-form-dialog';
import * as user from './action/user';

class Org extends Component {
  constructor() {
    super();
  }

  componentDidMount() {
    this.props.dispatch(user.initialize());
  }

  submitUserCreate(userName) {
    this.props.dispatch(user.submit(userName))
  }

  render() {
    return (
      <div>
        <UserFormDialog
          open={this.props.userFormOpen}
          userNameError={this.props.userNameError}
          submit={(name) => this.submitUserCreate(name)}
        />
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
      </div>
    )
  }
}

const store = createStoreWithMiddleware();

Org = connect(({user}) => {
  return {
    userFormOpen: user.formOpen,
    userNameError: user.formNameError
  }
})(Org)


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
