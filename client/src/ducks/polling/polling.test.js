import sinon from 'sinon';
import lolex from 'lolex';
import { useRedux } from '../../testHelpers/reduxHelpers';
import { useSinon } from '../../testHelpers';
import * as getters from '../../reducers';
import * as actions from './actions';
import * as messageActions from '../messages/actions';

describe("Polling", () => {
  useRedux();
  useSinon();

  beforeEach(function() {
    this.clock = lolex.install({
      toFake: ["setInterval", "clearInterval"]
    });
  });

  afterEach(function() {
    this.clock.uninstall();
  });


  describe("Initial state", () => {
    it("is not polling", function() {
      expect(getters.polling_isPolling(this.getState())).to.be.false;
    });
  });

  describe("Togling polling", () => {
    beforeEach(async function() {
      await this.dispatch(actions.togglePolling());
    });

    it('Sets polling to true', function() {
      expect(getters.polling_isPolling(this.getState())).to.be.true;
    });

    context('2 seconds have passed', () => {
      beforeEach(function() {
        this.actionMock = this.sinon.mock(messageActions).expects("fetchMessages").returns({type:"DUMMY"}).twice()
        this.clock.tick(2000);
      });

      it('dispatches FETCH actions', function() {
        this.actionMock.verify();
      });

      describe("toggling polling again", () => {
        beforeEach(async function() {
          await this.dispatch(actions.togglePolling());
        });

        it("does not dispatch any more actions", function() {
          this.clock.tick(2000);
          this.actionMock.verify();
        });
      });
    });
  });
});
