import * as api from '../../shared/api';

export function openForm() {
  return {
    type: "USER_FORM",
    open: true
  }
}

export function initialize() {
  return (dispatch) => {
    api.sessionGet()
    .then((response) => {
      let status = response.status
      if (status >= 200 && status < 300) {
        response.json().then((json) => dispatch(json['Key']));
      } else if (status === 404) {
        dispatch(openForm());
      } else {
        response.json().then((json) => console.error(json['Message']));
      }
    });
  }
}
