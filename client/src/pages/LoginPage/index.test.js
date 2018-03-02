import React from 'react';
import LoginPage from '.';
import { smokeTest } from '../../testHelpers/reactTestHelpers';

describe("LoginPage", () => {
  it('renders without crashing', () => {
    smokeTest(LoginPage);
  });
});
