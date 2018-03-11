import { combineReducers } from 'redux';
import { normalize, schema } from 'normalizr';
import { includes } from 'lodash';

const ids = (state = [], action) => {
  switch (action.type) {
    case "MESSAGES_FETCH_COMPLETED": return action.payload.result;
    case "MESSAGE_RECEIVED": {
      const id = action.payload.result;
      return includes(state, id) ? state : [...state, id];
    }
    default: return state;
  }
}

const messages = (state = {}, action) => {
  switch (action.type) {
    case "MESSAGE_RECEIVED":
    case "MESSAGES_FETCH_COMPLETED":
      return {...state, ...action.payload.entities.message };
    default: return state;
  }
}

const displayMessages = combineReducers({
  ids, messages,
});

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

export const getDisplayMessages = state => state.displayMessages.ids.map(id => state.displayMessages.messages[id]);
export const getLoadingState = state => state.loadingState;
