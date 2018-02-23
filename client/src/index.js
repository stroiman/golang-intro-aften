import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './index.css';
import App from './App';
import { unregister } from './registerServiceWorker';
import { createStore } from './store';
import * as actions from './ducks/messages/actions';

const store = createStore();
store.dispatch(actions.fetchMessages());

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>, document.getElementById('root'));
unregister();
