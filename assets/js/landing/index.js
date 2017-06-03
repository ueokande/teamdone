import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import { connect } from 'react-redux';
import {Provider} from "react-redux"

import OrgCreateForm from './component/org-create-form';
import createStoreWithMiddleware from './store/configure';
import * as form from './action/form'

const store = createStoreWithMiddleware();

class Landing extends Component {
  constructor() {
    super()
  }

  handleSubmit(orgName) {
    if (orgName.length == 0) {
      this.props.dispatch(form.clientError("Org name is required"))
      return
    }
    this.props.dispatch(form.submit(orgName, (key) => {
      window.location="/" + key;
    }))
  }

  render() {
    return (
      <OrgCreateForm
        submit={(o) => this.handleSubmit(o)}
        orgNameError={this.props.clientError}
      />
    )
  }
}

Landing = connect(({form}) => {
  return {
    key: form.key,
    clientError: form.clientError
  }
})(Landing)

window.addEventListener('load', () => {
  var parent = document.getElementById('teamdone-create-org');
  if (parent !== null) {
    ReactDOM.render(
      <Provider store={store}>
        <MuiThemeProvider>
          <Landing />
        </MuiThemeProvider>
      </Provider>,
      parent
    );
  }
});
