import * as getters from '../../reducers';
import * as actions from './actions';
import fetchMock from 'fetch-mock';
import { useRedux } from '../../testHelpers/reduxHelpers';
import { createMessage } from '../../testHelpers/factories';

const unsettledPromise = () => new Promise((resolve, reject) => {});

describe("messages duck", () => {
  useRedux();

  beforeEach(function() {
    this.getLoadingState = () => getters.messages_getLoadingState(this.getState());
    this.getDisplayMessages = () => getters.messages_getDisplayMessages(this.getState());
  });

  afterEach(() => fetchMock.restore());

  describe("initial state", () => {
    it("is an empty array", function() {
      this.getDisplayMessages().should.deep.equal([]);
    });

    it("has loadingState=NOT_LOADED", function() {
      this.getLoadingState().should.equal("NOT_LOADED");
    });
  });

  describe("Loading messages", () => {
    context("When server has not yet replied", () => {
      it("has loadingState=LOADING", async function() {
        fetchMock.get("/api/messages", unsettledPromise());
        this.dispatch(actions.fetchMessages());
        this.getLoadingState().should.equal("LOADING");
      });
    });

    context("When server has replied", () => {
      beforeEach(async function() {
        fetchMock.get("/api/messages", Promise.resolve([createMessage(), createMessage()]));
        await this.dispatch(actions.fetchMessages());
      });

      it("has loadingState=LOADED", async function() {
        this.getLoadingState().should.equal("LOADED");
      });

      it("has display messages as the data from the server", async function() {
        this.getDisplayMessages().should.have.length(2);
      });
    });

    context("When server fails", () => {
      beforeEach(() => {
        fetchMock.get("/api/messages", {status: 500});
      })

      it("has loadingState=FAILED", async function() {
        await this.dispatch(actions.fetchMessages());
        this.getLoadingState().should.equal("FAILED");
      });
    });
  });

  describe('message received', function() {
    context('empty context', function() {
      it('creates a new message', function() {
        this.dispatch(actions.messageReceived(createMessage()));
        this.getDisplayMessages().should.have.length(1);
      });
    });

    context('existing message is updated', function() {
      it('does not create a new message', function() {
        const message = createMessage();
        this.dispatch(actions.messageReceived(message));
        this.dispatch(actions.messageReceived({...message, message: 'new message'}));
        this.getDisplayMessages().should.have.length(1);
      });
    });
  });
});
