import * as api from '../../shared/api';

export function clientError(message) {
  return {
    type: "CLIENT_ERROR",
    message: message,
  };
}

export function serverError(message) {
  return {
    type: "SERVER_ERROR",
    message: message,
  };
}

export function submit(orgName, successed) {
  return (dispatch) => {
    api.orgCreate(orgName)
    .then((response) => {
      let status = response.status;
      if (status >= 200 && status < 300) {
        response.json().then((json) => successed(json['Key']));
      } else if (status >= 400 && status < 500) {
        response.json().then((json) => dispatch(clientError(json['Message'])));
      } else {
        response.json().then((json) => dispatch(serverError(json['Message'])));
      }
    })
  }
}
