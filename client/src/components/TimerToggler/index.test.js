import React from 'react';
import TimerToggler from '.';
import { smokeTest } from '../../testHelpers/reactTestHelpers';
import { shallow } from 'enzyme';
import { useSinon } from '../../testHelpers';
import { createMockStore } from '../../testHelpers/reduxHelpers';
import * as actions from './actions';

describe('TimerToggler', () => {
  it('Renders without crashing', () => {
    smokeTest(TimerToggler);
  });

  describe('Clicking', () => {
    useSinon();

    it('dispatches a toggle action', function() {
      const store = createMockStore();
      const dummyAction = {type: "DUMMY"};
      this.sinon.stub(actions, "togglePolling").returns(dummyAction)
      const wrapper = shallow(<TimerToggler store={store} />);
      wrapper.dive().simulate('click', { preventDefault: () => ({}) });
      expect(store.getActions()).to.deep.equal([dummyAction]);
    });
  });
});
