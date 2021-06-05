import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Donation, FaceBookLoginProvider } from './components';

function App() {
  return (
    <div className="App">
      <b>FB Login</b>
      <FaceBookLoginProvider>
        <Donation></Donation>
      </FaceBookLoginProvider>
    </div>
  );
}

export default App;
