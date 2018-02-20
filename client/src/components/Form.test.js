import React from 'react';
import { mount, shallow } from 'enzyme';
import sinon from 'sinon';
import Form from './Form';

describe("Form", () => {
  it("propagates events up", () => {
    const spy = sinon.spy();
    const wrapper = mount(
      <Form values={{}} onChange={spy} >
        <Form.TextField name="test" />
      </Form>);
    const input = wrapper.find(Form.TextField).find("input");
    input.simulate('change', { target: { value: "foobar" } });
    expect(spy).to.have.been.calledWith("test", "foobar");
  });
});
