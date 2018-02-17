import { combineReducers } from 'redux';

const loadingState = (state = "NOT_LOADING", action) => {
  switch(action.type){
    case "BLOGS_LOAD": return "LOADING";
    case "BLOGS_LOADED": return "LOADED";
    default: return state;
  }
};

export default combineReducers({
  loadingState
});

export const getLoadingState = state => state.loadingState;
