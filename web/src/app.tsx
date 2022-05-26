import React from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

import { SearchEvents } from './components/search_events';
import { ApiClient } from './api_client';

const client = new ApiClient('http://localhost:8000');

export const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/events" />} />
        <Route path="/events" element={<SearchEvents client={client} />} />
      </Routes>
    </BrowserRouter>
  )
};
