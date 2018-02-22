import * as getters from '../../reducers';
import { useRedux } from '../../testHelpers/reduxHelpers';

describe("messages duck", () => {
  useRedux();

  describe("initial state", () => {
    it("is an empty array", function() {
      const actual = getters.messages_getDisplayMessages(this.getState())
      expect(actual).to.deep.equal([]);
    });

    it("has loadingState=NOT_LOADED", function() {
      const actual = getters.messages_getLoadingState(this.getState())
      expect(actual).to.deep.equal("NOT_LOADED");
    });
  });
});
