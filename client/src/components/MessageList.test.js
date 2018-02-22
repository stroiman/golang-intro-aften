import React from 'react';
import MessageList from './MessageList';
import { smokeTest } from '../testHelpers/reactTestHelpers';

describe("MessageList", () => {
  it("Renders without crashing", () => {
    smokeTest(MessageList);
  });
});
