import React from 'react';
import { createMockStore } from '../testHelpers/reduxHelpers';
import MessageInput from './MessageInput';
import { mount } from 'enzyme';
import * as actions from '../ducks/message-input/actions';
import { useSinon } from '../testHelpers';

const mockAction = { type: "MOCK_ACTION" };

describe("MessageInput", () => {
  useSinon();

  beforeEach(function() {
    this.addMessageStub = this.sinon.stub(actions, "addMessage").returns(mockAction);
    this.store = createMockStore();
  });

  context("use has entered 'foobar' and submitted", () => {
    beforeEach(function() {
      this.wrapper = mount(<MessageInput store={this.store} />);
      this.wrapper.find("input").simulate("change", { target: { value: 'foobar' } } );
      this.wrapper.find("form").simulate("submit", { preventDefault: () => {} });
    });

    it("Sends an action on submit", function() {
      expect(this.store.getActions()).to.deep.equal([mockAction]);
    });

    it("Passes the message to the action", function() {
      expect(this.addMessageStub).to.have.been.calledWith("foobar");
    });
  });
});
