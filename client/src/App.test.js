import React from 'react';
import App, { MessagesPage } from './App';
import LoginPage from './pages/LoginPage';
import { smokeTest } from './testHelpers/reactTestHelpers';
import { useRedux } from './testHelpers/reduxHelpers';
import { shallow } from 'enzyme';
import * as authActions from './ducks/auth/actions';

describe('app.js', () => {
  it('renders without crashing', () => {
    smokeTest(App);
  });

  useRedux();

  context('user is not logged in', () => {
    it('renders the login page', function() {
      const wrapper = shallow(<App store={this.store}/>).dive();
      const page = wrapper.find(LoginPage);
      expect(page).to.have.length(1)
    });
  });

  context('user is logged in', () => {
    beforeEach(function() {
      this.dispatch(authActions.loginUser({}));
    });

    it('renders the messages page', function() {
      const wrapper = shallow(<App store={this.store}/>).dive();
      const page = wrapper.find(MessagesPage);
      expect(page).to.have.length(1)
    });
  });
});
