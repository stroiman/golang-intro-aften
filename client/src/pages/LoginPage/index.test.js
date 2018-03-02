import React from 'react';
import { mount } from 'enzyme'
import LoginPage from '.';
import { smokeTest } from '../../testHelpers/reactTestHelpers';
import * as actions from '../../ducks/auth/actions';
import { useSinon } from '../../testHelpers';
import { createMockStore } from '../../testHelpers/reduxHelpers';

describe("LoginPage", () => {
  it('renders without crashing', () => {
    smokeTest(LoginPage);
  });

  describe('filling out form', function() {
    useSinon();

    it('dispatches an action', function() {
      const dummyAction = { type: 'DUMMY' };
      //this.sinon.stub(actions, "loginUser").returns(dummyAction);
      const store = createMockStore();
      const wrapper = mount(<LoginPage store={store} />);
      const input = wrapper.find("input");
      input.simulate('change', { target: { value: 'foobar' } });
      const button = wrapper.find("button");
      button.simulate('click', { preventDefault: () => {} });
      expect(store.getActions()).to.deep.equal([
        actions.loginUser({
          username: 'foobar'
        })
      ]);
    });
  });
});
