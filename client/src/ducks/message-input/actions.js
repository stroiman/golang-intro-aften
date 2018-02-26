import uuid from 'uuid';
import * as api from '../api';
import * as getters from '../../reducers';

export const setInput = (message) => ({
  type: "MESSAGE_INPUT_SET",
  payload: message
});

export const addMessage = () => async (dispatch, getState) => {
  const message = getters.messageInput_getInput(getState());
  const result = await api.post("/api/messages", {
    id: uuid.v4(),
    message
  });
  if (result.ok) {
    dispatch({
      type: "MESSAGE_POSTED"
    });
  }
}
