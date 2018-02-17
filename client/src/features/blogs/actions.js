export const loadBlogs = () => async dispatch => {
  dispatch({
    type: "BLOGS_LOAD",
    payload: {
      status: "LOADING"
    }
  });
  let response = await fetch("/api/blogs");
  if (!response.ok) {
    dispatch({
      type: "BLOGS_LOAD",
      payload: {
        status: "LOAD_FAILED"
      }
    });
    return;
  };
  dispatch({
    type: "BLOGS_LOAD",
    payload: {
      status: "LOADED",
    }
  });
};


