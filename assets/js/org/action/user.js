import * as api from '../../shared/api';

function openForm() {
  return {
    type: "USER_FORM",
    open: true
  }
}

function closeForm() {
  return {
    type: "USER_FORM",
    open: false
  }
}

function setUser(id, name) {
  return {
    type: "USER_SET",
    id: id,
    name: name
  }
}

export function userNameError(message) {
  return {
    type: "USER_NAME_ERROR",
    message: message
  }
}

export function initialize() {
  return (dispatch) => {
    api.sessionGet()
    .then((response) => {
      let status = response.status
      if (status >= 200 && status < 300) {
        response.json().then((json) => {
          if (!json.hasOwnProperty("UserId")) {
            dispatch(openForm());
          }
        });
      } else {
        response.json().then((json) => console.error(json['Message']));
      }
    });
  }
}

export function submit(userName) {
  return (dispatch) => {
    api.sessionCreate(userName)
    .then((response) => {
      let status = response.status
      if (status >= 200 && status < 300) {
        response.json().then((json) => {
          dispatch(setUser(json['UserId'], json['UserName']));
          dispatch(closeForm());
        });
      } else if (status === 400) {
        response.json().then((json) => {
          let message = json['Message'];
          dispatch(userNameError(message));
        });
      } else {
        response.json().then((json) => console.error(json['Message']));
      }
    });
  }
}
