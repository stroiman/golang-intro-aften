import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as getters from '../reducers';
import * as actions from './actions'

export class Message extends React.Component {
  constructor(props) {
    super(props);
    this.editMessage = this.editMessage.bind(this);
  }

  editMessage(e) {
    e.preventDefault();
    this.props.editMessage(this.props.message);
  };

  render() {
    const message = this.props.message;
    return (
      <div className="card bg-light mb-3">
        <div className="card-body" style={{padding: "0.5rem"}} >
          {message.message}
        </div>
        <button role="edit" onClick={this.editMessage}>edit</button>
      </div>
    );
  }
};

class MessageList extends React.Component {
  scrollToBottom () {
    const shouldScroll =
      this.messagesEnd && this.messagesEnd.scrollIntoView &&
      this.knowMessageCount != this.props.messages.length
    shouldScroll && this.messagesEnd.scrollIntoView({ behavior: "smooth" });
    this.knowMessageCount = this.props.messages.length;
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
        { this.props.messages.map(x => <Message key={x.id} message={x} editMessage={this.props.editMessage} />) }
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

export default connect(mapStateToProps, actions)(MessageList);
