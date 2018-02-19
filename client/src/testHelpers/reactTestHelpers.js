import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore } from '../store';

/**
 * Simply tries to render a component to verify it renders
 */
export const smokeTest = (Component) => {
  const div = document.createElement('div');
  const store = createStore({disableDispatch: true});
  store.dispatch = () => {};
  ReactDOM.render(
    <Provider store={store}>
      <Component />
    </Provider>, div);
  ReactDOM.unmountComponentAtNode(div);
}
