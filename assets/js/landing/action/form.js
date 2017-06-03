import {getRequestToken} from '../../shared/request-token';

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
    let url = "/i/org/create"
    fetch(url, {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json',
        'X-Request-Token': getRequestToken()
      },
      body: JSON.stringify({
        OrgName: orgName
      })
    })
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
