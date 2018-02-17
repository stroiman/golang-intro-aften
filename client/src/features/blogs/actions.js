export const loadBlogs = () => dispatch => {
  dispatch({
    type: "BLOGS_LOAD"
  });
  return fetch("/api/blogs").then(r => r.json).then(
    x => dispatch({
      type: "BLOGS_LOADED"
    }));
};


