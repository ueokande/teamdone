function getRequestToken() {
  var fields = document.cookie.split(';');
  for (var i = 0; i < fields.length; ++i) {
    var kv = fields[i].split('=');
    if (kv[0].trim() === "_request_token") {
      return kv[1].trim();
    }
  }
  return null;
}

export { getRequestToken };
