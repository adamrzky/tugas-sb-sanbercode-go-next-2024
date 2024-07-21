import React from 'react';
import AppRouter from './TugasRouter/AppRouter';
import { ThemeProvider } from './TugasRouter/ThemeContext';

function App() {
  return (
    <ThemeProvider>
      <AppRouter />
    </ThemeProvider>
  );
}

export default App;
