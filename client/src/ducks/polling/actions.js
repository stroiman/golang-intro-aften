import * as getters from '../../reducers';
import { fetchMessages } from '../messages/actions';

const TIMER_INTERVAL = 1000;

const setTimerId = timerId => ({
  type: "MESSAGES_SETTIMERID",
  payload: timerId
});

export const togglePolling = () => async (dispatch, getState) => {
  const state = getState();
  if (getters.polling_isPolling(state)) {
    clearInterval(getters.polling_getTimerId(state));
    dispatch(setTimerId(null));
  } else {
    const timerId = setInterval(() => dispatch(fetchMessages()), TIMER_INTERVAL);
    dispatch(setTimerId(timerId));
  }
};
