import React from 'react';
import MessageList, { Message } from './MessageList';
import { smokeTest } from '../testHelpers/reactTestHelpers';
import { useRedux } from '../testHelpers/reduxHelpers';
import { shallow } from 'enzyme';
import { createMessage } from '../testHelpers/factories';
import * as actions from '../ducks/messages/actions';

describe("MessageList", () => {
  useRedux();

  it("Renders without crashing", () => {
    smokeTest(MessageList);
  });

  context("Initialized with two messages", () => {
    it("renders two messages", function() {
      this.dispatch(actions.messagesLoaded([createMessage(), createMessage()]));
      const wrapper = shallow(<MessageList store={this.store} />).dive();
      expect(wrapper.find(Message)).to.have.length(2);
    })
  });

  context("Initialized with two messages AFTER construction", () => {
    it("renders two messages", function() {
      const wrapper = shallow(
        <MessageList store={this.store}/>
      );
      this.dispatch(actions.messagesLoaded([createMessage(), createMessage()]));
      expect(wrapper.update().dive().find(Message)).to.have.length(2);
    })
  });
});
