import React from 'react';
import gopher from '../../gopher-front.svg';

export default props => (
  <div className="bg-dark" style={{display: "flex", flexDirection: "column", height: "100vh", maxHeight: "100vh" }}>
    <div className="valign-center">
      <img src={gopher} className="App-logo" alt="logo" style={{width:48, height: 48}} />
      <span className="h1 text-white" style={{verticalAlign: "middle", fontSize: "40px"}}>Go<em>ssip</em></span>
      <h1 className="text-white">Login</h1>
    </div>
  </div>
);
