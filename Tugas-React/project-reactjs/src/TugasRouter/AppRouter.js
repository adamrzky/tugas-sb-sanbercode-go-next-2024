import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import TugasHooks from '../Tugas-Hooks/TugasHooks';
import TugasIntroReactJS from '../Tugas-Intro-ReactJS/TugasIntroReactJS';
import BooksAxios from '../TugasAxios/Books';
import BooksContext from '../TugasContext/Books';
import TugasCRUDHooks from '../TugasCRUDHooks/TugasCRUDHooks';

const AppRouter = () => {
  return (
    <Router>
      <div>
        <nav>
          <ul className="nav">
            <li><Link to="">Home</Link></li>
            <li><Link to="/crud-hooks">CRUD Hooks</Link></li>
            <li><Link to="/axios">Axios</Link></li>
            <li><Link to="/context">Context</Link></li>
 
          </ul>
        </nav>
        <Routes>
          <Route path="/" element={<TugasIntroReactJS />} />
          <Route path="/crud-hooks" element={<TugasCRUDHooks />} />
          <Route path="/axios" element={<BooksAxios />} />
          <Route path="/context" element={<BooksContext />} />

        </Routes>
      </div>
    </Router>
  );
};

export default AppRouter;
