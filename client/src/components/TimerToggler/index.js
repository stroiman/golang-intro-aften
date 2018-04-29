import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
//import * as actions from '../../ducks/polling/actions';
import * as actions from './actions';
import * as getters from '../../reducers';

const TimerToggler = (props) => {
  const onClick = e => {
    e && e.preventDefault && e.preventDefault();
    props.togglePolling();
  };
  // const message = props.isPolling ? "Polling enabled" : "Polling disabled";
  const checked = props.isPolling ? "checked" : "";
  return (
    <button onClick={onClick} className="btn btn-primary">
      <input type="checkbox" checked={checked} className="mr-1" onChange={onClick} />
      Polling enabled
    </button>
  );
};

TimerToggler.propTypes = {
  isPolling: PropTypes.bool.isRequired,
  togglePolling: PropTypes.func.isRequired
}

const mapStateToProps = state => ({
  isPolling: getters.polling_isPolling(state)
});

export default connect(mapStateToProps, actions)(TimerToggler)
