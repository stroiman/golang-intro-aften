import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as actions from '../ducks/polling/actions';

const SocketToggler = (props) => {
  const onClick = e => {
    e.preventDefault();
    props.startWebSocket();
  };
  return (
    <button onClick={onClick} className="btn btn-primary">Start socket</button>
  );
}

SocketToggler.propTypes = {
  startWebSocket: PropTypes.func.isRequired
};

export default connect(null, actions)(SocketToggler);
