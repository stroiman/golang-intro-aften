export default (state = "", action) => {
  switch (action.type) {
    case "MESSAGE_INPUT_SET": return action.payload;
    default: return "";
  }
};

export const getInput = state => state;
