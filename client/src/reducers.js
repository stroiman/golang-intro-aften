import { combineReducers } from 'redux';
import messages, * as fromMessages from './ducks/messages/reducers';
import messageInput, * as fromMessageInput from './ducks/message-input/reducers';
import polling, * as fromPolling from './ducks/polling/reducers';

const dummy = (state = {}, action) => state;

export default combineReducers({
  messages,
  messageInput,
  polling
});

export const messages_getDisplayMessages = state => fromMessages.getDisplayMessages(state.messages);
export const messages_getLoadingState = state => fromMessages.getLoadingState(state.messages);
export const polling_isPolling = state => fromPolling.getIsPolling(state.polling);
export const polling_getTimerId = state => fromPolling.getTimerId(state.polling);

export const messageInput_getInput = state => fromMessageInput.getInput(state.messageInput);
