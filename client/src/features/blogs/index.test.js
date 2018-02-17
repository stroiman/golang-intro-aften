import { createStore, applyMiddleware, combineReducers } from 'redux';
import thunk from 'redux-thunk';
import fetchMock from 'fetch-mock';
import blogsReducer from './reducers';
import * as actions from './actions';
import reducers, * as fromReducers from '../../reducers';

const getLoadingState = fromReducers.getBlogsLoadingState;

describe("blogs", () => {
  beforeEach(function() {
    this.store = createStore(reducers, applyMiddleware(thunk));
    this.getState = this.store.getState;
    this.dispatch = this.store.dispatch;
  });

  afterEach(() => {
    fetchMock.restore();
  });

  describe("load", () => {
    describe("initial state", () => {
      it("is NOT_LOADING", function() {
        expect(getLoadingState(this.getState())).to.equal("NOT_LOADING");
      });
    });

    context("No response has been returned", () => {
      it("has status LOADING", function() {
        fetchMock.get("/api/blogs", new Promise((resolve,reject) => {}));
        this.dispatch(actions.loadBlogs());
        expect(getLoadingState(this.getState())).to.equal("LOADING");
      });
    });

    context("A response has been returned", () => {
      it("has a status LOADED", async function() {
        fetchMock.get("/api/blogs", Promise.resolve({}));
        await this.dispatch(actions.loadBlogs());
        expect(getLoadingState(this.getState())).to.equal("LOADED");
      });
    });
  });
});
