import { combineReducers } from 'redux';
import { normalize, schema } from 'normalizr';
import { includes } from 'lodash';

const message = new schema.Entity('message');
const messages = new schema.Array(message);

const displayMessages = (state = {ids: [], messages: {}}, action) => {
  switch (action.type) {
    case "MESSAGES_FETCH_COMPLETED": {
      const normalized = normalize(action.payload.messages, messages);
      return { ids: normalized.result, messages: normalized.entities.message }
    }
    case "MESSAGE_RECEIVED": {
      const normalized = normalize(action.payload, message);
      const messages = {...state.messages, ...normalized.entities.message};
      const id = normalized.result;
      const ids = includes(state.ids, id) ? state.ids : [...state.ids, id];
      return { ids, messages };
    }
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

export const getDisplayMessages = state => state.displayMessages.ids.map(id => state.displayMessages.messages[id]);
export const getLoadingState = state => state.loadingState;
