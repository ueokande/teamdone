const initialState = {
  userId: null,
  userName: "",

  formOpen: false,
}

export default function tasks(state = initialState, action) {
  switch (action.type) {
  case "USER_FORM":
    return Object.assign({}, state, {
      formOpen: action.open
    });
  default:
    return state;
  }
}
