import React from 'react';
import { connect } from 'react-redux';
import gopher from '../../gopher-front.svg';
import * as actions from '../../ducks/auth/actions';

class LoginPage extends React.Component {
  constructor(props) {
    super(props)
    this.state = { userName: '' };

    this.onChange = this.onChange.bind(this);
    this.onClick = this.onClick.bind(this);
  }

  onChange(e) {
    this.setState({userName: e.target.value});
  }

  onClick (e) {
    e.preventDefault();
    this.props.loginUser({ username: this.state.userName });
  }
  render() {
    return (
      <div className="bg-dark" style={{display: "flex",
        flexDirection: "column", height: "100vh", maxHeight: "100vh",
        alighItems: "center",
        justifyContent: "center",
      }}>
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-md-3">
            <div className="mb-3 text-align-center">
              <img src={gopher} className="App-logo" alt="logo" style={{width:48, height: 48}} />
              <span className="h1 text-white" style={{verticalAlign: "middle", fontSize: "40px"}}>Go<em>ssip</em></span>
            </div>
            <form>
              <div className="form-group">
                <input className="form-control" placeholder="username"
                  value={this.state.userName} onChange={this.onChange} />
              </div>
              <button className="btn btn-outline-success form-control" onClick={this.onClick} >Login</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    );
  }
}

export default connect(null, actions)(LoginPage);
