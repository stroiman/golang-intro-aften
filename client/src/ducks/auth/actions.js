export const loginUser = ({username}) => ({
  type: "AUTH_LOGIN_USER",
  payload: {username}
});

export const setUserNameInput = input => ({
  type: "AUTH_SET_USERNAME",
  payload: input
});
