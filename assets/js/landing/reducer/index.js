import {combineReducers} from "redux"

const initialState = {
  orgName: "",
  clientError: "",
  serverError: "",
}

function form(state = initialState, action) {
  switch (action.type) {
  case "SUBMIT_OK":
    return Object.assign({}, state, {
      key: action.key
    });
  case "CLIENT_ERROR":
    return Object.assign({}, state, {
      clientError: action.message
    });
  case "SERVER_ERROR":
    return Object.assign({}, state, {
      serverError: action.message
    });
  default:
    return state;
  }
}

export default combineReducers({
  form
});
