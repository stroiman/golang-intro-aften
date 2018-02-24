import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as actions from '../ducks/message-input/actions';
import * as getters from '../reducers';

class MessageInput extends React.Component {
  static propTypes = {
    message: PropTypes.string.isRequired,
    setInput: PropTypes.func.isRequired,
    addMessage: PropTypes.func.isRequired
  };

  constructor(props) {
    super(props)
    this.state = { message: props.message || "" }

    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentWillReceiveProps(newProps) {
    const { message } = newProps
    this.setState({message});
  }

  onChange(e) {
    this.props.setInput(e.target.value);
  }

  onSubmit(e) {
    e.preventDefault();
    this.props.addMessage();
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

const mapStateToProps = state => ({
  message: getters.messageInput_getInput(state)
});

export default connect(mapStateToProps, actions)(MessageInput);
