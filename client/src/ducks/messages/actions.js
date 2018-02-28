import uuid from 'uuid';
import * as api from '../api';
import * as getters from '../../reducers';

export const messagesLoaded = messages => ({
  type: "MESSAGES_FETCH_COMPLETED",
  payload: { messages }
});

export const fetchMessages = () => async (dispatch) => {
  dispatch({
    type: "MESSAGES_FETCH"
  });
  const response = await api.get("/api/messages");
  //const response = await fetch("/api/messages");
  if (!response.ok) {
    return dispatch({
      type: "MESSAGES_FETCH_FAILED",
    })
  }
  const messages = await response.json;
  dispatch(messagesLoaded(messages));
};
