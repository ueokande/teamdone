const initialState = {
  userId: null,
  userName: "",

  formOpen: false,
  formNameError: "",
}

export default function user(state = initialState, action) {
  switch (action.type) {
  case "USER_FORM":
    return Object.assign({}, state, {
      formOpen: action.open
    });
  case "USER_SET":
    return Object.assign({}, state, {
      userId: state.id,
      userName: state.name
    });
  case "USER_NAME_ERROR":
    return Object.assign({}, state, {
      formNameError: action.message
    });
  default:
    return state;
  }
}
