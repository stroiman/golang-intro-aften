import React, { Component } from 'react';
//import './App.css';
import 'bootstrap/dist/css/bootstrap.css';
import gopher from './gopher-front.svg';
import MessageList from './components/MessageList';

class App extends Component {
  render() {
    return (
      <div style={{display: "flex", flexDirection: "column", height: "100vh", maxHeight: "100vh" }}>
        <div className="navbar bg-dark" style={{paddingLeft: "1rem"}} >
          <a className="navbar-brand" href="#">
            <img src={gopher} className="App-logo" alt="logo" style={{width:48, height: 48}} />
            <span className="h1 text-white" style={{verticalAlign: "middle", fontSize: "40px"}}>Go<em>ssip</em></span>
          </a>
        </div>
        <div style={{flexGrow: "1", margin: "1rem" }}>
          <MessageList />
        </div>
        <div>
          <form>
            <div className="form-group" style={{margin: "1rem" }}>
              <input className="form-control" type="text" placeholder="Skriv noget..." />
            </div>
          </form>
        </div>
      </div>
    );
  }
}

export default App;
