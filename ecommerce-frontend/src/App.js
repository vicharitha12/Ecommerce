import React from 'react';
import { BrowserRouter as Router, Route, Switch, Routes } from 'react-router-dom';
import Login from './components/Login';
import ItemList from './components/ItemList';

function App() {
  return (
    <Router>

      <Routes>
        <Route exact path="/" element={<Login />} />
        <Route path="/items" element={<ItemList/> } />
      </Routes>

    </Router>
  );
}

export default App;
