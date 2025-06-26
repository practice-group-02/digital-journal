import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Register from './Register';
import Login from './Login';

function App() {
  return (
    <Router>
      <nav style={{ padding: '1rem', background: '#f4f4f4' }}>
        <Link to="/register" style={{ marginRight: '1rem' }}>Register</Link>
        <Link to="/login">Login</Link>
      </nav>
      <div style={{ padding: '1rem' }}>
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;