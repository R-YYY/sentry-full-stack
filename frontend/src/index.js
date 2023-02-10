import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import * as Sentry from "@sentry/browser";
import { BrowserTracing } from "@sentry/tracing";
import App from './App';
import reportWebVitals from './reportWebVitals';

Sentry.init({
  dsn: "https://49d0e621960545d3be74c96f364989ca@o4504648733032448.ingest.sentry.io/4504654757691392",

  integrations:[
    new BrowserTracing({
      tracePropagationTargets: ['localhost'],
    })
  ],

  // Set tracesSampleRate to 1.0 to capture 100%
  // of transactions for performance monitoring.
  // We recommend adjusting this value in production
  tracesSampleRate: 1.0,
});

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
