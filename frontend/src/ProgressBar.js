import React, {useState, useEffect} from 'react';
import './ProgressBar.css';
const Range = (props) => {
    return (
        <div className="range" style={{width: `${props.percentRange}%`}}/>
    );
};

const ProgressBar = (props) => {
  return (
      <div className="progress-bar">
          <Range percentRange={props.percentRange}/>
      </div>
  );
};



export const ProgressBarContainer = props => {
  let [percentRange, setProgress] = useState(0);
  return (
      <div className="container">
          <ProgressBar percentRange={props.percentRange}/>
      </div>
  );
};
