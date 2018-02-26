import * as getters from '../../reducers';
import * as actions from './actions';
import fetchMock from 'fetch-mock';
import { useRedux } from '../../testHelpers/reduxHelpers';


describe("messages duck", () => {
  useRedux();

  afterEach(function() {
    fetchMock.restore();
  });

  describe("Adding new messages", () => {
    it("Posts", async function() {
      fetchMock.get("/api/messages", {status: 200, body: []});
      fetchMock.post("/api/messages", {status: 200, body: {status: 'ok'}}, {name: "msgPost"});
      this.dispatch(actions.setInput("foobar"));
      await this.dispatch(actions.addMessage())
      const lastCall = fetchMock.lastCall("msgPost").pop();
      const lastBody = JSON.parse(lastCall.body);
      expect(lastBody.message).to.equal("foobar");
    });

    it('clears the message input', async function() {
      fetchMock.get("/api/messages", {status: 200, body: []});
      fetchMock.post("/api/messages", {status: 200, body: {status: 'ok'}}, {name: "msgPost"});
      this.dispatch(actions.setInput("foobar"));
      await this.dispatch(actions.addMessage())
      const actual = getters.messageInput_getInput(this.getState());
      expect(actual).to.equal('');
    });
  });
});
