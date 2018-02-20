import React from 'react';
import PropTypes from 'prop-types';

const idFor = name => `input-${name}`;
const helpFor = name => `${name}-help`;

class TextField extends React.Component {
  constructor(props) {
    super(props);
    this.onChange = this.onChange.bind(this);
  };

  onChange (e) {
    const target = e.target;
    const arg = {[this.props.name]: target.value};
    this.context.onChange && this.context.onChange(arg);
  };

  render() {
    const { name, helpText } = this.props;
    const value = this.context.values[name];
    return (
      <div className="form-group">
        <label htmlFor={idFor(name)}>{this.props.heading || ""}</label>
        <input type="text" className="form-control" id={idFor(name)}
          aria-describedby={helpFor(name)} value={value} onChange={this.onChange} />
        { helpText && <small id={helpFor(name)} className="text-muted form-text">{helpText}</small> }
      </div>
    )}
};

TextField.contextTypes = {
  values: PropTypes.object,
  onChange: PropTypes.func
};

TextField.propTypes = {
  name: PropTypes.string.isRequired,
  helpText: PropTypes.string
};

class TextAreaField extends React.Component {
  constructor(props) {
    super(props);
    this.onChange = this.onChange.bind(this);
  };

  onChange (e) {
    const target = e.target;
    const arg = {[this.props.name]: target.value};
    this.context.onChange && this.context.onChange(arg);
  };

  render() {
    const { name, helpText } = this.props;
    const value = this.context.values[name];
    return (
      <div className="form-group">
        <label htmlFor={idFor(name)}>{this.props.heading || ""}</label>
        <textarea className="form-control" id={idFor(name)}
          aria-describedby={helpFor(name)} value={value} onChange={this.onChange} />
        { helpText && <small id={helpFor(name)} className="text-muted form-text">{helpText}</small> }
      </div>
    )}
};

TextAreaField.contextTypes = {
  values: PropTypes.object,
  onChange: PropTypes.func
};

TextAreaField.propTypes = {
  name: PropTypes.string.isRequired,
  helpText: PropTypes.string
};

class Form extends React.Component {
  static TextField = TextField;
  static TextAreaField = TextAreaField;

  constructor(props) {
    super(props);
    this.state = { values: props.values };
    this.onChange = this.onChange.bind(this);
  };

  componentWillReceiveProps(newProps) {
    const { values } = newProps;
    this.setState({values});
  }

  getChildContext() {
    return {
      values: this.state.values,
      onChange: this.onChange
    };
  }

  onChange(name, value) {
    this.props.onChange(name, value)
  }

  render() {
    return (
      <form>
        { this.props.children }
      </form>);
  }
}

Form.childContextTypes = {
  values: PropTypes.object,
  onChange: PropTypes.func
};

export default Form;
