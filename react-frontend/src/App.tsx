// import React from 'react';

import { CssBaseline, MuiThemeProvider } from "@material-ui/core";
import { Mapping } from "./components/Mapping";
import theme from './theme';


function App() {
  return (
    <div className="App">
      <MuiThemeProvider theme={theme}>
        <CssBaseline/>
      <Mapping/>
      </MuiThemeProvider>
    </div>
  );
}

export default App;
