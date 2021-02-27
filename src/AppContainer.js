// import React, { useState, useEffect } from 'react';
import App from './App/App';
import data from './vendorData.json';

// function mockData() {
//   return data;
// }

function AppContainer() {
  //Use incase data was moved to an extrernal source
  //   const [data, setData] = useState([]);
  //   useEffect(() => {
  //     setTimeout(() => {
  //       setData(mockData());
  //     }, 600);
  //   }, []);

  return <App data={data} />;
}

export default AppContainer;
