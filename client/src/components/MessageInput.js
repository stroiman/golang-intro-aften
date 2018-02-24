import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as actions from '../ducks/message-input/actions';

class MessageInput extends React.Component {
  constructor(props) {
    super(props)
    this.state = { message: "" }

    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onChange(e) {
    this.setState({message: e.target.value});
  }

  onSubmit(e) {
    e.preventDefault();
    this.props.addMessage(this.state.message);
  }

  render() {
    return(
      <form onSubmit={this.onSubmit}>
        <div className="form-group mb-0">
          <input className="form-control" type="text" placeholder="Skriv noget..."
            autoFocus onChange={this.onChange} value={this.state.message} />
        </div>
      </form>);
  }
}

export default connect(() => ({}), actions)(MessageInput);
