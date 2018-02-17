import React, { Component } from 'react';
import logo from './logo.svg';
//import './App.css';
import 'bootstrap/dist/css/bootstrap.css';
import gopher from './gopher-front.svg';

class App extends Component {
  render() {
    return (
      <div>
        <div className="navbar bg-dark">
          <div className="container">
            <a className="navbar-brand" href="#">
              <img src={gopher} className="App-logo" alt="logo" style={{width:48, height: 48}} />
            </a>
          </div>
        </div>
        <div className="App container">
          <header className="App-header">
            <h1 className="App-title">Welcome to React</h1>
          </header>
          <p className="App-intro">
            To get started, edit <code>src/App.js</code> and save to reload.
          </p>
        </div>
      </div>
    );
  }
}

export default App;
