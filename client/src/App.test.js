import App from './App';
import fetchMock from 'fetch-mock';
import { smokeTest } from './testHelpers/reactTestHelpers';

describe('app.js', () => {
  it('renders without crashing', () => {
    smokeTest(App);
  });
});
