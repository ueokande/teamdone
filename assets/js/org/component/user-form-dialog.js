import React, { Component } from 'react';
import Dialog from 'material-ui/Dialog';
import PropTypes from 'prop-types';
import TextField from 'material-ui/TextField';
import FlatButton from 'material-ui/FlatButton';

export default class UserFormDialog extends Component {
  constructor() {
    super();
  }

  handleCreate() {
    if (typeof this.props.submit === 'undefined') {
      return
    }
    let userName = this.userName.getValue();
    this.props.submit(userName)
  }

  render() {
    return (
      <Dialog open={this.props.open} >
        <TextField
          ref={(e) => { this.userName = e; }}
          hintText="Alice"
          floatingLabelText="Name"
          errorText={this.props.userNameError}
          name="name"
          fullWidth={true}
        /><br />
        <div style={{ textAlign: "right" }}>
          <FlatButton label="OK" onTouchTap={() => this.handleCreate()} />
        </div>
      </Dialog>
    );
  }
}

UserFormDialog.defaultProps = {
};

UserFormDialog.propTypes = {
  open: PropTypes.bool.isRequired,
  userNameError: PropTypes.string,
  submit: PropTypes.func
};

