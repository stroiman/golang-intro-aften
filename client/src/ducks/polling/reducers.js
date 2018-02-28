import { combineReducers } from 'redux';

const timerId = (state = null, action) => {
  switch(action.type) {
    case "MESSAGES_SETTIMERID": return action.payload;
    default: return state;
  }
}

export default combineReducers({
  timerId
});

export const getIsPolling = state => {
  return state.timerId !== null;
}
export const getTimerId = state => state.timerId;
