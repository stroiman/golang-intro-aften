import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import sinon from 'sinon';
import * as actions from './features/blogs/actions';
import fetchMock from 'fetch-mock';

describe('app.js', () => {
  beforeEach(function() {
    fetchMock.mock("*", {});
  });

  afterEach(function() {
    fetchMock.restore();
  });

  it('renders without crashing', () => {
    const div = document.createElement('div');
    const store = createStore();
    ReactDOM.render(<Provider store={store}><App /></Provider>, div);
    ReactDOM.unmountComponentAtNode(div);
  });
});
