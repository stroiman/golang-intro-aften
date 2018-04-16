import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './index.css';
import App from './App';
import { unregister } from './registerServiceWorker';
import { createStore } from './store';
import * as actions from './ducks/auth/actions';
import * as messageActions from './ducks/messages/actions';
import * as pollActions from './ducks/polling/actions';

const store = createStore({logging: true});
if (process.env.NODE_ENV === "development") {
  store.dispatch(actions.loginUser({username: 'stroiman'}));
  // store.dispatch(pollActions.startWebSocket());
  store.dispatch(messageActions.fetchMessages());
}

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>, document.getElementById('root'));
unregister();
