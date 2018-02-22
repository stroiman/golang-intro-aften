import { combineReducers } from 'redux';

const displayMessages = (state = [], action) => state;

export default combineReducers({
  displayMessages
})

export const getDisplayMessages = state => state.displayMessages;
export const getLoadingState = state => "NOT_LOADED";
