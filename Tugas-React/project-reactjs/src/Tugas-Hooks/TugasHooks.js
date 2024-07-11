import React, { useState, useEffect } from 'react';

const TugasHooks = () => {
  const [currentTime, setCurrentTime] = useState(new Date());
  const [countdown, setCountdown] = useState(100);

  useEffect(() => {
    const timeInterval = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);

    const countdownInterval = setInterval(() => {
      setCountdown((prevCountdown) => {
        if (prevCountdown > 0) {
          return prevCountdown - 1;
        } else {
          clearInterval(countdownInterval);
          return prevCountdown;
        }
      });
    }, 1000);

    return () => {
      clearInterval(timeInterval);
      clearInterval(countdownInterval);
    };
  }, []);

  if (countdown <= 0) {
    return null;
  }

  return (
     <div className="task-container">
      <div className="time-container">
        <h2>Now At - {currentTime.toLocaleTimeString()}</h2>
        <h2>Countdown: {countdown}</h2>
      </div>
    </div>
  );
};

export default TugasHooks;
