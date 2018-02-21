import App from './App';
import { smokeTest } from './testHelpers/reactTestHelpers';

describe('app.js', () => {
  it('renders without crashing', () => {
    smokeTest(App);
  });
});
