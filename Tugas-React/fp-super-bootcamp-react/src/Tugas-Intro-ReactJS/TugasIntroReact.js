import React, { useState } from 'react';
// import './App.css';

// Komponen untuk Checkbox
const Checkbox = ({ label }) => (
  <div className="checkbox">
    <input type="checkbox" id={label} />
    <label htmlFor={label}>{label}</label>
  </div>
);

// Data daftar kegiatan
let thingsToDo = [
  "Belajar GIT & CLI",
  "Belajar HTML & CSS",
  "Belajar Javascript",
  "Belajar ReactJS Dasar",
  "Belajar ReactJS Advance"
];

const TugasIntroReact = () => {
  const [isClicked, setIsClicked] = useState(false);

  const handleClick = () => {
    setIsClicked(!isClicked);
  };

  return (
    <div className="container">
      <img src="logo1.png" alt="Sanbercode Logo" className="logo" />
      <h1>THINGS TO DO</h1>
      <p>During bootcamp in sanbercode</p>
      <hr />
      <div className="checkbox-container">
        {thingsToDo.map((item, index) => (
          <Checkbox key={index} label={item} />
        ))}
      </div>
      <button
        type="button"
        className={`submit-button ${isClicked ? 'clicked' : ''}`}
        onClick={handleClick}
      >
        SEND
      </button>
    </div>
  );
};

export default TugasIntroReact;
