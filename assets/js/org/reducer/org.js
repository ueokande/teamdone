const initialiState = {
  orgId: null,
  orgName: ""
}

export default function org(state = initialiState, action) {
  switch (action.type) {
  case "ORG_SET":
    return Object.assign({}, state, {
      orgId: action.id,
      orgName: action.name,
      orgKey: action.key
    });
  default:
    return state;
  }
}
