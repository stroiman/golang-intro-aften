import { combineReducers } from 'redux';
import messages, * as fromMessages from './ducks/messages/reducers';

const dummy = (state = {}, action) => state;

export default combineReducers({
  messages
});

export const messages_getDisplayMessages = state => fromMessages.getDisplayMessages(state.messages);
export const messages_getLoadingState = state => fromMessages.getLoadingState(state.messages);
