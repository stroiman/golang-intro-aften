import { combineReducers } from 'redux';

const input = (state = "", action) => {
  switch (action.type) {
    case "MESSAGE_INPUT_SET": return action.payload;
    case "MESSAGE_POSTED": return "";
    default: return state;
  }
};

export default combineReducers({
  input
})

export const getInput = state => state.input;
