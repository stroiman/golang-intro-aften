import * as redux from 'redux';
import thunk from 'redux-thunk';
import reducers from './reducers';
import logger from 'redux-logger';

export const rootReducer = reducers;


export const createStore = (options = {}) => {
  let middlewares = [thunk];
  const { initialState } = options;
  if (options.middlewares) {
    middlewares = [...middlewares, ...options.middlewares];
  };

  if (options.logging) {
    middlewares.push(logger);
  }

  if (initialState) {
    return redux.createStore(rootReducer, initialState, redux.applyMiddleware(...middlewares));
  } else {
    return redux.createStore(rootReducer, redux.applyMiddleware(...middlewares));
  }
};

