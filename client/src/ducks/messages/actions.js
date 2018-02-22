export const fetchMessages = () => async (dispatch) => {
  dispatch({
    type: "MESSAGES_FETCH"
  });
  const response = await fetch("/api/messages");
  if (!response.ok) {
    return dispatch({
      type: "MESSAGES_FETCH_FAILED",
    })
  }
  const json = await response.json();
  dispatch({
    type: "MESSAGES_FETCH_COMPLETED",
    payload: {
      messages: json
    }
  });
};
