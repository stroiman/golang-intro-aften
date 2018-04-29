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
    this.onCancel = this.onCancel.bind(this);
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

  onCancel(e) {
    e.preventDefault();
    this.props.cancelEditing();
  }

  render() {
    return(
      <form onSubmit={this.onSubmit}>
        <div className="form-group mb-0 input-group">
          <input className="form-control" type="text" placeholder="Skriv noget..."
            autoFocus onChange={this.onChange} value={this.state.message} />
          { this.props.isEditing &&
              <div className="input-group-append">
                <button type="button" className="btn btn-secondary" onClick={this.onCancel} >Afbryd</button>
              </div>
          }
        </div>
      </form>);
  }
}

const mapStateToProps = state => ({
  message: getters.messageInput_getInput(state),
  isEditing: getters.messageInput_isEditing(state),
});

export default connect(mapStateToProps, actions)(MessageInput);
