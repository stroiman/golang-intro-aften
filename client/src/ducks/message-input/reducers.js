export default (state = "", action) => {
  switch (action.type) {
    case "MESSAGE_INPUT_SET": return action.payload;
    case "MESSAGE_POSTED": return "";
    default: return state;
  }
};

export const getInput = state => state;
