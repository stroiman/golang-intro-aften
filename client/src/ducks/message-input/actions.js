import uuid from 'uuid';
import * as api from '../api';
import * as getters from '../../reducers';

export const setInput = (message) => ({
  type: "MESSAGE_INPUT_SET",
  payload: message
});

export const addMessage = () => async (dispatch, getState) => {
  const message = getters.messageInput_getInput(getState());
  await api.post("/api/messages", {
    id: uuid.v4(),
    message
  });
}
