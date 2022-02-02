import React from 'react';
import { render, screen } from '@testing-library/react';
import App, { BuildAPIEndpointURL } from './App';
import { BrowserRouter } from 'react-router-dom';

it('renders learn react link', () => {
  render(<BrowserRouter><App /></BrowserRouter>);
});

it('buildAPIEndpointURL returns expected url', () => {
  const inputSearchParams = new URLSearchParams("c=12&b=12&a=abc");
  const outcome = BuildAPIEndpointURL(inputSearchParams);
  const expected = process.env.REACT_APP_RACER_READY_API + "?" + inputSearchParams.toString()
  
  expect(outcome).toEqual(expected)
});
