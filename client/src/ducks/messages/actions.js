import uuid from 'uuid';
import { normalize, schema } from 'normalizr';
import * as api from '../api';
import * as getters from '../../reducers';

const messageSchema = new schema.Entity('message');
const messageListSchema = new schema.Array(messageSchema);

export const messageReceived = message => ({
  type: "MESSAGE_RECEIVED",
  payload: normalize(message, messageSchema),
});

export const messagesLoaded = messages => ({
  type: "MESSAGES_FETCH_COMPLETED",
  payload: normalize(messages, messageListSchema)
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
