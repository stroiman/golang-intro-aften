import * as getters from '../../reducers';
import * as actions from './actions';
import * as authActions from '../auth/actions';
import fetchMock from 'fetch-mock';
import { useRedux } from '../../testHelpers/reduxHelpers';

import { createMessage } from '../../testHelpers/factories';

fetchMock.lastBodyJson = function() {
  const opts = fetchMock.lastOptions.call(this, ...arguments)
  return JSON.parse(opts.body);
}

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
      const lastBody = fetchMock.lastBodyJson("msgPost");
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

    it('puts the username on the message', async function() {
      fetchMock.post("/api/messages", {status: 200, body: {status: 'ok'}}, {name: "msgPost"});
      this.dispatch(authActions.loginUser({username: "johndoe"}));
      console.log(this.getState());
      this.dispatch(actions.setInput("foobar"));
      await this.dispatch(actions.addMessage())
      const lastBody = fetchMock.lastBodyJson("msgPost");
      expect(lastBody.userName).to.equal("johndoe");
    });
  });

  describe("Opening message for edit", () => {
    it("Sets the input field", function() {
      this.message = {...createMessage(), message: "old value"};
      const messageId = this.message.id;
      fetchMock.put(`/api/messages/${messageId}`, {status: 200, body: {}}, {name: "msgPut"});
      this.dispatch(actions.editMessage(this.message));
      const actual = getters.messageInput_getInput(this.getState());
      expect(actual).to.equal('old value');
    })
  });

  describe("Editing an existing message", () => {
    beforeEach(async function() {
      this.message = createMessage();
      const messageId = this.message.id;
      fetchMock.put(`/api/messages/${messageId}`, {status: 200, body: {}}, {name: "msgPut"});
      this.dispatch(actions.editMessage(this.message));
      this.dispatch(actions.setInput("new input"));
      await this.dispatch(actions.addMessage())
    });

    it("sends a PUT request with the message id", async function() {
      const lastCall = fetchMock.lastOptions("msgPut");
      const lastBody = JSON.parse(lastCall.body);
      const expected = { ...this.message, message: "new input" };
      expect(lastBody).to.deep.equal(expected);
    });

    it("clears the input", async function() {
      const actual = getters.messageInput_getInput(this.getState());
      expect(actual).to.equal('');
    });

    describe("subsequent edits and changes", () => {
      beforeEach(async function() {
        fetchMock.post(`/api/messages`, {status: 200, body: {}}, {name: "msgPost"});
        this.dispatch(actions.setInput("new message"));
        await this.dispatch(actions.addMessage())
      });

      it("subsequently posts a new message", function() {
        expect(fetchMock.calls("msgPost")).to.have.length(1);
        expect(fetchMock.calls("msgPut")).to.have.length(1);
      })

      it("clears the input", function() {
        // new test case - clears the input
        const actual = getters.messageInput_getInput(this.getState());
        expect(actual).to.equal('');
      })
    });
  });
});
