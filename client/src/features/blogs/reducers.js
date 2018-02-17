import { combineReducers } from 'redux';

const loadingState = (state = "NOT_LOADING", action) => {
  switch(action.type){
    case "BLOGS_LOAD": return action.payload.status;
    default: return state;
  }
};

export default combineReducers({
  loadingState
});

export const getLoadingState = state => state.loadingState;
