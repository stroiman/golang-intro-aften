import { combineReducers } from 'redux';

const input = (state = "", action) => {
  switch (action.type) {
    case "MESSAGE_INPUT_SET": return action.payload;
    case "MESSAGE_EDIT": return action.payload.message;
    case "MESSAGE_EDIT_CANCEL": return "";
    case "MESSAGE_POSTED": return "";
    default: return state;
  }
};

const editingMessageId = (state = null, action) => {
  switch (action.type) {
    case "MESSAGE_EDIT": return action.payload.id;
    case "MESSAGE_EDIT_CANCEL": return null;
    case "MESSAGE_POSTED": return null;
    default: return state;
  }
}

export default combineReducers({
  input,
  editingMessageId
})

export const getInput = state => state.input;
export const isEditing = state => state.editingMessageId !== null;
