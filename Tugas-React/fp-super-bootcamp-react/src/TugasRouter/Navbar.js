// src/TugasRouter/Navbar.js
import React, { useContext } from 'react';
import { ThemeContext } from './ThemeContext';
import { Link } from 'react-router-dom';
import './Navbar.css';

function Navbar() {
  const { theme, toggleTheme } = useContext(ThemeContext);

  return (
    <nav className={`nav ${theme}`}>
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/crud-hooks">CRUD Hooks</Link></li>
        <li><Link to="/axios">Axios</Link></li>
        <li><Link to="/context">Context</Link></li>
        <li><button onClick={toggleTheme} className="theme-toggle">
          Switch Theme
        </button></li>
      </ul>
    </nav>
  );
}

export default Navbar;
