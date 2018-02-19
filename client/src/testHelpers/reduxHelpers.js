import { rootReducer, createStore } from '../store';
import deepFreeze from 'deep-freeze';

const initialState = deepFreeze(rootReducer(undefined, {type:null}));

const deepFreezeMiddleware = store => next => action => {
  return deepFreeze(next(action));
}
/**
 * Helper function for using a redux store in tests
 *
 * options: disableDispatch. Replaces `dispatch` with a dummy
 * function. Usefule if you don't want to react to event dispatched
 * during mounting
 */
export const useRedux = (options = {}) => {
  beforeEach(function() {
    this.store = createStore({initialState, middlewares: [deepFreezeMiddleware]});
    this.dispatch = this.store.dispatch;
    this.getState = this.store.getState;
    this.disableDispatch = () => this.store.dispatch = () => {};
    if (options.disableDispatch) this.disableDispatch();
  })
}

