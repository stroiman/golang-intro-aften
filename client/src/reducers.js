import { combineReducers } from 'redux';
import blogs, * as fromBlogs from './features/blogs/reducers';

export default combineReducers({
  blogs
});

export const getBlogsLoadingState = state => fromBlogs.getLoadingState(state.blogs);

