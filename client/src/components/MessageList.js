import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as getters from '../reducers';

export const Message = () => (<div>Message</div>);

class MessageList extends React.Component {
  render() {
    this.rendered = true;
    return (
      <div>
        { this.props.messages.map(x => <Message key={x.id} />) }
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
