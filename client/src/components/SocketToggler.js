import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as actions from '../ducks/polling/actions';
import classNames from 'classnames';

const SocketToggler = (props) => {
  const onClick = e => {
    e.preventDefault();
    props.startWebSocket();
  };
  const className = classNames("btn", "btn-primary", props.className);
  return (
    <button onClick={onClick} className={className}>Start socket</button>
  );
}

SocketToggler.propTypes = {
  startWebSocket: PropTypes.func.isRequired
};

export default connect(null, actions)(SocketToggler);
