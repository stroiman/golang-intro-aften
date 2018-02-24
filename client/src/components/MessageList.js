import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as getters from '../reducers';

export const Message = ({message}) => (
  <div className="card bg-light mb-3">
    <div className="card-body" style={{padding: "0.5rem"}} >
      {message.message}
    </div>
  </div>
);

class MessageList extends React.Component {
  render() {
    this.rendered = true;
    return (
      <div>
        { this.props.messages.map(x => <Message key={x.id} message={x} />) }
      </div>);
  }
}

MessageList.propTypes = {
  messages: PropTypes.array.isRequired
};

const mapStateToProps = state => ({
  messages: getters.messages_getDisplayMessages(state)
});

export default connect(mapStateToProps)(MessageList);
