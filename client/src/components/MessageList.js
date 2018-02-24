import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as getters from '../reducers';

export class Message extends React.Component {
  render() {
    const message = this.props.message;
    return (
      <div className="card bg-light mb-3">
        <div className="card-body" style={{padding: "0.5rem"}} >
          {message.message}
        </div>
      </div>
    );
  }
};

class MessageList extends React.Component {
  scrollToBottom () {
    this.messageEnd && this.messagesEnd.scrollIntoView({ behavior: "smooth" });
  }

  componentDidMount() {
    this.scrollToBottom();
  }

  componentDidUpdate() {
    this.scrollToBottom();
  }

  render() {
    return (
      <div>
        { this.props.messages.map(x => <Message key={x.id} message={x} />) }
        <div style={{ float:"left", clear: "both" }}
          ref={(el) => { this.messagesEnd = el; }}>
        </div>
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
