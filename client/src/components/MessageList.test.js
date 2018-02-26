import React from 'react';
import MessageList, { Message } from './MessageList';
import { smokeTest } from '../testHelpers/reactTestHelpers';
import { useRedux, createMockStore } from '../testHelpers/reduxHelpers';
import { shallow, mount } from 'enzyme';
import { createMessage } from '../testHelpers/factories';
import * as actions from '../ducks/messages/actions';
import * as editActions from '../ducks/message-input/actions';

describe("MessageList", () => {
  it("Renders without crashing", () => {
    smokeTest(MessageList);
  });

  describe("Editing", () => {
    const message = createMessage();
    const store = createMockStore(
      actions.messagesLoaded([message])
    );
    const wrapper = mount(<MessageList store={store} />);
    wrapper.find("[role='edit']").simulate('click', { preventDefault: () => {} });
    const actual = store.getActions();
    const expected = [editActions.editMessage(message)];
    expect(actual).to.deep.equal(expected);
  });

  describe("Rendering messages", () => {
    useRedux();

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
});
