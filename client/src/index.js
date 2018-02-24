import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './index.css';
import App from './App';
import { unregister } from './registerServiceWorker';
import { createStore } from './store';
import * as actions from './ducks/messages/actions';

const store = createStore({logging: true});
store.dispatch(actions.fetchMessages());

setInterval(() => store.dispatch(actions.fetchMessages()), 1000);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>, document.getElementById('root'));
unregister();
