import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';
import { BrowserRouter } from 'react-router-dom';

it('renders learn react link', () => {
  render(<BrowserRouter><App /></BrowserRouter>);
});
