import * as uuid from 'uuid'
import * as getters from '../../reducers';
import * as actions from './actions';
import fetchMock from 'fetch-mock';
import { useRedux } from '../../testHelpers/reduxHelpers';

const unsettledPromise = () => new Promise((resolve, reject) => {});

const createMessage = () => ({
  id: uuid.v4()
});

describe("messages duck", () => {
  useRedux();

  beforeEach(function() {
    this.getLoadingState = () => getters.messages_getLoadingState(this.getState());
    this.getDisplayMessages = () => getters.messages_getDisplayMessages(this.getState());
  });

  afterEach(() => fetchMock.restore());

  describe("initial state", () => {
    it("is an empty array", function() {
      const actual = this.getDisplayMessages();
      expect(actual).to.deep.equal([]);
    });

    it("has loadingState=NOT_LOADED", function() {
      const actual = this.getLoadingState();
      expect(actual).to.equal("NOT_LOADED");
    });
  });

  describe("Loading messages", () => {
    context("When server has not yet replied", () => {
      it("has loadingState=LOADING", async function() {
        fetchMock.get("/api/messages", unsettledPromise());
        this.dispatch(actions.fetchMessages());
        const actual = this.getLoadingState();
        expect(actual).to.equal("LOADING");
      });
    });

    context("When server has replied", () => {
      beforeEach(() => {
        fetchMock.get("/api/messages", Promise.resolve([createMessage(), createMessage()]));
      });

      it("has loadingState=LOADED", async function() {
        await this.dispatch(actions.fetchMessages());
        const actual = this.getLoadingState();
        expect(actual).to.equal("LOADED");
      });

      it("has display messages as the data from the server", async function() {
        await this.dispatch(actions.fetchMessages());
        const actual = this.getDisplayMessages();
        expect(actual).to.have.length(2);
      });
    });

    context("When server fails", () => {
      beforeEach(() => {
        fetchMock.get("/api/messages", {status: 500});
      })

      it("has loadingState=FAILED", async function() {
        await this.dispatch(actions.fetchMessages());
        const actual = this.getLoadingState();
        expect(actual).to.equal("FAILED");
      });
    });
  });
});
