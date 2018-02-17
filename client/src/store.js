import * as redux from 'redux';
import thunk from 'redux-thunk';
import reducers from './reducers';

export const createStore = () => {
  return redux.createStore(reducers, redux.applyMiddleware(thunk));
};

