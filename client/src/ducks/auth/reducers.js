export default (state = null, action) => {
  switch(action.type) {
    case "AUTH_LOGIN_USER": return action.payload;
    default: return state;
  }
}
