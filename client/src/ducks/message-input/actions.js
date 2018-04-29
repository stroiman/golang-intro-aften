import uuid from 'uuid';
import * as api from '../api';
import * as getters from '../../reducers';

export const setInput = (message) => ({
  type: "MESSAGE_INPUT_SET",
  payload: message
});

export const addMessage = () => async (dispatch, getState) => {
  const state = getState();
  const message = getters.messageInput_getInput(state);
  const messageId = getState().messageInput.editingMessageId;
  const userName = getters.auth_getUserName(state);

  if (messageId) {
    const body = {
      id: messageId, message,
      userName: getters.auth_getUserName(getState()),
    }
    const result = await api.put(`/api/messages/${messageId}`, body);
    if (result.ok) {
      dispatch({
        type: "MESSAGE_POSTED"
      });
    }
    return;
  }
  const result = await api.post("/api/messages", {
    // id: uuid.v4(),
    userName: getters.auth_getUserName(getState()),
    message,
  });
  if (result.ok) {
    dispatch({
      type: "MESSAGE_POSTED"
    });
  }
}

export const editMessage = message => ({
  type: "MESSAGE_EDIT",
  payload: message
});

export const cancelEditing = () => ({
  type: "MESSAGE_EDIT_CANCEL",
});
