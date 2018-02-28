import { combineReducers } from 'redux';

const displayMessages = (state = [], action) => {
  switch (action.type) {
    case "MESSAGES_FETCH_COMPLETED": return action.payload.messages;
    default: return state;
  }
}

const loadingState = (state = "NOT_LOADED", action) => {
  switch(action.type) {
    case "MESSAGES_FETCH": return "LOADING";
    case "MESSAGES_FETCH_COMPLETED": return "LOADED";
    case "MESSAGES_FETCH_FAILED": return "FAILED";
    default: return state;
  }
};

export default combineReducers({
  displayMessages,
  loadingState,
})

export const getDisplayMessages = state => state.displayMessages;
export const getLoadingState = state => state.loadingState;
