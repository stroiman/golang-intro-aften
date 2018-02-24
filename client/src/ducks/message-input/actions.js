import uuid from 'uuid';
import * as api from '../api';

export const addMessage = (message) => async (dispatch) => {
  await api.post("/api/messages", {
    id: uuid.v4(),
    message
  });
}
