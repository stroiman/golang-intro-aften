import React from 'react';
import PropTypes from 'prop-types';

class MessageInput extends React.Component {
  constructor(props) {
    super(props)
    this.state = { message: "" }

    this.onChange = this.onChange.bind(this);
  }

  onChange(e) {
    this.setState({message: e.target.value});
  }

  render() {
    return(
      <form>
        <div className="form-group mb-0">
          <input className="form-control" type="text" placeholder="Skriv noget..."
            onChange={this.onChange} value={this.state.message} />
        </div>
      </form>);
  }
}

export default MessageInput;
