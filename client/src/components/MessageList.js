import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as getters from '../reducers';
import * as actions from './actions';
import * as icons from './icons';

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
        <div className="card-body media" style={{padding: "0.5rem"}} >
          <div className="media-body">
            {message.message}
          </div>
          <a role="edit" onClick={this.editMessage} aria-label="edit"><icons.Edit aria-hidden="true" /></a>
        </div>
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
