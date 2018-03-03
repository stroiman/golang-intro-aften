import React, { Component } from 'react';
import { connect } from 'react-redux';
import gopher from './gopher-front.svg';
import MessageList from './components/MessageList';
import MessageInput from './components/MessageInput';
import TimerToggler from './components/TimerToggler';
import SocketToggler from './components/SocketToggler';
import LoginPage from './pages/LoginPage';

import 'bootstrap/dist/css/bootstrap.css';

export const MessagesPage = props => (
  <div style={{display: "flex", flexDirection: "column", height: "100vh", maxHeight: "100vh" }}>
    <div className="navbar bg-dark" style={{paddingLeft: "1rem"}} >
      <a className="navbar-brand" href="#">
        <img src={gopher} className="App-logo" alt="logo" style={{width:48, height: 48}} />
        <span className="h1 text-white" style={{verticalAlign: "middle", fontSize: "40px"}}>Go<em>ssip</em></span>
      </a>
      <SocketToggler />
      <TimerToggler />
    </div>
    <div style={{flexGrow: "1", padding: "1rem", overflow: "auto" }}>
      <MessageList />
    </div>
    <div style={{padding: "1rem", borderTop: "1px solid #CCC" }}>
      <MessageInput />
    </div>
  </div>
);

class App extends Component {
  render() {
    return this.props.user ? <MessagesPage /> : <LoginPage />
  }
}

const mapStateToProps = state => ({
  user: state.auth
});

export default connect(mapStateToProps)(App);
