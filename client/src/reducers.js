import { combineReducers } from 'redux';
import messages, * as fromMessages from './ducks/messages/reducers';
import messageInput, * as fromMessageInput from './ducks/message-input/reducers';

const dummy = (state = {}, action) => state;

export default combineReducers({
  messages,
  messageInput
});

export const messages_getDisplayMessages = state => fromMessages.getDisplayMessages(state.messages);
export const messages_getLoadingState = state => fromMessages.getLoadingState(state.messages);

export const messageInput_getInput = state => fromMessageInput.getInput(state.messageInput);
