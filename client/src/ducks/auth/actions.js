export const loginUser = (user) => ({
  type: "AUTH_LOGIN_USER",
  payload: user
});

export const setUserNameInput = input => ({
  type: "AUTH_SET_USERNAME",
  payload: input
});
