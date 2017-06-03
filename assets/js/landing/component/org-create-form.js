import React, { Component } from 'react';
import Dialog from 'material-ui/Dialog';
import MenuItem from 'material-ui/MenuItem';
import PropTypes from 'prop-types';
import SelectField from 'material-ui/SelectField';
import TextField from 'material-ui/TextField';
import FlatButton from 'material-ui/FlatButton';

export default class OrgCreateForm extends Component {
  constructor() {
    super();
    this.state = {
      orgName: "",
      orgNameError: ""
    }
  }

  handleCreate() {
    if (typeof this.props.submit === "undefined") {
      return
    }
    let orgName = this.orgName.getValue();
    this.props.submit(orgName);
  }

  render() {
    return (
      <div>
        <TextField
          ref={(e) => { this.orgName = e; }}
          hintText="Org name"
          errorText={this.props.orgNameError}
          floatingLabelText="Org Name"
          name="orgName"
          fullWidth={true}
        />
        <br />
        <div style={{ textAlign: "right" }}>
          <FlatButton label="Create" onTouchTap={() => this.handleCreate()} />
        </div>
      </div>
    );
  }
}

OrgCreateForm.defaultProps = {
  orgName: "",
  orgNameError: "",
};

OrgCreateForm.propTypes = {
  orgName: PropTypes.string.isRequired,
  orgNameError: PropTypes.string.isRequired,
  submit: PropTypes.func
};
