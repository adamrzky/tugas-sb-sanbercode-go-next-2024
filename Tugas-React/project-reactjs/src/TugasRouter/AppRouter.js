// src/TugasRouter/AppRouter.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import TugasHooks from '../Tugas-Hooks/TugasHooks'; 
import TugasIntroReactJS from '../Tugas-Intro-ReactJS/TugasIntroReactJS'; 
import BooksAxios from '../TugasAxios/Books';
import BooksContext from '../TugasContext/Books';
import TugasCRUDHooks from '../TugasCRUDHooks/TugasCRUDHooks';
import Navbar from './Navbar';

const AppRouter = () => {
  return (
    <Router>
      <div>
        <Navbar />
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
