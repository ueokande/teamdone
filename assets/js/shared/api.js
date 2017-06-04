import {getRequestToken} from '../shared/request-token';

function post(url, params) {
    return fetch(url, {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json',
        'X-Request-Token': getRequestToken()
      },
      body: JSON.stringify(params)
    })
}

export function orgCreate(orgName) {
  return post('/i/org/create', {
    OrgName: orgName
  })
}
