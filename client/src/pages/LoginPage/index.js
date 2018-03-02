import React from 'react';
import gopher from '../../gopher-front.svg';

const LoginPage = props => (
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
            <input className="form-control" placeholder="username" />
          </div>
          <div className="form-group">
            <input className="form-control" placeholder="e-mail (optional)" />
            <small className="text-muted">
              If you have a gravatar setup, you can enter this, and get your
              gravatar logo next to your posts.
            </small>
          </div>
          <button className="btn btn-outline-success form-control">Login</button>
        </form>
      </div>
    </div>
  </div>
</div>
);

export default LoginPage;
