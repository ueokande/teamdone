import * as api from '../../shared/api';

function setOrg(id, name, key) {
  return {
    type: "ORG_SET",
    id: id,
    name: name,
    key: key
  }
}

export function initialize(key) {
  return (dispatch) => {
    api.orgGet(key)
    .then((response) => {
      let status = response.status;
      if (status >= 200 && status < 300) {
        response.json().then((json) => {
          dispatch(setOrg(json["OrgId"], json["OrgName"], json["OrgKey"]));
        })
      }
    })
  }
}
